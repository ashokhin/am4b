package bot

import (
	"context"
	"log/slog"
	"strings"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/ashokhin/am4bot/internal/utils"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func (b *Bot) Money(ctx context.Context) error {
	var accElemList []*cdp.Node

	slog.Info("check account money")
	slog.Debug("get accounts list")

	if err := chromedp.Run(ctx,
		utils.ClickElement(model.BUTTON_MAIN_ACCOUNT),
		chromedp.Nodes(model.LIST_ACCOUNT_ACCOUNTS, &accElemList, chromedp.ByQueryAll),
	); err != nil {
		slog.Warn("error in bot.Money > get accounts list", "error", err)

		return err
	}

	defer utils.DoClickElement(ctx, model.BUTTON_COMMON_CLOSE_POPUP)

	for idx, accountElem := range accElemList {
		var accountName string
		var accountBalance float64

		slog.Debug("check account", "index", idx)

		if err := chromedp.Run(ctx,
			chromedp.Text(model.TEXT_ACCOUNT_ACCOUNT_NAME, &accountName, chromedp.ByQuery, chromedp.FromNode(accountElem)),
			utils.GetFloatFromChildElement(model.TEXT_ACCOUNT_ACCOUNT_BALANCE, &accountBalance, accountElem),
		); err != nil {
			slog.Warn("error in Bot.Money > get account info", "error", err)

			return err
		}

		accountName = strings.TrimSpace(accountName)

		slog.Debug("account balance", "name", accountName, "value", accountBalance)

		b.PrometheusMetrics.CompanyMoney.WithLabelValues(accountName).Set(accountBalance)

		if accountName == "Airline account" {
			b.AccountBalance = accountBalance
		}
	}

	b.calcBudget()

	return nil
}

func (b *Bot) calcBudget() {
	slog.Debug("calculate budgets")

	b.Conf.BudgetMoney.Maintenance = (b.AccountBalance * (b.Conf.BudgetPercent.Maintenance * 0.01))
	b.Conf.BudgetMoney.Marketing = (b.AccountBalance * (b.Conf.BudgetPercent.Marketing * 0.01))
	b.Conf.BudgetMoney.Fuel = (b.AccountBalance * (b.Conf.BudgetPercent.Fuel * 0.01))

	slog.Debug("calculated budget",
		"maintenancePercent", b.Conf.BudgetPercent.Maintenance,
		"maintenanceMoney", int(b.Conf.BudgetMoney.Maintenance),
		"marketingPercent", b.Conf.BudgetPercent.Marketing,
		"marketingMoney", int(b.Conf.BudgetMoney.Marketing),
		"fuelPercent", int(b.Conf.BudgetPercent.Fuel),
		"fuelMoney", int(b.Conf.BudgetMoney.Fuel))
}
