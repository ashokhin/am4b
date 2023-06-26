# Airline Manager 4 Bot

### Install
1. Download and install [Python3](https://www.python.org/downloads/) in your system (Don't forget to add to PATH).
2. Download a [Google Chrome](https://www.google.com/chrome/).
3. Import am4b (see examples in [Examples]() section).
4. Install Python requirements `pip install -r requirements.txt`.

### Features
- ✔️ Autostart marketing companies.
- ✔️ Autodepartures.
- ✔️ Autobuy Fuel when price is low.
- ✔️ Autobuy CO2 when price is low.
- ✔️ Autorepair.

### AM4Bot variables
| Variable | Type | Default value | Description | Usage example |
|---|---|---|---|---|
| am4_base_url | str | `"https://www.airlinemanager.com/"` | Base URL string for AM4 site | `am4bot.am4_base_url = "https://airlinemanager.com/"` |
| username | str | `""` | Username/email for AM4 | `am4bot.username = "bob@example.com"` |
| password | str | `""` | Password for AM4 | `am4bot.password = "PA5sW0rD"` |
| fuel_good_price | int | `350` | Fuel good price for buying | `am4bot.fuel_good_price = 450` |
| co2_good_price | int | `120` | CO2 good price for buying | `am4bot.co2_good_price = 130` |
| fuel_budget_percent | int | `30` | Budged in percent from total amount of money for buying fuel and CO2 | `am4bot.fuel_budget_percent = 70` |
| maintanance_budget_percent | int | `30` | Budged in percent from total amount of money for maintanance (repairs and A-Checks) | `am4bot.maintanance_budget_percent = 40` |
| marketing_budget_percent | int | `30` | Budged in percent from total amount of money for marketing companies | `am4bot.marketing_budget_percent = 40` |
| aircraft_wear_percent | int | `20` | Maximum percent of wear aircraft for repair | `am4bot.aircraft_wear_percent = 30` |
| aircraft_max_hours_to_acheck | int | `24` | Maximum hours before aircraft needs A-Check | `am4bot.aircraft_max_hours_to_acheck = 12` |

### AM4Bot methods

| Method | Arguments | Description | Usage example |
|---|---|---|---|
| `AirlineManager4Bot` | None | Initialize AM4Bot class | `am4bot = am4b.AirlineManager4Bot()` |
| `login` | None | Login to AM4 manager | `am4bot.login()` |
| `marketing_companies` | None | Check and activate marketing companies | `am4bot.marketing_companies()` |
| `depart` | None | Depart all landed aircrafts | `am4bot.depart()` |
| `check_money` | None | Check available account money | `am4bot.check_money()` |
| `check_fuel` | None | Check fuel and CO2 prices and capacity | `am4bot.check_fuel()` |
| `get_info` | None | Check available account money, fuel and CO2 prices and capacity then print they | `am4bot.get_info()` |
| `buy_fuel` | None | Buy fuel and CO2 if they available for good prices and enough volumes for budget | `am4bot.buy_fuel()` |
| `do_maintanance` | None | Do maintanance (repairs and A-Checks) for all aircrafts which it need and which towards to hub | `am4bot.do_maintanance()` |
| `run_once` | None | Run all actions in following order: `login` -> `marketing_companies` -> `do_maintanance` -> `depart` -> `buy_fuel` | `am4bot.run_once()` |
| `run_service` | `seconds_to_sleep: int` | Run `run_once` in infinity loop with pause every `seconds_to_sleep` | `am4bot.run_service(seconds_to_sleep=300)` |

### Examples
1. Create `app.py`
2. Set `username`, `password` etc. (See available variables in [AM4Bot variables]() section).
3. Run `app.py`:
``` bash
python app.py
```

`app.py` example with rewrite all available variables and run as a service:

``` python
import am4b

am4bot = am4b.AirlineManager4Bot()

am4bot.am4_base_url = "https://airlinemanager.com/"
am4bot.username = "bob@example.com"
am4bot.password = "PA5sW0rD"
am4bot.fuel_good_price = 450
am4bot.co2_good_price = 130
am4bot.fuel_budget_percent = 70
am4bot.maintanance_budget_percent = 40
am4bot.marketing_budget_percent = 40
am4bot.aircraft_wear_percent = 30
am4bot.aircraft_max_hours_to_acheck = 12

am4bot.run_service(300)
```

`app.py` example with logging, minimum rewrites and run once:
``` python
import logging
import sys

import am4b


"""
logging.basicConfig(format=u'[%(asctime)s][%(levelname)-8s][PID:%(process)d] %(funcName)s.%(lineno)d: %(message)s', 
                    level=logging.DEBUG, stream=sys.stdout)
"""

logging.basicConfig(format=u'[%(asctime)s] %(message)s', 
                    level=logging.INFO, stream=sys.stdout)

am4bot = am4b.AirlineManager4Bot()

am4bot.username = "bob@example.com"
am4bot.password = "PA5sW0rD"

am4bot.run_once()
```

`app.py` example with minimum rewrites and full flow control (except `run_once` and/or `run_service`):
``` python
import am4b


am4bot = am4b.AirlineManager4Bot()

am4bot.username = "bob@example.com"
am4bot.password = "PA5sW0rD"

am4bot.login()
am4bot.marketing_companies()
am4bot.buy_fuel()
am4bot.do_maintanance()
am4bot.depart()
```