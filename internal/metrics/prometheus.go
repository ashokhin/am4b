package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "am4"
)

type Metrics struct {
	Up                    prometheus.Gauge
	StartTime             prometheus.Gauge
	Duration              prometheus.Gauge
	CompanyRank           prometheus.Gauge
	CompanyTrainingPoints prometheus.Gauge
	CompanyFleetSize      prometheus.Gauge
	RoutesNumber          prometheus.Gauge
	HubsNumber            prometheus.Gauge
	HangarCapacity        prometheus.Gauge
	SharePrice            prometheus.Gauge
	FlightsOperated       prometheus.Gauge
	PassengersTransported *prometheus.GaugeVec
	CargoTransported      *prometheus.GaugeVec
	AircraftStatus        *prometheus.GaugeVec
	CompanyReputation     *prometheus.GaugeVec
	CompanyMoney          *prometheus.GaugeVec
	HubStats              *prometheus.GaugeVec
	StaffSalary           *prometheus.GaugeVec
	FuelHolding           *prometheus.GaugeVec
	FuelLimit             *prometheus.GaugeVec
	FuelPrice             *prometheus.GaugeVec
}

func New() *Metrics {
	return &Metrics{
		Up: prometheus.NewGauge(
			prometheus.GaugeOpts{
				// Namespace: namespace,
				Name: "up",
				Help: "Was the last execution successful.",
			},
		),
		StartTime: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "process_start_time_seconds",
				Help:      "Start time of the process since unix epoch in seconds.",
			},
		),
		Duration: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "duration_seconds",
				Help:      "Duration of execution in seconds.",
			},
		),

		CompanyRank: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "company_rank",
				Help:      "Company rank value.",
			},
		),
		CompanyTrainingPoints: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "company_training_points",
				Help:      "Company training points value.",
			},
		),
		CompanyFleetSize: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "ac_fleet_size",
				Help:      "Company fleet size value.",
			},
		),
		RoutesNumber: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "ac_routes",
				Help:      "Company routes number value.",
			},
		),
		HubsNumber: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "company_hubs",
				Help:      "Company hubs number value.",
			},
		),
		HangarCapacity: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "ac_hangar_capacity",
				Help:      "Company hangar capacity value.",
			},
		),
		SharePrice: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "company_share_value",
				Help:      "Company share price value.",
			},
		),
		FlightsOperated: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "stats_flights_operated",
				Help:      "Company flights operated value.",
			},
		),
		PassengersTransported: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "stats_passengers_transported",
				Help:      "Passengers transported by type.",
			},
			[]string{"type"},
		),
		CargoTransported: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "stats_cargo_transported",
				Help:      "Cargo transported by type.",
			},
			[]string{"type"},
		),
		AircraftStatus: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "ac_status",
				Help:      "Aircraft status by type.",
			},
			[]string{"type"},
		),
		CompanyReputation: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "company_reputation",
				Help:      "Company reputation by company type.",
			},
			[]string{"type"},
		),
		CompanyMoney: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "company_money",
				Help:      "Company money by account type.",
			},
			[]string{"type"},
		),
		HubStats: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "hub_stats",
				Help:      "Company hub info by hub name and stat type.",
			},
			[]string{"name", "type"},
		),
		StaffSalary: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "company_staff_salary",
				Help:      "Company staff salary by staff type.",
			},
			[]string{"type"},
		),
		FuelHolding: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "company_fuel_holding",
				Help:      "Fuel amount holding by fuel type.",
			},
			[]string{"type"},
		),
		FuelLimit: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "company_fuel_limit",
				Help:      "Fuel amount limit by fuel type.",
			},
			[]string{"type"},
		),
		FuelPrice: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "market_fuel_price",
				Help:      "Fuel amount price by fuel type.",
			},
			[]string{"type"},
		),
	}
}

func (m *Metrics) RegisterMetrics(registry *prometheus.Registry) {
	registry.MustRegister(
		m.Up,
		m.StartTime,
		m.Duration,
		m.CompanyRank,
		m.CompanyTrainingPoints,
		m.CompanyFleetSize,
		m.RoutesNumber,
		m.HubsNumber,
		m.HangarCapacity,
		m.SharePrice,
		m.FlightsOperated,
		m.PassengersTransported,
		m.CargoTransported,
		m.AircraftStatus,
		m.CompanyReputation,
		m.CompanyMoney,
		m.HubStats,
		m.StaffSalary,
		m.FuelHolding,
		m.FuelLimit,
		m.FuelPrice,
	)
}
