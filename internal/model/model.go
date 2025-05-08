package model

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
