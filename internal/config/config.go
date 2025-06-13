package config

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

// url: "https://www.airlinemanager.com/"
// username: "z1odeypnd@gmail.com"
// password: "tW+L-L9Xp7@,MFR"
// budget_percent:
//   maintenance: 50
//   marketing: 70
//   fuel: 70
// good_price:
//   fuel: 500
//   co2: 120
// aircraft_wear_percent: 80
// aircraft_max_hours_to_check: 24
// service_cron_string: "*/5 * * * *"
// services:
//   - "check_staff_morale"
//   - "check_lounges"

type Config struct {
	Url                     string `yaml:"url"`
	User                    string `yaml:"username"`
	Password                string `yaml:"password"`
	LogLevel                string `yaml:"log_level"`
	BudgetPercent           Budget `yaml:"budget_percent"`
	BudgetMoney             Budget
	FuelPrice               Price    `yaml:"good_price"`
	AircraftWearPercent     float64  `yaml:"aircraft_wear_percent"`
	AircraftMaxHoursToCheck float64  `yaml:"aircraft_max_hours_to_check"`
	AircraftModifyLimit     float64  `yaml:"aircraft_modify_limit"`
	FuelCriticalPercent     float64  `yaml:"fuel_critical_percent"`
	CronSchedule            string   `yaml:"service_cron_string"`
	Services                []string `yaml:"services"`
	passwordRunes           []rune   // most safe storage for password in memory
}

type Budget struct {
	Maintenance float64 `yaml:"maintenance"`
	Marketing   float64 `yaml:"marketing"`
	Fuel        float64 `yaml:"fuel"`
}

type Price struct {
	Fuel float64 `yaml:"fuel"`
	Co2  float64 `yaml:"co2"`
}

func (c *Config) String() string {
	return fmt.Sprintf("%+v", *c)
}

// convert password string into array of runes
func (c *Config) SafeStorePassword() {
	c.passwordRunes = []rune(c.Password)
	c.Password = ""
}

// getter for returning password as a string
func (c *Config) GetPassword() string {
	return string(c.passwordRunes)
}

func New(filePath string) (*Config, error) {
	var err error
	var c Config

	slog.Info("loading config file", "file", filePath)

	if err := loadYaml(filePath, &c); err != nil {
		return &c, err
	}

	c.SafeStorePassword()

	slog.Debug("yaml loaded", "yaml", c)

	return &c, err
}

func loadYaml(filePath string, out interface{}) error {
	var err error
	var f []byte

	slog.Debug("read file", "file", filePath)

	if f, err = os.ReadFile(filePath); err != nil {
		return err
	}

	slog.Debug("load file as yaml", "file", filePath)

	if err := yaml.Unmarshal(f, out); err != nil {
		return err
	}

	return err
}
