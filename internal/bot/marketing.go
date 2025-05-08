package bot

import (
	"context"
	"log/slog"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/ashokhin/am4bot/internal/utils"
)

type marketingCompany struct {
	name          string
	companyRow    string
	companyButton string
}

var marketingCompaniesList []marketingCompany = []marketingCompany{
	{
		name:          "Airline reputation",
		companyRow:    "",
		companyButton: "",
	},
	{
		name:          "Eco friendly",
		companyRow:    "",
		companyButton: "",
	},
	{
		name:          "Cargo reputation",
		companyRow:    "",
		companyButton: "",
	},
}

func (b *Bot) marketingCompanies(ctx context.Context) error {
	slog.Info("check marketing companies")

	// open finance pop-up
	utils.DoClickElement(ctx, model.BUTTON_MAIN_FINANCE)
	// to to marketing tab
	utils.DoClickElement(ctx, model.BUTTON_COMMON_TAB2)
	defer utils.DoClickElement(ctx, model.BUTTON_COMMON_CLOSE_POPUP)

	for _, markComp := range marketingCompaniesList {
		slog.Warn("marketing company", "company", markComp.name)
	}

	return nil
}
