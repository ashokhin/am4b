import argparse
import logging
import sys


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
                        default=500)
    parser.add_argument('--co2-good-price', type=int, help="Good price for 1000 Quotas of CO2",
                        default=120)
    parser.add_argument('--fuel-budget-percent', type=int, help="Percent of account money available for buying fuel/CO2",
                        default=70)
    parser.add_argument('--maintenance-budget-percent', type=int, help="Percent of account money available for maintenance",
                        default=50)
    parser.add_argument('--marketing-budget-percent', type=int, help="Percent of account money available for marketing",
                        default=70)
    parser.add_argument('--aircraft-wear-percent', type=int, help="Percent of aircraft wear for repair",
                        default=80)
    parser.add_argument('--aircraft-max-hours-to-acheck', type=int, help="Percent of aircraft A-Check for repair",
                        default=24)
    parser.add_argument('--run-mode', type=str, choices=['once', 'service'], help="Run mode",
                        default='once')
    parser.add_argument('--service-sleep-sec', type=int, help="Seconds to sleep between iterations when program runs as a service",
                        default=300)
    parser.add_argument('--scanner-file', type=str, help="Path to file for save scanner results",
                        default="am4scanner.csv")

    return parser.parse_args()


def run_bot(args: argparse.Namespace) -> None:
    import am4.bot as bot

    am4bot = bot.AirlineManager4Bot()

    am4bot.am4_base_url = args.base_url
    am4bot.username = args.username
    am4bot.password = args.password
    am4bot.fuel_good_price = args.fuel_good_price
    am4bot.co2_good_price = args.co2_good_price
    am4bot.fuel_budget_percent = args.fuel_budget_percent
    am4bot.maintenance_budget_percent = args.maintenance_budget_percent
    am4bot.marketing_budget_percent = args.marketing_budget_percent
    am4bot.aircraft_wear_percent = args.aircraft_wear_percent
    am4bot.aircraft_max_hours_to_a_check = args.aircraft_max_hours_to_acheck
    
    try:
        if args.run_mode == 'once':
            am4bot.run_once()
        else:
            am4bot.run_service(seconds_to_sleep=args.service_sleep_sec)
    except KeyboardInterrupt:
            logging.info("Program interrupted by user")
            sys.exit(0)
    except Exception as ex:
        logging.exception("Exception:\n{}".format(ex))
        sys.exit(1)


def run_scanner(args: argparse.Namespace) -> None:
    import am4.scanner as scanner

    am4scanner = scanner.AM4Scanner()

    am4scanner.am4_base_url = args.base_url
    am4scanner.username = args.username
    am4scanner.password = args.password
    am4scanner.file_path = args.scanner_file

    am4scanner.scan()


def main():
    """
    logging.basicConfig(format=u'[%(asctime)s][%(levelname)-8s][PID:%(process)d] %(funcName)s.%(lineno)d: %(message)s', 
                        level=logging.DEBUG, stream=sys.stdout)
    """
    
    logging.basicConfig(format=u'[%(asctime)s][%(levelname)-8s] %(message)s', 
                        level=logging.INFO, stream=sys.stdout)

    args = parse_arguments()

    if args.run_mode == 'once' or args.run_mode == 'service':
        run_bot(args)
    
    if args.run_mode == 'scanner':
        run_scanner(args)


if __name__ == "__main__":
    main()
