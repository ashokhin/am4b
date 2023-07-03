import argparse
import logging
import sys

import am4.bot as bot

def parse_arguments() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        prog="AM4Bot",
        description="Airline Manager 4 Bot",
    )
    parser.add_argument('--base-url', type=str,
                        default="https://www.airlinemanager.com/")
    parser.add_argument('--username', type=str, help="AM4 username/email")
    parser.add_argument('--password', type=str, help="AM4 password")
    parser.add_argument('--fuel-good-price', type=int, help="Good price for 1000 Lbs of fuel",
                        default=450)
    parser.add_argument('--co2-good-price', type=int, help="Good price for 1000 Quotas of CO2",
                        default=120)
    parser.add_argument('--fuel-budget-percent', type=int, help="Percent of account money available for buying fuel/CO2",
                        default=70)
    parser.add_argument('--maintanance-budget-percent', type=int,
                        default=50)
    parser.add_argument('--marketing-budget-percent', type=int,
                        default=70)
    parser.add_argument('--aircraft-wear-percent', type=int,
                        default=30)
    parser.add_argument('--aircraft-max-hours-to-acheck', type=int,
                        default=12)
    parser.add_argument('--run-mode', type=str, choices=['once', 'service'], help="Run mode",
                        default='once')
    parser.add_argument('--service-sleep-sec', type=int, help="Seconds to sleep between iterations when program runs as a service",
                        default=300)

    return parser.parse_args()

def main():
    """
    logging.basicConfig(format=u'[%(asctime)s][%(levelname)-8s][PID:%(process)d] %(funcName)s.%(lineno)d: %(message)s', 
                        level=logging.DEBUG, stream=sys.stdout)
    """
    
    logging.basicConfig(format=u'[%(asctime)s] %(message)s', 
                        level=logging.INFO, stream=sys.stdout)

    args = parse_arguments()

    am4bot = bot.AirlineManager4Bot()

    am4bot.am4_base_url = args.base_url
    am4bot.username = args.username
    am4bot.password = args.password
    am4bot.fuel_good_price = args.fuel_good_price
    am4bot.co2_good_price = args.co2_good_price
    am4bot.fuel_budget_percent = args.fuel_budget_percent
    am4bot.maintanance_budget_percent = args.maintanance_budget_percent
    am4bot.marketing_budget_percent = args.marketing_budget_percent
    am4bot.aircraft_wear_percent = args.aircraft_wear_percent
    am4bot.aircraft_max_hours_to_acheck = args.aircraft_max_hours_to_acheck
    
    try:
        if args.run_mode == 'once':
            am4bot.run_once()
        else:
            am4bot.run_service(seconds_to_sleep=args.service_sleep_sec)
    except KeyboardInterrupt:
            logging.info("Program interupted by user")
            sys.exit(0)
    except Exception as ex:
         logging.exception("Exception:\n{}".format(ex))
         sys.exit(1)


if __name__ == "__main__":
    main()
