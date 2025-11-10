![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ashokhin/am4b)
![GitHub issues](https://img.shields.io/github/issues/ashokhin/am4b)
![Docker Pulls](https://img.shields.io/docker/pulls/ashokhin/am4bot)

🐳[Docker Hub](https://hub.docker.com/r/ashokhin/am4bot)


# Airline Manager Bot

Automated bot for managing your Airline Manager company.

It can automatically handle various tasks such as starting marketing campaigns,
scheduling departures, purchasing fuel and CO2, improving staff morale,
managing hubs, collecting company statistics, performing repairs, A-Checks,
and modifications.

The bot is designed to run periodically based on a cron schedule, allowing for
continuous management of your airline without manual intervention.

It uses a headless browser to interact with the Airline Manager web interface,
simulating user actions to perform the necessary tasks.


## Features

- Automatic start marketing companies (Available: `Airline reputation`, `Cargo reputation`, `Eco friendly`).
- Automatic departures.
- Automatic buy Fuel when price is low or fuel level is critical.
- Automatic buy CO2 when price is low or quota level is critical.
- Automatic staff morale improvement.
- Automatic hub management.
- Automatic catering purchase.
- Automatic company statistics collection.
- Automatic alliance statistics collection.
- Automatic repair.
- Automatic A-Check.
- Automatic modification.
- Prometheus metrics support.


## Installation

1. Install Docker from https://www.docker.com/get-started
2. Create `config.yaml` file based on [Configuration](#configuration) section. For example:
   ```bash
   mkdir -p /opt/ambot/conf
   nano /opt/ambot/conf/config.yaml
   ```
   Paste your configuration and save the file.
3. Run the bot:
   ```bash
   docker run --rm --name ambot --volume /opt/ambot/conf/config.yaml:/config.yaml ashokhin/am4bot:latest
   ```
   
   For collecting Prometheus metrics, you can expose port 9150 (default in the config option `prometheus_address`) from container to host:
   ```bash
   docker run --rm --name ambot --volume /opt/ambot/conf/config.yaml:/config.yaml -p 9150:9150 ashokhin/am4bot:latest
   ```
4. (Optional) To run the bot as a [systemd service](https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html), create a file `/etc/systemd/system/am4bot.service` with the following content:
   ```ini
   [Unit]
   Description=Airline Manager bot
   Documentation="https://github.com/ashokhin/am4b"
   After=docker.service
   Requires=docker.service

   [Service]
   Type=simple
   Restart=always
   ExecStartPre=-/usr/bin/docker pull ashokhin/am4bot:latest
   ExecStart=/usr/bin/docker run --rm --name %n --volume /opt/ambot/conf/config.yaml:/config.yaml --publish 9150:9150 ashokhin/am4bot:latest

   [Install]
   WantedBy=multi-user.target
   ```
   
   Then enable and start the service:
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable am4bot.service --now
   ```


## Configuration

### Available options:
| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `url` | string | `"https://www.airlinemanager.com/"` | Airline Manager URL. |
| `username` | string | `""` | Username for login. |
| `password` | string | `""` | Password for login. |
| `log_level` | string | `"info"` | Logging level (debug, info, warn, error). |
| `budget_percent` | map of strings to int | see below | Percentage of budget to use for each category. |
| `budget_percent.fuel` | int | `70` | Percentage of budget for Fuel. |
| `budget_percent.maintenance` | int | `30` | Percentage of budget for Maintenance. |
| `budget_percent.marketing` | int | `70` | Percentage of budget for Marketing. |
| `good_price` | map of strings to int | see below | Good price thresholds for resources. |
| `good_price.fuel` | int | `500` | Good price for Fuel (per 1,000 Lbs). |
| `good_price.co2` | int | `120` | Good price for CO2 (per 1,000 Quotas). |
| `buy_catering_if_missing` | bool | `true` | Whether to buy catering if missing in hubs. |
| `catering_duration_hours` | string | `"168"` | Catering duration in hours to set when buying catering. Possible values: `6`, `12`, `18`, `24`, `48`, `72`, `96`, `120`, `144`, `168` |
| `catering_amount_option` | string | `"20000"` | Catering amount option to select when buying catering. Possible values: `200`, `500`, `1000`, `2000`, `3000`, `4000`, `5000`, `10000`, `15000`, `20000`, `50000`, `100000`, `200000` |
| `aircraft_wear_percent` | float | `80` | Aircraft wear percentage to trigger maintenance. |
| `aircraft_max_hours_to_check` | int | `24` | Max hours to next A-Check to trigger it. |
| `aircraft_modify_limit` | int | `3` | Max aircraft for modifications checks. |
| `fuel_critical_percent` | float | `20` | Fuel level percentage to trigger refuel. Even the price isn't good. |
| `cron_schedule` | string | `"*/5 * * * *"` | Cron schedule for services. Default: Every 5 minutes. |
| `services` | list of strings | `["company_stats",` `"staff_morale",` `"alliance_stats",` `"hubs",` `"buy_fuel",` `"depart",` `"marketing",` `"ac_maintenance"]` | List of services to run. Possible values: `company_stats`, `alliance_stats`, `staff_morale`, `hubs`, `buy_fuel`, `depart`, `marketing`, `ac_maintenance`. |
| `timeout_seconds` | int | `120` | Timeout for full round in seconds. |
| `chrome_headless` | bool | `true` | Run browser in headless mode. |
| `prometheus_address` | string | `":9150"` | Address to expose Prometheus metrics. |

#### Example of `config.yaml` with the non-default options:
```yaml
url: "https://www.airlinemanager.com/"
username: "your_email@example.com"
password: "YourPasswordHere"
log_level: "warn"
budget_percent:
  fuel: 75
  maintenance: 25
  marketing: 75
good_price:
  fuel: 550
  co2: 140
buy_catering_if_missing: true
catering_duration_hours: "24"
catering_amount_option: "5000"
aircraft_wear_percent: 75
aircraft_max_hours_to_check: 48
aircraft_modify_limit: 5
fuel_critical_percent: 15
cron_schedule: "*/10 * * * *"
services:
  - "company_stats"
  - "alliance_stats"
  - "staff_morale"
  - "hubs"
  - "buy_fuel"
  - "marketing"
  - "ac_maintenance"
  - "depart"
timeout_seconds: 180
# Not recommended to change this option
# on systems without GUI support
chrome_headless: true
prometheus_address: ":9150"
```

#### Minimal configuration example:
```yaml
username: "username@email.example"
password: "your_password_here"
```

#### Service descriptions:
- `company_stats`: Collects and exposes company statistics as Prometheus metrics.
- `alliance_stats`: Collects and exposes alliance statistics as Prometheus metrics.
- `staff_morale`: Improves staff morale if below 100%.
- `hubs`: Manages hubs, including buying catering if missing.
- `buy_fuel`: Buys fuel and CO2 based on good price thresholds and critical levels.
- `marketing`: Starts marketing campaigns based on budget percentage.
- `ac_maintenance`: Performs aircraft maintenance, A-Checks, and modifications based on configured thresholds.
- `depart`: Schedules departures for flights ready to depart.

> [!NOTE]
> 
> Note that the order of services in the configuration matters.
> All services are executed sequentially in the order they are listed.


## Prometheus Metrics

<details>
	<summary>Prometheus metrics example output</summary>

```
# HELP am4_ac_fleet_size Company fleet size value.
# TYPE am4_ac_fleet_size gauge
am4_ac_fleet_size 154
# HELP am4_ac_hangar_capacity Company hangar capacity value.
# TYPE am4_ac_hangar_capacity gauge
am4_ac_hangar_capacity 170
# HELP am4_ac_routes Company routes number value.
# TYPE am4_ac_routes gauge
am4_ac_routes 154
# HELP am4_ac_status Aircraft status by type.
# TYPE am4_ac_status gauge
am4_ac_status{type="in_flight"} 146
am4_ac_status{type="pending_delivery"} 0
am4_ac_status{type="pending_maintenance"} 10
am4_ac_status{type="wo_route"} 0
# HELP am4_alliance_contributed_per_day Alliance contributed per day value.
# TYPE am4_alliance_contributed_per_day gauge
am4_alliance_contributed_per_day 31318
# HELP am4_alliance_contributed_total Alliance contributed total value.
# TYPE am4_alliance_contributed_total gauge
am4_alliance_contributed_total 79152
# HELP am4_alliance_flights Alliance flights value.
# TYPE am4_alliance_flights gauge
am4_alliance_flights 479
# HELP am4_alliance_season_money Alliance season money value.
# TYPE am4_alliance_season_money gauge
am4_alliance_season_money 147
# HELP am4_build_info A metric with a constant '1' value labeled by version, revision, branch, goversion from which am4 was built, and the goos and goarch for the build.
# TYPE am4_build_info gauge
am4_build_info{branch="tags/1.28",goarch="amd64",goos="linux",goversion="go1.24.7",revision="84a34f6b9b1352cda47a517a8ee748c306c7d2e5",tags="unknown",version="1.28"} 1
# HELP am4_company_fuel_holding Fuel amount holding by fuel type.
# TYPE am4_company_fuel_holding gauge
am4_company_fuel_holding{type="co2"} 2.5868711e+07
am4_company_fuel_holding{type="fuel"} 2.0085064e+07
# HELP am4_company_fuel_limit Fuel amount limit by fuel type.
# TYPE am4_company_fuel_limit gauge
am4_company_fuel_limit{type="co2"} 2.85e+07
am4_company_fuel_limit{type="fuel"} 2.55e+07
# HELP am4_company_hubs Company hubs number value.
# TYPE am4_company_hubs gauge
am4_company_hubs 2
# HELP am4_company_money Company money by account type.
# TYPE am4_company_money gauge
am4_company_money{type="Airline account"} 3.030865737e+09
am4_company_money{type="Savings"} 1.102106418e+09
# HELP am4_company_rank Company rank value.
# TYPE am4_company_rank gauge
am4_company_rank 3051
# HELP am4_company_reputation Company reputation by company type.
# TYPE am4_company_reputation gauge
am4_company_reputation{type="airline"} 90
am4_company_reputation{type="cargo"} 88
# HELP am4_company_share_value Company share price value.
# TYPE am4_company_share_value gauge
am4_company_share_value 1547.3
# HELP am4_company_staff_salary Company staff salary by staff type.
# TYPE am4_company_staff_salary gauge
am4_company_staff_salary{type="crew"} 159
am4_company_staff_salary{type="engineers"} 264
am4_company_staff_salary{type="pilots"} 211
am4_company_staff_salary{type="technicians"} 236
# HELP am4_company_training_points Company training points value.
# TYPE am4_company_training_points gauge
am4_company_training_points 0
# HELP am4_duration_seconds Duration of execution in seconds.
# TYPE am4_duration_seconds gauge
am4_duration_seconds 76.877199534
# HELP am4_hub_stats Company hub info by hub name and stat type.
# TYPE am4_hub_stats gauge
am4_hub_stats{name="BRAZIL, BRASÍLIA",type="arrivals"} 4513
am4_hub_stats{name="BRAZIL, BRASÍLIA",type="departures"} 5951
am4_hub_stats{name="BRAZIL, BRASÍLIA",type="paxArrived"} 1.397119e+06
am4_hub_stats{name="BRAZIL, BRASÍLIA",type="paxDeparted"} 1.920807e+06
am4_hub_stats{name="UNITED STATES, NEW YORK JFK",type="arrivals"} 16269
am4_hub_stats{name="UNITED STATES, NEW YORK JFK",type="departures"} 16389
am4_hub_stats{name="UNITED STATES, NEW YORK JFK",type="paxArrived"} 3.759363e+06
am4_hub_stats{name="UNITED STATES, NEW YORK JFK",type="paxDeparted"} 3.825323e+06
# HELP am4_market_fuel_price Fuel amount price by fuel type.
# TYPE am4_market_fuel_price gauge
am4_market_fuel_price{type="co2"} 151
am4_market_fuel_price{type="fuel"} 1713
# HELP am4_stats_cargo_transported Cargo transported by type.
# TYPE am4_stats_cargo_transported gauge
am4_stats_cargo_transported{type="heavy"} 6.35925e+08
am4_stats_cargo_transported{type="large"} 6.36211e+08
# HELP am4_stats_flights_operated Company flights operated value.
# TYPE am4_stats_flights_operated gauge
am4_stats_flights_operated 82737
# HELP am4_stats_passengers_transported Passengers transported by type.
# TYPE am4_stats_passengers_transported gauge
am4_stats_passengers_transported{type="business"} 4.010239e+06
am4_stats_passengers_transported{type="economy"} 1.528383e+07
am4_stats_passengers_transported{type="first"} 2.27574e+06
```

</details>


## Grafana Dashboard

You can use the following [Grafana dashboard](https://grafana.com/grafana/dashboards/24308-airline-manager/) to visualize the Prometheus metrics collected by the bot.

![Grafana dashboard](resources/Grafana_dashboard.png?raw=true "Grafana Dashboard Screenshot")


## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


## Maintainer
GitHub: [@ashokhin](https://github.com/ashokhin)
