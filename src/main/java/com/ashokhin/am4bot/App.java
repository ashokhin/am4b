package com.ashokhin.am4bot;

import java.time.Duration;
import java.time.Instant;
import java.util.Optional;

import org.apache.commons.cli.CommandLine;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import com.ashokhin.am4bot.bot.Bot;
import com.ashokhin.am4bot.model.BotMode;
import com.ashokhin.am4bot.utils.CliArgumentParser;

public final class App {
    private static final Logger logger = LogManager.getLogger(App.class);
    private static CommandLine cmd;
    private static String baseUrl = "https://www.airlinemanager.com/";
    private static String username;
    private static String password;
    private static Integer fuelGoodPrice = 450;
    private static Integer co2GoodPrice = 120;
    private static Integer criticalFuelLevelPercent = 20;
    private static Integer fuelBudgetPercent = 50;
    private static Integer maintenanceBudgetPercent = 50;
    private static Integer marketingBudgetPercent = 50;
    private static Integer aircraftWearPercent = 80;
    private static Integer aircraftMaximumHoursBeforeACheck = 24;
    private static long daemonSecondsWaitInterval = 66;
    private static String runMode = "once";

    public static String getCmdArgOrDefault(String cmdValue, String defaultValue) {
        try {
            return Optional.ofNullable(cmd.getOptionValue(cmdValue)).orElse(defaultValue);
        } catch (NullPointerException e) {
            return defaultValue;
        }
    }

    public static Integer getCmdArgOrDefault(String cmdValue, Integer defaultValue) {
        String cmdArgString;
        Integer result = defaultValue;
        try {
            cmdArgString = Optional.ofNullable(cmd.getOptionValue(cmdValue)).orElse(defaultValue.toString());
        } catch (NullPointerException e) {
            cmdArgString = defaultValue.toString();
        }

        try {
            result = Integer.parseInt(cmdArgString);
        } catch (NumberFormatException e) {
            logger.fatal(String.format("Error: %s", e));
            System.exit(1);
        }

        return result;
    }

