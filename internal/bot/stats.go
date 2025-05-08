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
		slog.Debug("error in Bot.staffMorale", "error", err)

		return err
	}

	acWithoutRoute := (fleetSize - (acPendingDelivery + routes))

	b.PrometheusMetrics.CompanyReputation.WithLabelValues("airline").Set(airlineReputation)
	b.PrometheusMetrics.CompanyReputation.WithLabelValues("cargo").Set(cargoReputation)
	b.PrometheusMetrics.CompanyFleetSize.Set(fleetSize)
	b.PrometheusMetrics.AircraftStatus.WithLabelValues("in_flight").Set(acInflight)
	b.PrometheusMetrics.AircraftStatus.WithLabelValues("pending_delivery").Set(acPendingDelivery)
	b.PrometheusMetrics.AircraftStatus.WithLabelValues("pending_maintenance").Set(acPendingMaintenance)
	b.PrometheusMetrics.AircraftStatus.WithLabelValues("wo_route").Set(acWithoutRoute)
	b.PrometheusMetrics.RoutesNumber.Set(routes)
	b.PrometheusMetrics.HubsNumber.Set(hubs)
	b.PrometheusMetrics.HangarCapacity.Set(hangarCapacity)
	b.PrometheusMetrics.SharePrice.Set(sharePrice)
	b.PrometheusMetrics.FlightsOperated.Set(flightsOperated)
	b.PrometheusMetrics.PassengersTransported.WithLabelValues("economy").Set(passengersEconomyTransported)
	b.PrometheusMetrics.PassengersTransported.WithLabelValues("business").Set(passengersBusinessTransported)
	b.PrometheusMetrics.PassengersTransported.WithLabelValues("first").Set(passengersFirstTransported)
	b.PrometheusMetrics.CargoTransported.WithLabelValues("large").Set(cargoTransportedLarge)
	b.PrometheusMetrics.CargoTransported.WithLabelValues("heavy").Set(cargoTransportedHeavy)

	return nil
}
