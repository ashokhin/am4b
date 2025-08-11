package bot

import (
	"context"
	"log/slog"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/ashokhin/am4bot/internal/utils"
	"github.com/chromedp/chromedp"
)

func (b *Bot) companyStats(ctx context.Context) error {
	var (
		airlineReputation             float64
		cargoReputation               float64
		fleetSize                     float64
		acInflight                    float64
		acPendingMaintenance          float64
		acPendingDelivery             float64
		routes                        float64
		hubs                          float64
		hangarCapacity                float64
		sharePrice                    float64
		flightsOperated               float64
		passengersEconomyTransported  float64
		passengersBusinessTransported float64
		passengersFirstTransported    float64
		cargoTransportedLarge         float64
		cargoTransportedHeavy         float64
	)

	slog.Info("check company stats")
	slog.Debug("open pop-up window", "window", "overview")

	if err := chromedp.Run(ctx,
		chromedp.Click(model.BUTTON_FI_OVERVIEW, chromedp.ByQuery),
		chromedp.WaitReady(model.TEXT_OVERVIEW_AIRLINE_REPUTATION, chromedp.ByQuery),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_AIRLINE_REPUTATION, &airlineReputation),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_CARGO_REPUTATION, &cargoReputation),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_FLEET_SIZE, &fleetSize),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_AC_PENDING_DELIVERY, &acPendingDelivery),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_ROUTES, &routes),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_HUBS, &hubs),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_AC_PENDING_MAINTENANCE, &acPendingMaintenance),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_HANGAR_CAPACITY, &hangarCapacity),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_AC_INFLIGHT, &acInflight),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_SHARE_PRICE, &sharePrice),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_FLIGHTS_OPERATED, &flightsOperated),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_PASSENGERS_ECONOMY_TRANSPORTED, &passengersEconomyTransported),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_PASSENGERS_BUSINESS_TRANSPORTED, &passengersBusinessTransported),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_PASSENGERS_FIRST_TRANSPORTED, &passengersFirstTransported),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_CARGO_TRANSPORTED_LARGE, &cargoTransportedLarge),
		utils.GetFloatFromElement(model.TEXT_OVERVIEW_CARGO_TRANSPORTED_HEAVY, &cargoTransportedHeavy),
		chromedp.Click(model.BUTTON_COMMON_CLOSE_POPUP, chromedp.ByQuery),
		//utils.Screenshot(),
	); err != nil {
		slog.Debug("error in Bot.companyStats", "error", err)

		return err
	}

	acWithoutRoute := (fleetSize - (acPendingDelivery + routes))

	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.CompanyReputation.WithLabelValues("airline"), airlineReputation)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.CompanyReputation.WithLabelValues("cargo"), cargoReputation)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.CompanyFleetSize, fleetSize)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.AircraftStatus.WithLabelValues("in_flight"), acInflight)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.AircraftStatus.WithLabelValues("pending_delivery"), acPendingDelivery)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.AircraftStatus.WithLabelValues("pending_maintenance"), acPendingMaintenance)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.AircraftStatus.WithLabelValues("wo_route"), acWithoutRoute)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.RoutesNumber, routes)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.HubsNumber, hubs)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.HangarCapacity, hangarCapacity)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.SharePrice, sharePrice)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.FlightsOperated, flightsOperated)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.PassengersTransported.WithLabelValues("economy"), passengersEconomyTransported)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.PassengersTransported.WithLabelValues("business"), passengersBusinessTransported)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.PassengersTransported.WithLabelValues("first"), passengersFirstTransported)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.CargoTransported.WithLabelValues("large"), cargoTransportedLarge)
	utils.SetPromGaugeNonNeg(b.PrometheusMetrics.CargoTransported.WithLabelValues("heavy"), cargoTransportedHeavy)

	return nil
}
