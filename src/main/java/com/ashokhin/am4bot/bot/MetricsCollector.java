package com.ashokhin.am4bot.bot;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.atomic.AtomicBoolean;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import com.ashokhin.am4bot.model.APIXpath;

public final class MetricsCollector implements Runnable {
    private static final int COLLECTION_INTERVAL_SEC = 300;
    private static final Logger logger = LogManager.getLogger(Bot.class);
    private static Map<String, Float> metricsMap = buildMetricsMap();
    private static AtomicBoolean metricsCollected = new AtomicBoolean(false);
    private Bot bot;

    public MetricsCollector(Bot bot) {
        this.bot = bot;
    }

    private static Map<String, Float> buildMetricsMap() {
        metricsMap = new HashMap<>();
        metricsMap.put("AirlineAccountMoney", 0.0f);
        metricsMap.put("AirlineReputation", 0.0f);
        metricsMap.put("CargoReputation", 0.0f);
        metricsMap.put("FuelCost", 0.0f);
        metricsMap.put("Co2Cost", 0.0f);
        metricsMap.put("FleetSize", 0.0f);
        metricsMap.put("PendingDelivery", 0.0f);
        metricsMap.put("Routes", 0.0f);
        metricsMap.put("Hubs", 0.0f);
        metricsMap.put("PendingMaintenance", 0.0f);
        metricsMap.put("FuelHolding", 0.0f);
        metricsMap.put("HangarCapacity", 0.0f);
        metricsMap.put("Co2Quota", 0.0f);
        metricsMap.put("ACInflight", 0.0f);
        metricsMap.put("ShareValue", 0.0f);
        metricsMap.put("FlightsOperated", 0.0f);
        metricsMap.put("YClassPax", 0.0f);
        metricsMap.put("JClassPax", 0.0f);
        metricsMap.put("FClassPax", 0.0f);
        metricsMap.put("LargeLoad", 0.0f);
        metricsMap.put("HeavyLoad", 0.0f);

        return metricsMap;
    }

    private final void sleep() throws InterruptedException {
        try {
            logger.trace(String.format("Sleeping for %d sec.", MetricsCollector.COLLECTION_INTERVAL_SEC));
            Thread.sleep(MetricsCollector.COLLECTION_INTERVAL_SEC * 1000);
        } catch (InterruptedException e) {
            logger.debug("Metrics collection interrupted. Exit.");
            Thread.currentThread().interrupt();
            throw e;
        }
    }

    private final void collect() throws Exception {
        logger.debug("Collecting AM4 metrics...");
        bot.startBot();

        while (true) {
            logger.debug("Collecting metrics...");
            logger.trace("Getting 'AirlineAccountMoney'...");
            metricsMap.put(
                    "AirlineAccountMoney",
                    bot.getMoney().floatValue());

            logger.trace("Opening 'Overview' popup...");
            bot.clickButton(APIXpath.xpathButtonOverviewMenu);

            logger.trace("Getting 'AirlineReputation'...");
            metricsMap.put(
                    "AirlineReputation",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewAirlineReputation)
                            .floatValue());

            logger.trace("Getting 'CargoReputation'...");
            metricsMap.put(
                    "CargoReputation",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewCargoReputation).floatValue());

            logger.trace("Getting 'FuelCost'...");
            metricsMap.put(
                    "FuelCost",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewFuelCost).floatValue());

            logger.trace("Getting 'Co2Cost'...");
            metricsMap.put(
                    "Co2Cost",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewCo2Cost).floatValue());

            logger.trace("Getting 'FleetSize'...");
            metricsMap.put(
                    "FleetSize",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewFleetSize).floatValue());

            logger.trace("Getting 'PendingDelivery'...");
            metricsMap.put(
                    "PendingDelivery",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewPendingDelivery).floatValue());

            logger.trace("Getting 'Routes'...");
            metricsMap.put(
                    "Routes",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewRoutes).floatValue());

            logger.trace("Getting 'Hubs'...");
            metricsMap.put(
                    "Hubs",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewHubs).floatValue());

            logger.trace("Getting 'PendingMaintenance'...");
            metricsMap.put(
                    "PendingMaintenance",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewPendingMaintenance)
                            .floatValue());

            logger.trace("Getting 'FuelHolding'...");
            metricsMap.put(
                    "FuelHolding",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewFuelHolding).floatValue());

            logger.trace("Getting 'HangarCapacity'...");
            metricsMap.put(
                    "HangarCapacity",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewHangarCapacity).floatValue());

            logger.trace("Getting 'Co2Quota'...");
            metricsMap.put(
                    "Co2Quota",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewCo2Quotas).floatValue());

            logger.trace("Getting 'ACInflight'...");
            metricsMap.put(
                    "ACInflight",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewACInflight).floatValue());

            logger.trace("Getting 'ShareValue'...");
            metricsMap.put(
                    "ShareValue",
                    bot.getFloatFromElement(APIXpath.xpathTextOverviewShareValue));

            logger.trace("Getting 'FlightsOperated'...");
            metricsMap.put(
                    "FlightsOperated",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewFlightsOperated).floatValue());

            logger.trace("Getting 'YClassPax'...");
            metricsMap.put(
                    "YClassPax",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewYClassPax).floatValue());

            logger.trace("Getting 'JClassPax'...");
            metricsMap.put(
                    "JClassPax",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewJClassPax).floatValue());

            logger.trace("Getting 'FClassPax'...");
            metricsMap.put(
                    "FClassPax",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewFClassPax).floatValue());

            logger.trace("Getting 'LargeLoad'...");
            metricsMap.put(
                    "LargeLoad",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewLargeLoad).floatValue());

            logger.trace("Getting 'HeavyLoad'...");
            metricsMap.put(
                    "HeavyLoad",
                    bot.getIntFromElement(APIXpath.xpathTextOverviewHeavyLoad).floatValue());

            bot.clickButton(APIXpath.xpathButtonPopupClose);

            MetricsCollector.metricsCollected = new AtomicBoolean(true);

            logger.debug("Metrics collected.");
            this.sleep();
        }
    }

    @Override
    public final void run() {
        try {
            this.collect();
        } catch (InterruptedException e) {
            e.printStackTrace();
            Thread.currentThread().interrupt();
        } catch (Exception e) {
            e.printStackTrace();
            bot.quit();
            System.exit(1);
        }
    }

    public final Map<String, Float> getMetrics() {
        return MetricsCollector.metricsMap;
    }

    public final AtomicBoolean isUpdated() {
        return MetricsCollector.metricsCollected;
    }
}
