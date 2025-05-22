package bot

import (
	"context"
	"log/slog"
	"time"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/ashokhin/am4bot/internal/utils"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func (b *Bot) hubs(ctx context.Context) error {
	var needRepair bool

	slog.Info("check hubs")

	// check Alert icon for lounge on the "Flight Info" menu
	needRepair = utils.IsElementVisible(ctx, model.ICON_FI_LOUNGE_ALERT)

	slog.Debug("repair status", "need_repair", needRepair)
	slog.Debug("open pop-up window", "window", "hubs")

	if err := chromedp.Run(ctx,
		utils.ClickElement(model.BUTTON_MAIN_HUBS),
	); err != nil {
		slog.Warn("error in Bot.hubs > open hubs", "error", err)

		return err
	}

	defer utils.DoClickElement(ctx, model.BUTTON_COMMON_CLOSE_POPUP)

	slog.Debug("get list of hubs")

	var hubsElemList []*cdp.Node

	// get list of Hubs
	if err := chromedp.Run(ctx,
		chromedp.Nodes(model.LIST_HUBS_HUBS, &hubsElemList, chromedp.ByQueryAll),
	); err != nil {
		slog.Warn("error in Bot.hubs > get hubs list", "error", err)

		return err
	}

	// get metrics for all hubs in humElemList
	for _, hubElem := range hubsElemList {
		var hub model.Hub

		slog.Debug("hubElem", "elem", hubElem)

		if err := chromedp.Run(ctx,
			chromedp.Text(model.TEXT_HUBS_HUB_NAME, &hub.Name, chromedp.ByQuery, chromedp.FromNode(hubElem)),
			utils.GetFloatFromChildElement(model.TEXT_HUBS_HUB_DEPARTURES, &hub.Departures, hubElem),
			utils.GetFloatFromChildElement(model.TEXT_HUBS_HUB_ARRIVALS, &hub.Arrivals, hubElem),
			utils.GetFloatFromChildElement(model.TEXT_HUBS_HUB_PAX_DEPARTED, &hub.PaxDeparted, hubElem),
			utils.GetFloatFromChildElement(model.TEXT_HUBS_HUB_PAX_ARRIVED, &hub.PaxArrived, hubElem),
		); err != nil {
			slog.Warn("error in Bot.hubs > get gub info", "error", err)

			return err
		}

		slog.Debug("hub found", "name", hub.Name)

		hub.HasCatering = utils.IsSubElementVisible(ctx, model.ICON_HUBS_CATERING, hubElem)

		b.PrometheusMetrics.HubStats.WithLabelValues(hub.Name, "departures").Set(hub.Departures)
		b.PrometheusMetrics.HubStats.WithLabelValues(hub.Name, "arrivals").Set(hub.Arrivals)
		b.PrometheusMetrics.HubStats.WithLabelValues(hub.Name, "paxDeparted").Set(hub.PaxDeparted)
		b.PrometheusMetrics.HubStats.WithLabelValues(hub.Name, "paxArrived").Set(hub.PaxArrived)

		// repair lounge if global 'Alert icon' for lounges is displayed
		if needRepair {
			slog.Debug("repair lounge", "lounge", hub.Name)

			if err := b.repairLounge(ctx, hubElem); err != nil {
				slog.Warn("error in Bot.hubs > Bot.repairLounge", "error", err)

				return err
			}
		}

		if !hub.HasCatering {
			if err := b.buyCatering(ctx, hub, hubElem); err != nil {
				slog.Warn("error in Bot.hubs > Bot.buyCatering", "error", err)

				return err
			}
		}
	}

	slog.Debug("close pop-up window", "window", "hubs")

	return nil
}

