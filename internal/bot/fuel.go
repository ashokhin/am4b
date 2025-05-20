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

var fuelList []model.Fuel = []model.Fuel{
	{
		FuelType: "fuel",
	},
	{
		FuelType: "co2",
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
		slog.Debug("processing fuel type", "type", fuelEntry.FuelType)

		if err := b.checkFuelType(ctx, &fuelEntry); err != nil {
			slog.Warn("error from Bot.fuel > checkFuelType", "type",
				fuelEntry.FuelType, "error", err)

			return err
		}

		slog.Debug("fuel collected", "type", fuelEntry.FuelType, "fuel", fuelEntry)

		if fuelEntry.IsFull {
			slog.Debug("fuel is full", "type", fuelEntry.FuelType)

			continue
		}

		if err := b.buyFuelType(ctx, &fuelEntry); err != nil {
			slog.Warn("error from Bot.fuel > buyFuelType", "type",
				fuelEntry.FuelType, "error", err)

			return err
		}
	}

	return nil
}

func (b *Bot) checkFuelType(ctx context.Context, fuelStruct *model.Fuel) error {
	slog.Debug("check fuel type", "type", fuelStruct.FuelType)

	switch fuelStruct.FuelType {
	case "fuel":
		utils.DoClickElement(ctx, model.BUTTON_COMMON_TAB1)
	case "co2":
		utils.DoClickElement(ctx, model.BUTTON_COMMON_TAB2)
	}

	if err := chromedp.Run(ctx,
		utils.GetFloatFromElement(model.TEXT_FUEL_FUEL_PRICE, &fuelStruct.Price),
		utils.GetFloatFromElement(model.TEXT_FUEL_FUEL_HOLDING, &fuelStruct.Holding),
		utils.GetFloatFromElement(model.TEXT_FUEL_FUEL_CAPACITY, &fuelStruct.Capacity),
	); err != nil {
		slog.Warn("error in Bot.checkFuelType", "type", fuelStruct.FuelType, "error", err)

		return err
	}

	slog.Debug("set prometheus metrics", "type", fuelStruct.FuelType, "fuel", *fuelStruct)

	b.PrometheusMetrics.FuelHolding.WithLabelValues(fuelStruct.FuelType).Set(fuelStruct.Holding)
	b.PrometheusMetrics.FuelLimit.WithLabelValues(fuelStruct.FuelType).Set(fuelStruct.Capacity)
	b.PrometheusMetrics.FuelPrice.WithLabelValues(fuelStruct.FuelType).Set(fuelStruct.Price)

	// calc how much fuel do we need to max.capacity and compare it with FUEL_MINIMUM_AMOUNT
	// needAmount = capacity - holding
	// if needAmount less than FUEL_MINIMUM_AMOUNT then isFull = true
	fuelStruct.IsFull = (fuelStruct.Capacity - fuelStruct.Holding) < FUEL_MINIMUM_AMOUNT

	return nil
}

func (b *Bot) buyFuelType(ctx context.Context, fuelStruct *model.Fuel) error {
	var fuelExpectedPrice float64

	slog.Debug("buy fuel type", "type", fuelStruct.FuelType)

	switch fuelStruct.FuelType {
	case "fuel":
		utils.DoClickElement(ctx, model.BUTTON_COMMON_TAB1)
		fuelExpectedPrice = b.Conf.FuelPrice.Fuel
	case "co2":
		utils.DoClickElement(ctx, model.BUTTON_COMMON_TAB2)
		fuelExpectedPrice = b.Conf.FuelPrice.Co2
	}

	fuelNeedAmount := fuelStruct.Capacity - fuelStruct.Holding
	fuelKeepAmountPercent := (fuelStruct.Holding / fuelStruct.Capacity) * 100
	// price per 1000 Lbs/Quotas
	amountPrice := (fuelNeedAmount * fuelStruct.Price) / 1000

	// if fuel less than critical_percent then buy fuel anyway
	if fuelKeepAmountPercent <= b.Conf.FuelCriticalPercent {
		slog.Info("not enough fuel (less than fuel_critical_percent)", "type", fuelStruct.FuelType,
			"keepPercent", int(fuelKeepAmountPercent),
			"critical_percent", int(b.Conf.FuelCriticalPercent))
	} else if fuelStruct.Price > fuelExpectedPrice { // else if fuelPrice more that expectedPrice then exit
		slog.Info("fuel is too expensive", "type", fuelStruct.FuelType, "price", int(fuelStruct.Price),
			"expected", int(fuelExpectedPrice))

		return nil
	} else if amountPrice > b.Conf.BudgetMoney.Fuel { // else if amountPrice more than budget then exit
		slog.Info("not enough money for buying fuel", "type", fuelStruct.FuelType, "need", int(amountPrice),
			"budget", int(b.Conf.BudgetMoney.Fuel))

		return nil
	}

	fuelNeedAmountString := fmt.Sprintf("%d", int(fuelNeedAmount))

	slog.Debug("buying fuel", "type", fuelStruct.FuelType, "amount", fuelNeedAmountString, "price", int(amountPrice))

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(model.TEXT_FIELD_FUEL_AMOUNT, fuelNeedAmountString, chromedp.ByQuery),
		utils.ClickElement(model.BUTTON_FUEL_BUY),
	); err != nil {
		slog.Warn("error in Bot.buyFuelType", "type", fuelStruct.FuelType, "error", err)

		return err
	}

	slog.Debug("money before", "AccountBalance", int(b.AccountBalance), "fuelBudget", int(b.Conf.BudgetMoney.Fuel))
	b.AccountBalance -= amountPrice
	b.Conf.BudgetMoney.Fuel -= amountPrice
	slog.Debug("money after", "AccountBalance", int(b.AccountBalance), "fuelBudget", int(b.Conf.BudgetMoney.Fuel))

	return nil
}
