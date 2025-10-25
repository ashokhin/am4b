# Airline Manager 4 Bot

### Features
- ✔️ Automatic start marketing companies.
- ✔️ Automatic departures.
- ✔️ Automatic buy Fuel when price is low or fuel level is critical.
- ✔️ Automatic buy CO2 when price is low or quota level is critical.
- ✔️ Automatic staff morale improvement.
- ✔️ Automatic hub management.
- ✔️ Automatic company statistics collection.
- ✔️ Automatic repair.
- ✔️ Automatic A-Check.
- ✔️ Automatic modification.

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
   docker run --rm --name ambot --volume /opt/ambot/conf/config.yaml:/app/conf/config.yaml ashokhin/am4bot:latest
   ```
   
   For collecting Prometheus metrics, you can expose port 9150:
   ```bash
   docker run --rm --name ambot --volume /opt/ambot/conf/config.yaml:/app/conf/config.yaml -p 9150:9150 ashokhin/am4bot:latest
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
   ExecStart=/usr/bin/docker run --rm --name %n --volume /opt/ambot/conf/config.yaml:/app/conf/config.yaml --publish 9150:9150 ashokhin/am4bot:latest

   [Install]
   WantedBy=multi-user.target
   ```
   
   Then enable and start the service:
   ```bash
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
| `good_price.co2` | int | `120` | Good price for CO2 (per Quotas). |
| `aircraft_wear_percent` | int | `80` | Aircraft wear percentage to trigger maintenance. |
| `aircraft_max_hours_to_check` | int | `24` | Max hours to next A-Check to trigger it. |
| `aircraft_modify_limit` | int | `3` | Max aircraft for modifications checks. |
| `fuel_critical_percent` | float | `20` | Fuel level percentage to trigger refuel. |
| `service_cron_string` | string | `"*/5 * * * *"` | Cron schedule for services. Default: Every 5 minutes. |
| `services` | list of strings | `["company_stats","staff_morale",`<br />`"hubs","buy_fuel","depart",`<br />`"marketing_companies","ac_maintenance"]` | List of services to run. Possible values: `company_stats`, `staff_morale`,<br />`hubs`, `buy_fuel`, `depart`,<br />`marketing_companies`, `ac_maintenance`. |
| `timeout_seconds` | int | `120` | Timeout for full round in seconds. |
| `chrome_headless` | bool | `true` | Run browser in headless mode. |

Example of `config.yaml` with the default options:
```yaml
url: "https://www.airlinemanager.com/"
username: ""
password: ""
log_level: "info"
budget_percent:
  fuel: 70
  maintenance: 30
  marketing: 70
good_price:
  fuel: 500
  co2: 120
aircraft_wear_percent: 80
aircraft_max_hours_to_check: 24
aircraft_modify_limit: 3
fuel_critical_percent: 20
service_cron_string: "*/5 * * * *"
services:
  - "company_stats"
  - "staff_morale"
  - "hubs"
  - "buy_fuel"
  - "marketing_companies"
  - "ac_maintenance"
  - "depart"
timeout_seconds: 120
chrome_headless: true
```

Minimal configuration example:
```yaml
username: "username@email.example"
password: "your_password_here"
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
