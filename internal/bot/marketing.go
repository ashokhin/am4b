package bot

import (
	"context"
	"log/slog"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/ashokhin/am4bot/internal/utils"
	"github.com/chromedp/chromedp"
)

var marketingCompaniesList = []model.MarketingCompany{
	{
		Name:               "Airline reputation",
		CompanyRow:         model.ELEM_FINANCE_MARKETING_INC_AIRLINE_REP,
		CompanyOptionValue: model.OPTION_FINANCE_MARKETING_INC_AIRLINE_REP_24H_VALUE,
		CompanyCost:        model.TEXT_FINANCE_MARKETING_INC_AIRLINE_REP_COST,
		CompanyButton:      model.BUTTON_FINANCE_MARKETING_INC_AIRLINE_REP_BUY,
	},
	{
		Name:               "Cargo reputation",
		CompanyRow:         model.ELEM_FINANCE_MARKETING_INC_CARGO_REP,
		CompanyOptionValue: model.OPTION_FINANCE_MARKETING_INC_CARGO_REP_24H_VALUE,
		CompanyCost:        model.TEXT_FINANCE_MARKETING_INC_CARGO_REP_COST,
		CompanyButton:      model.BUTTON_FINANCE_MARKETING_INC_CARGO_REP_BUY,
	},
	{
		Name:               "Eco friendly",
		CompanyRow:         model.ELEM_FINANCE_MARKETING_ECO_FRIENDLY,
		CompanyOptionValue: "",
		CompanyCost:        model.TEXT_FINANCE_MARKETING_ECO_FRIENDLY_COST,
		CompanyButton:      model.BUTTON_FINANCE_MARKETING_ECO_FRIENDLY_BUY,
	},
}

// marketingCompanies checks and activates marketing companies based on budget and status.
func (b *Bot) marketingCompanies(ctx context.Context) error {
	slog.Info("check marketing companies")

	// open finance pop-up
	utils.DoClickElement(ctx, model.BUTTON_MAIN_FINANCE)
	defer utils.DoClickElement(ctx, model.BUTTON_COMMON_CLOSE_POPUP)

	for _, markComp := range marketingCompaniesList {
		slog.Debug("marketing company", "company", markComp.Name)

		if err := b.activateMarketingCompany(ctx, markComp); err != nil {
			slog.Warn("error in Bot.marketingCompanies > Bot.activateMarketingCompany", "company", markComp.Name, "error", err)

			return err
		}
	}

	return nil
}

// activateMarketingCompany activates a specific marketing company if it is affordable and not already active.
func (b *Bot) activateMarketingCompany(ctx context.Context, mc model.MarketingCompany) error {
	var marketingCompanyElemAttributes map[string]string

	slog.Debug("Check marketing company", "company", mc.Name)

	// search marketingCompany element attributes
	if err := chromedp.Run(ctx,
		utils.ClickElement(model.BUTTON_COMMON_TAB2),
		utils.ClickElement(model.BUTTON_FINANCE_MARKETING_NEW_COMPANY),
		chromedp.Attributes(mc.CompanyRow, &marketingCompanyElemAttributes, chromedp.ByQuery),
	); err != nil {
		slog.Warn("error in Bot.activateMarketingCompany > get company elem attributes", "company", mc.Name, "error", err)

		return err
	}

	slog.Debug("attributes found", "company", mc.Name, "attributes", marketingCompanyElemAttributes)

	if marketingCompanyElemAttributes["class"] == "not-active" {
		slog.Debug("marketing company inactive", "company", mc.Name)

		return nil
	}

	utils.DoClickElement(ctx, mc.CompanyRow)

	var companyCost float64
	// get marketing company cost
	switch mc.Name {
	// in case of "Eco friendly" marketing company we skip "select option" actions
	case "Eco friendly":
		if err := chromedp.Run(ctx,
			utils.GetFloatFromElement(mc.CompanyCost, &companyCost),
		); err != nil {
			slog.Warn("error in Bot.activateMarketingCompany > get company cost", "company", mc.Name, "error", err)

			return err
		}
	default:
		if err := chromedp.Run(ctx,
			chromedp.SetValue(model.SELECT_FINANCE_MARKETING_COMPANY_DURATION, mc.CompanyOptionValue, chromedp.ByQuery),
			utils.GetFloatFromElement(mc.CompanyCost, &companyCost),
		); err != nil {
			slog.Warn("error in Bot.activateMarketingCompany > get company cost", "company", mc.Name, "error", err)

			return err
		}
	}

	slog.Debug("company cost", "company", mc.Name, "cost", int(companyCost))

	if companyCost > b.Conf.BudgetMoney.Marketing {
		slog.Warn("marketing company is too expensive", "company", mc.Name,
			"cost", int(companyCost), "budget", int(b.Conf.BudgetMoney.Marketing))

		return nil
	}

	// buy marketing company
	if err := chromedp.Run(ctx,
		utils.ClickElement(mc.CompanyButton),
	); err != nil {
		slog.Warn("error in Bot.activateMarketingCompany > buy company", "company", mc.Name, "error", err)

		return err
	}

	b.Conf.BudgetMoney.Marketing -= companyCost
	b.AccountBalance -= companyCost

	return nil
}
