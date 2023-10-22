package com.ashokhin.am4bot;

import java.time.Duration;
import java.time.Instant;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import com.ashokhin.am4bot.bot.Bot;

public final class App {
    private static final Logger logger = LogManager.getLogger(App.class);
    private static final String BASE_URL_STRING = "https://www.airlinemanager.com/";
    private static String loginString;
    private static String passwordString;

    public static final void main(String[] args) {
        logger.info("Start App");
        App.loginString = System.getenv("AM4_USERNAME");

        if (App.loginString == null) {
            logger.error("Environment variable 'AM4_USERNAME' should be set!");
            return;
        }

        App.passwordString = System.getenv("AM4_PASSWORD");

        if (App.passwordString == null) {
            logger.error("Environment variable 'AM4_PASSWORD' should be set!");
        }

        Instant start = Instant.now();

        Bot bot = new Bot(BASE_URL_STRING, App.loginString, App.passwordString);
        bot.setSettings(450, 120, 20, 50,
                50, 50, 80,
                24);

        try {
            bot.startOnce();
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