func (b *Bot) repairLounge(ctx context.Context, hubElem *cdp.Node) error {
	slog.Debug("repair lounge")

	if err := chromedp.Run(ctx,
		chromedp.Click(model.ELEMENT_HUB, chromedp.ByQuery, chromedp.FromNode(hubElem)),
	); err != nil {
		slog.Warn("error in Bot.repairLounge > select hub", "error", err)

		return err
	}

	// return to list of hubs when exiting from function
	defer utils.DoClickElement(ctx, model.BUTTON_HUBS_HUB_MANAGE_BACK)

	if !utils.IsElementVisible(ctx, model.BUTTON_HUBS_HUB_MANAGE) {
		slog.Warn("button 'Manage' isn't visible")

		return nil
	}

	utils.DoClickElement(ctx, model.BUTTON_HUBS_HUB_MANAGE)

	if !utils.IsElementVisible(ctx, model.BUTTON_HUBS_HUB_MANAGE_REPAIR) {
		slog.Warn("button 'Repair' isn't visible")

		return nil
	}

	var repairCost float64

	if err := chromedp.Run(ctx,
		utils.GetFloatFromElement(model.TEXT_HUBS_HUB_MANAGE_REPAIR_COST, &repairCost),
	); err != nil {
		slog.Warn("error in Bot.repairLounge > get repair cost", "error", err)

		return err
	}

	slog.Debug("repair cost", "value", int(repairCost))

	slog.Debug("available money", "value", int(b.Conf.BudgetMoney.Maintenance))

	if repairCost > b.Conf.BudgetMoney.Maintenance {
		slog.Warn("lounge repair is too expensive", "cost", int(repairCost),
			"budget", int(b.Conf.BudgetMoney.Maintenance))

		return nil
	}

	slog.Info("repair lounge", "repairCost", int(repairCost),
		"BudgetMoney.Maintenance", int(b.Conf.BudgetMoney.Maintenance))

	utils.DoClickElement(ctx, model.BUTTON_HUBS_HUB_MANAGE_REPAIR)

	// reduce current account money and maintenance budged by repair cost
	b.AccountBalance -= repairCost
	b.Conf.BudgetMoney.Maintenance -= repairCost

	return nil
}

func (b *Bot) buyCatering(ctx context.Context, hub model.Hub, hubElem *cdp.Node) error {
	slog.Debug("Buy catering", "hub", hub.Name)

	if err := chromedp.Run(ctx,
		chromedp.Click(model.ELEMENT_HUB, chromedp.ByQuery, chromedp.FromNode(hubElem)),
	); err != nil {
		slog.Warn("error in Bot.buyCatering > select hub", "error", err)

		return err
	}

	// return to list of hubs when exiting from function
	defer utils.DoClickElement(ctx, model.BUTTON_HUBS_HUB_MANAGE_BACK)

	if !utils.IsElementVisible(ctx, model.BUTTON_HUBS_ADD_CATERING) {
		slog.Warn("button '+ Add catering' isn't visible")

		return nil
	}

	var cateringCost float64

	if err := chromedp.Run(ctx,
		utils.ClickElement(model.BUTTON_HUBS_ADD_CATERING),
		chromedp.WaitReady(model.ELEM_HUBS_CATERING_OPTION_3, chromedp.ByQuery),
		utils.ClickElement(model.ELEM_HUBS_CATERING_OPTION_3),
		chromedp.SetValue(model.SELECT_HUBS_CATERING_DURATION, model.OPTION_HUBS_CATERING_DURATION_1W, chromedp.ByQuery),
		chromedp.SetValue(model.SELECT_HUBS_CATERING_AMOUNT, model.OPTION_HUBS_CATERING_AMOUNT_20000, chromedp.ByQuery),
		utils.GetFloatFromElement(model.TEXT_HUBS_CATERING_COST, &cateringCost),
	); err != nil {
		slog.Warn("error in Bot.buyCatering > select hub", "error", err)

		return err
	}

	if cateringCost > b.Conf.BudgetMoney.Maintenance {
		slog.Warn("catering is too expensive", "cost", int(cateringCost),
			"budget", int(b.Conf.BudgetMoney.Maintenance), "hub", hub.Name)

		return nil
	}

	time.Sleep(5 * time.Second)

	// buy catering
	if err := chromedp.Run(ctx,
		utils.ClickElement(model.BUTTON_HUBS_CATERING_BUY),
	); err != nil {
		slog.Warn("error in Bot.buyCatering > buy catering", "hub", hub.Name, "error", err)

		return err
	}

	time.Sleep(10 * time.Second)

	return nil
}