    public static final void main(String[] args) {
        if (args.length > 0) {
            cmd = CliArgumentParser.parseArgs(args);
        }

        logger.info("Start App");

        username = getCmdArgOrDefault("username", System.getenv("AM4_USERNAME"));

        if (username == null) {
            logger.error("Argument '--username' or environment variable 'AM4_USERNAME' should be set!");

            System.exit(1);
        }

        password = getCmdArgOrDefault("password", System.getenv("AM4_PASSWORD"));

        if (password == null) {
            logger.error("Argument '--password' or environment variable 'AM4_PASSWORD' should be set!");

            System.exit(1);
        }

        baseUrl = getCmdArgOrDefault("url", baseUrl);
        fuelGoodPrice = getCmdArgOrDefault("fuel-good-price", fuelGoodPrice);
        co2GoodPrice = getCmdArgOrDefault("co2-good-price", co2GoodPrice);
        criticalFuelLevelPercent = getCmdArgOrDefault("crit-fuel-level", criticalFuelLevelPercent);
        fuelBudgetPercent = getCmdArgOrDefault("fuel-budget-percent", fuelBudgetPercent);
        maintenanceBudgetPercent = getCmdArgOrDefault("maint-budget-percent", maintenanceBudgetPercent);
        marketingBudgetPercent = getCmdArgOrDefault("market-budget-percent", marketingBudgetPercent);
        aircraftWearPercent = getCmdArgOrDefault("air-wear-percent", aircraftWearPercent);
        aircraftMaximumHoursBeforeACheck = getCmdArgOrDefault("air-acheck-hours", aircraftMaximumHoursBeforeACheck);
        runMode = getCmdArgOrDefault("mode", runMode);

        Instant start = Instant.now();

        Bot bot = new Bot(baseUrl, username, password);
        bot.setSettings(fuelGoodPrice, co2GoodPrice, criticalFuelLevelPercent, fuelBudgetPercent,
                maintenanceBudgetPercent, marketingBudgetPercent, aircraftWearPercent,
                aircraftMaximumHoursBeforeACheck);

        try {
            if (runMode.equals("once")) {
                bot.startOnce();
            } else if (runMode.equals("daemon")) {
                bot.setDaemon(BotMode.ALL, daemonSecondsWaitInterval);
                bot.run();
            } else if (runMode.equals("test")) {
                bot.setDaemon(BotMode.REPAIR_LOUNGE, daemonSecondsWaitInterval);
                bot.run();
            } else {
                logger.fatal(String.format("Run mode '%s' is unknown! Supported modes: 'once' and 'daemon'", runMode));
                System.exit(1);
            }

        } finally {
            bot.quit();
        }

        /*
         * 
         * Runnable botAllOperations = new Bot(_baseURL, _loginString, _passwordString,
         * BotMode.ALL, 500, 120, 20, 70,
         * 50, 70, 90,
         * 24);
         * 
         * Thread botAllOperationsThread = new Thread(botAllOperations);
         * botAllOperationsThread.setDaemon(true);
         * 
         * Runnable botBuyFuel = new Bot(_baseURL, _loginString, _passwordString,
         * BotMode.BUY_FUEL, 500, 120, 20, 70,
         * 50, 70, 90,
         * 24);
         * 
         * Thread botBuyFuelThread = new Thread(botBuyFuel);
         * botBuyFuelThread.setDaemon(true);
         * 
         * Runnable botMaintenance = new Bot(_baseURL, _loginString, _passwordString,
         * BotMode.MAINTENANCE, 500, 120, 20,
         * 70,
         * 50, 70, 90,
         * 24);
         * 
         * Thread botMaintenanceThread = new Thread(botMaintenance);
         * botMaintenanceThread.setDaemon(true);
         * 
         * Runnable botMarketing = new Bot(_baseURL, _loginString, _passwordString,
         * BotMode.MARKETING, 500, 120, 20,
         * 70,
         * 50, 70, 90,
         * 24);
         * 
         * Thread botMarketingThread = new Thread(botMarketing);
         * botMarketingThread.setDaemon(true);
         * 
         * Runnable botDispatcher = new Bot(_baseURL, _loginString, _passwordString,
         * BotMode.DEPART, 500, 120, 20,
         * 70,
         * 50, 70, 90,
         * 24);
         * 
         * Thread botDispatcherThread = new Thread(botDispatcher);
         * botDispatcherThread.setDaemon(true);
         * 
         * // botAllOperationsThread.start();
         * 
         * try {
         * botAllOperationsThread.join();
         * } catch (Exception e) {
         * e.printStackTrace();
         * }
         * 
         * botBuyFuelThread.start();
         * botMaintenanceThread.start();
         * 
         * try {
         * botBuyFuelThread.join();
         * } catch (Exception e) {
         * e.printStackTrace();
         * }
         * 
         * botMarketingThread.start();
         * 
         * try {
         * botMaintenanceThread.join();
         * } catch (Exception e) {
         * e.printStackTrace();
         * }
         * 
         * try {
         * botMarketingThread.join();
         * } catch (Exception e) {
         * e.printStackTrace();
         * }
         * 
         * botDispatcherThread.start();
         * 
         * try {
         * botDispatcherThread.join();
         * } catch (Exception e) {
         * e.printStackTrace();
         * }
         * 
         * Thread botBuyFuelAgainThread = new Thread(botBuyFuel);
         * botBuyFuelAgainThread.setDaemon(true);
         * 
         * botBuyFuelAgainThread.start();
         * 
         * try {
         * botBuyFuelAgainThread.join();
         * } catch (Exception e) {
         * e.printStackTrace();
         * }
         */

        Instant finish = Instant.now();
        Duration timeElapsed = Duration.between(start, finish);
        logger.info(String.format("Time elapsed: %s", timeElapsed));
        logger.info("End App");

    }
}
