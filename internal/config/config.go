package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Url                     string   `default:"https://www.airlinemanager.com/" yaml:"url"`
	User                    string   `yaml:"username"`
	Password                string   `yaml:"password"`
	LogLevel                string   `default:"info" yaml:"log_level"`
	BudgetPercent           Budget   `yaml:"budget_percent"`
	FuelPrice               Price    `yaml:"good_price"`
	AircraftWearPercent     float64  `default:"80" yaml:"aircraft_wear_percent"`
	AircraftMaxHoursToCheck float64  `default:"24" yaml:"aircraft_max_hours_to_check"`
	AircraftModifyLimit     float64  `default:"3" yaml:"aircraft_modify_limit"`
	FuelCriticalPercent     float64  `default:"20" yaml:"fuel_critical_percent"`
	CronSchedule            string   `default:"*/5 * * * *" yaml:"service_cron_string"`
	Services                []string `default:"[\"company_stats\",\"staff_morale\",\"hubs\",\"buy_fuel\",\"marketing_companies\",\"ac_maintenance\",\"depart\"]" yaml:"services"`
	BudgetMoney             Budget
	passwordRunes           []rune // most safe storage for password in memory
}

type Budget struct {
	Maintenance float64 `default:"50" yaml:"maintenance"`
	Marketing   float64 `default:"70" yaml:"marketing"`
	Fuel        float64 `default:"70" yaml:"fuel"`
}

type Price struct {
	Fuel float64 `default:"500" yaml:"fuel"`
	Co2  float64 `default:"120" yaml:"co2"`
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

	defaults.Set(&c)

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
