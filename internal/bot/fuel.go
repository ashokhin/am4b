package bot

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/ashokhin/am4bot/internal/utils"
	"github.com/chromedp/chromedp"
)

const FUEL_MINIMUM_AMOUNT float64 = 1000.00

type fuel struct {
	// fuel type
	fType    string
	price    float64
	holding  float64
	capacity float64
	isFull   bool
}

var fuelList []fuel = []fuel{
	{
		fType: "fuel",
	},
	{
		fType: "co2",
	},
}

// fuel is checking fuel levels and buying them if price is good
// or levels are critical
func (b *Bot) fuel(ctx context.Context) error {
	slog.Info("check fuel")

	// open fuel window
	utils.DoClickElement(ctx, model.BUTTON_MAIN_FUEL)
	defer utils.DoClickElement(ctx, model.BUTTON_COMMON_CLOSE_POPUP)

	for _, fuelEntry := range fuelList {
		slog.Debug("processing fuel type", "type", fuelEntry.fType)

		if err := b.checkFuelType(ctx, &fuelEntry); err != nil {
			slog.Warn("error from Bot.fuel > checkFuelType", "type",
				fuelEntry.fType, "error", err)

			return err
		}

		slog.Debug("fuel collected", "type", fuelEntry.fType, "fuel", fuelEntry)

		if fuelEntry.isFull {
			slog.Debug("fuel is full", "type", fuelEntry.fType)

			continue
		}

		if err := b.buyFuelType(ctx, &fuelEntry); err != nil {
			slog.Warn("error from Bot.fuel > buyFuelType", "type",
				fuelEntry.fType, "error", err)

			return err
		}
	}

	return nil
}

func (b *Bot) checkFuelType(ctx context.Context, fuelStruct *fuel) error {
	slog.Debug("check fuel type", "type", fuelStruct.fType)

	switch fuelStruct.fType {
	case "fuel":
		utils.DoClickElement(ctx, model.BUTTON_COMMON_TAB1)
	case "co2":
		utils.DoClickElement(ctx, model.BUTTON_COMMON_TAB2)
	}

	if err := chromedp.Run(ctx,
		utils.GetFloatFromElement(model.TEXT_FUEL_FUEL_PRICE, &fuelStruct.price),
		utils.GetFloatFromElement(model.TEXT_FUEL_FUEL_HOLDING, &fuelStruct.holding),
		utils.GetFloatFromElement(model.TEXT_FUEL_FUEL_CAPACITY, &fuelStruct.capacity),
	); err != nil {
		slog.Warn("error in Bot.checkFuelType", "type", fuelStruct.fType, "error", err)

		return err
	}

	slog.Debug("set prometheus metrics", "type", fuelStruct.fType, "fuel", *fuelStruct)

	b.PrometheusMetrics.FuelHolding.WithLabelValues(fuelStruct.fType).Set(fuelStruct.holding)
	b.PrometheusMetrics.FuelLimit.WithLabelValues(fuelStruct.fType).Set(fuelStruct.capacity)
	b.PrometheusMetrics.FuelPrice.WithLabelValues(fuelStruct.fType).Set(fuelStruct.price)

	// calc how much fuel do we need to max.capacity and compare it with FUEL_MINIMUM_AMOUNT
	// needAmount = capacity - holding
	// if needAmount less than FUEL_MINIMUM_AMOUNT then isFull = true
	fuelStruct.isFull = (fuelStruct.capacity - fuelStruct.holding) < FUEL_MINIMUM_AMOUNT

	return nil
}

func (b *Bot) buyFuelType(ctx context.Context, fuelStruct *fuel) error {
	var fuelExpectedPrice float64

	slog.Debug("buy fuel type", "type", fuelStruct.fType)

	switch fuelStruct.fType {
	case "fuel":
		utils.DoClickElement(ctx, model.BUTTON_COMMON_TAB1)
		fuelExpectedPrice = b.Conf.FuelPrice.Fuel
	case "co2":
		utils.DoClickElement(ctx, model.BUTTON_COMMON_TAB2)
		fuelExpectedPrice = b.Conf.FuelPrice.Co2
	}

	fuelNeedAmount := fuelStruct.capacity - fuelStruct.holding
	fuelKeepAmountPercent := (fuelStruct.holding * fuelStruct.capacity) * 100
	// price per 1000 Lbs/Quotas
	amountPrice := (fuelNeedAmount * fuelStruct.price) / 1000

	// if fuel less than critical_percent then buy fuel anyway
	if fuelKeepAmountPercent <= b.Conf.FuelCriticalPercent {
		slog.Info("not enough fuel (less than fuel_critical_percent)", "type", fuelStruct.fType,
			"keepPercent", int(fuelKeepAmountPercent),
			"critical_percent", int(b.Conf.FuelCriticalPercent))
	} else if fuelStruct.price > fuelExpectedPrice { // else if fuelPrice more that expectedPrice then exit
		slog.Info("fuel is too expensive", "type", fuelStruct.fType, "price", int(fuelStruct.price),
			"expected", int(fuelExpectedPrice))

		return nil
	} else if amountPrice > b.Conf.BudgetMoney.Fuel { // else if amountPrice more than budget then exit
		slog.Info("not enough money for buying fuel", "type", fuelStruct.fType, "need", int(amountPrice),
			"budget", int(b.Conf.BudgetMoney.Fuel))

		return nil
	}

	fuelNeedAmountString := fmt.Sprintf("%d", int(fuelNeedAmount))

	slog.Debug("buying fuel", "type", fuelStruct.fType, "amount", fuelNeedAmountString, "price", int(amountPrice))

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(model.TEXT_FIELD_FUEL_AMOUNT, fuelNeedAmountString, chromedp.ByQuery),
		chromedp.Click(model.BUTTON_FUEL_BUY, chromedp.ByQuery),
	); err != nil {
		slog.Warn("error in Bot.buyFuelType", "type", fuelStruct.fType, "error", err)

		return err
	}

	slog.Debug("money before", "AccountBalance", int(b.AccountBalance), "fuelBudget", int(b.Conf.BudgetMoney.Fuel))
	b.AccountBalance -= amountPrice
	b.Conf.BudgetMoney.Fuel -= amountPrice
	slog.Debug("money after", "AccountBalance", int(b.AccountBalance), "fuelBudget", int(b.Conf.BudgetMoney.Fuel))

	return nil
}
