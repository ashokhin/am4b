package com.ashokhin.am4bot.utils;

import org.apache.commons.cli.CommandLine;
import org.apache.commons.cli.CommandLineParser;
import org.apache.commons.cli.DefaultParser;
import org.apache.commons.cli.HelpFormatter;
import org.apache.commons.cli.Options;
import org.apache.commons.cli.ParseException;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import com.ashokhin.am4bot.bot.Bot;

public final class CliArgumentParser {
    private static final Logger logger = LogManager.getLogger(Bot.class);
    private static CommandLine cmd;
    private static Options options = new Options();
    private static CommandLineParser parser = new DefaultParser();

    public static CommandLine parseArgs(String[] args) {
        options.addOption(null, "url", true, "URL for the connection (Default: 'https://www.airlinemanager.com/')");
        options.addOption("u", "username", true, "Username for authentication");
        options.addOption("p", "password", true, "Password for authentication");
        options.addOption(null, "fuel-good-price", true, "Good price for buying 1000 Lbs of the fuel (Default: 500)");
        options.addOption(null, "co2-good-price", true, "Good price for buying 1000 Quotas of the CO2 (Default: 120)");
        options.addOption(null, "crit-fuel-level", true,
                "Critical percent of fuel/CO2 level for buying beyond this point w/o waiting for a good price (Default: 20)");
        options.addOption(null, "fuel-budget-percent", true,
                "Percent of account money available for buying fuel/CO2 (Default: 50)");
        options.addOption(null, "maint-budget-percent", true,
                "Percent of account money available for maintenance aircraft (Default: 50)");
        options.addOption(null, "market-budget-percent", true,
                "Percent of account money available for marketing companies (Default: 50)");
        options.addOption(null, "air-wear-percent", true,
                "Critical percent of aircraft wear for repearing beyond this point (Default: 80)");
        options.addOption(null, "air-acheck-hours", true, "Number of hours for aircraft A-Check (Default: 24)");
        options.addOption("m", "mode", true, "Run mode. Supported values: 'once','daemon' (Default: 'once')");
        options.addOption("v", "verbose", false, "Enable verbose mode");
        options.addOption("h", "help", false, "Print this help");

        try {
            logger.debug("Parsing arguments...");
            cmd = parser.parse(options, args);
            boolean help = cmd.hasOption("help");
            if (help) {
                String header = "\nAM4Bot application for automate management of Airline Manager 4.\n\n";
                String footer = "\n© Andrei Shokhin (Github: ashokhin)";
                HelpFormatter formatter = new HelpFormatter();
                formatter.printHelp("com.ashokhin.am4bot.App", header, options, footer, true);
                System.exit(0);
            }
        } catch (ParseException e) {
            logger.fatal(String.format("Error: %s", e));
            System.exit(1);
        }

        return cmd;
    }
}
