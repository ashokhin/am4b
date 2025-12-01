package model

import (
	"fmt"

	"github.com/chromedp/cdproto/cdp"
)

type MaintenanceType int

const (
	A_CHECK MaintenanceType = iota
	REPAIR
	MODIFY
)

// StaffEntry represents a staff category with associated UI elements for salary and morale management.
type StaffEntry struct {
	Name             string
	TextSalary       string
	TextMorale       string
	ButtonSalaryUp   string
	ButtonSalaryDown string
}

var StaffEntires = []StaffEntry{
	{
		"pilots",
		"#pilotSalary",
		"#pilotMorale",
		`button[onclick="Ajax('staff_action.php?type=pilot&mode=raise','runme',this);"]`,
		`button[onclick="Ajax('staff_action.php?type=pilot&mode=cut','runme',this);"]`,
	},
	{
		"crew",
		"#crewSalary",
		"#crewMorale",
		`button[onclick="Ajax('staff_action.php?type=crew&mode=raise','runme',this);"]`,
		`button[onclick="Ajax('staff_action.php?type=crew&mode=cut','runme',this);"]`,
	},
	{
		"engineers",
		"#engineerSalary",
		"#engineerMorale",
		`button[onclick="Ajax('staff_action.php?type=engineer&mode=raise','runme',this);"]`,
		`button[onclick="Ajax('staff_action.php?type=engineer&mode=cut','runme',this);"]`,
	},
	{
		"technicians",
		"#techSalary",
		"#techMorale",
		`button[onclick="Ajax('staff_action.php?type=tech&mode=raise','runme',this);"]`,
		`button[onclick="Ajax('staff_action.php?type=tech&mode=cut','runme',this);"]`,
	},
}

// Fuel represents fuel information for an aircraft.
type Fuel struct {
	FuelType string
	Price    float64
	Holding  float64
	Capacity float64
	IsFull   bool
}

// Aircraft represents an aircraft in the fleet.
type Aircraft struct {
	RegNumber   string
	AcType      string
	WearPercent float64
	HoursACheck int
}

// MarketingCompany represents a marketing company with associated UI elements for activation and cost.
type MarketingCompany struct {
	Name               string
	CompanyRow         string
	CompanyOptionValue string
	CompanyCost        string
	CompanyButton      string
}

// Hub represents an airport hub with various statistics.
type Hub struct {
	Departures    float64
	Arrivals      float64
	PaxDeparted   float64
	PaxArrived    float64
	HasCatering   bool
	NeedsRepair   bool
	HubCdpNode    *cdp.Node
	LoungeCdpNode *cdp.Node
}

func (h Hub) String() string {
	return fmt.Sprint("{Departures:", h.Departures, ", Arrivals:", h.Arrivals,
		", PaxDeparted:", h.PaxDeparted, ", PaxArrived:", h.PaxArrived,
		", HasCatering:", h.HasCatering, ", NeedsRepair:", h.NeedsRepair, "}")
}
