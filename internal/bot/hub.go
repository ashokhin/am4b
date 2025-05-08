package bot

import (
	"context"
	"log/slog"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/ashokhin/am4bot/internal/utils"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Hub struct {
	name        string
	departures  float64
	arrivals    float64
	paxDeparted float64
	paxArrived  float64
}

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
		var hub Hub

		slog.Debug("hubElem", "elem", &hubElem)

		if err := chromedp.Run(ctx,
			chromedp.Text(model.TEXT_HUBS_HUB_NAME, &hub.name, chromedp.ByQuery, chromedp.FromNode(hubElem)),
			utils.GetFloatFromChildElement(model.TEXT_HUBS_HUB_DEPARTURES, &hub.departures, hubElem),
			utils.GetFloatFromChildElement(model.TEXT_HUBS_HUB_ARRIVALS, &hub.arrivals, hubElem),
			utils.GetFloatFromChildElement(model.TEXT_HUBS_HUB_PAX_DEPARTED, &hub.paxDeparted, hubElem),
			utils.GetFloatFromChildElement(model.TEXT_HUBS_HUB_PAX_ARRIVED, &hub.paxArrived, hubElem),
		); err != nil {
			slog.Warn("error in Bot.hubs > get gub info", "error", err)

			return err
		}

		slog.Debug("hub found", "name", hub.name)

		b.PrometheusMetrics.HubStats.WithLabelValues(hub.name, "departures").Set(hub.departures)
		b.PrometheusMetrics.HubStats.WithLabelValues(hub.name, "arrivals").Set(hub.arrivals)
		b.PrometheusMetrics.HubStats.WithLabelValues(hub.name, "paxDeparted").Set(hub.paxDeparted)
		b.PrometheusMetrics.HubStats.WithLabelValues(hub.name, "paxArrived").Set(hub.paxArrived)

		// repair lounge if global 'Alert icon' for lounges is displayed
		if needRepair {
			slog.Debug("repair lounge", "lounge", hub.name)

			if err := b.repairLounge(ctx, hubElem); err != nil {
				slog.Warn("error in Bot.hubs > Bot.repairLounge", "error", err)

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

	// return to list of hubs when editing from function
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

	slog.Debug("repair cost", "value", repairCost)

	slog.Debug("available money", "value", b.Conf.BudgetMoney.Maintenance)

	if b.Conf.BudgetMoney.Maintenance > repairCost {
		slog.Info("repair lounge", "repairCost", int(repairCost),
			"BudgetMoney.Maintenance", int(b.Conf.BudgetMoney.Maintenance))

		utils.DoClickElement(ctx, model.BUTTON_HUBS_HUB_MANAGE_REPAIR)

		// reduce current account money and maintenance budged by repair cost
		b.AccountBalance -= repairCost
		b.Conf.BudgetMoney.Maintenance -= repairCost
	}

	return nil
}

// BUTTON_HUBS_HUB_REPAIR
// TEXT_HUBS_HUB_MANAGE_REPAIR_COST
