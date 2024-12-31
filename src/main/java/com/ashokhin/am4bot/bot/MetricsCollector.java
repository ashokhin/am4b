package com.ashokhin.am4bot.bot;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.concurrent.atomic.AtomicBoolean;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.openqa.selenium.WebElement;

import com.ashokhin.am4bot.model.APIXpath;
import com.ashokhin.am4bot.model.HubsStats;

public final class MetricsCollector implements Runnable {
    private static final int COLLECTION_INTERVAL_SEC = 60;
    private static final Logger logger = LogManager.getLogger(Bot.class);
    private static HashMap<String, Float> metricsMap = buildMetricsMap();
    private static HashMap<String, HashMap<String, Float>> complexMetricsMap = buildComplexMap();
    private static List<HubsStats> hubsStatsList = new ArrayList<HubsStats>();
    private static AtomicBoolean metricsCollected = new AtomicBoolean(false);
    private Bot bot;

    public MetricsCollector(Bot bot) {
        this.bot = bot;
    }

    private static HashMap<String, Float> buildMetricsMap() {
        metricsMap = new HashMap<>();
        // metricsMap.put("AirlineAccountMoney", 0.0f);
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
        metricsMap.put("FuelLimit", 0.0f);
        metricsMap.put("Co2Limit", 0.0f);

        return metricsMap;
    }

    private static HashMap<String, HashMap<String, Float>> buildComplexMap() {
        complexMetricsMap = new HashMap<String, HashMap<String, Float>>();
        complexMetricsMap.put("AccountMoney", null);

        return complexMetricsMap;
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

    private final void checkMoney() throws Exception {
        logger.trace("Opening 'Account' popup...");

        bot.clickButton(APIXpath.xpathButtonAccount);

        logger.trace("Getting data for each account...");

        List<WebElement> accountsList = bot.getElements(APIXpath.xpathElementListAccountLines);
        HashMap<String, Float> accountData = new HashMap<String, Float>();

        for (WebElement accountWebElement : accountsList) {
            String accountName = "N/A";
            WebElement accountNameWebElement = bot.getSubElement(accountWebElement, APIXpath.xpathElementAccountLine,
                    0);

            if (accountNameWebElement != null) {
                accountName = bot.getTextFromElement(accountNameWebElement);
            }

            float accountAmount = -1.0f;
            WebElement accountAmountWebElement = bot.getSubElement(accountWebElement, APIXpath.xpathElementAccountLine,
                    1);

            if (accountAmountWebElement != null) {
                accountAmount = bot.getIntFromElement(accountAmountWebElement).floatValue();
            }

            logger.trace(String.format("Account '%s' has $%f", accountName, accountAmount));

            accountData.put(accountName, accountAmount);

        }

        complexMetricsMap.put("AccountMoney", accountData);
        bot.clickButton(APIXpath.xpathButtonPopupClose);
    }

    private final void checkHubs() throws Exception {
        logger.trace("Open 'Hubs' popup...");

        bot.clickButton(APIXpath.xpathButtonHubs);
        List<WebElement> hubsList = bot.getElements(APIXpath.xpathElementListHubs);

        for (WebElement hubWebElement : hubsList) {
            WebElement hubNameElement = bot.getSubElement(hubWebElement, APIXpath.xpathTextHubName);

            String hubName = bot.getTextFromElement(hubNameElement);

            // Departures
            logger.debug(String.format("Get 'Departures' for the hub '%s'", hubName));
            WebElement hubStatWebElement = bot.getSubElement(hubWebElement, APIXpath.xpathTextHubDepartures);
            Float HubStat = bot.getIntFromElement(hubStatWebElement).floatValue();
            logger.trace(String.format("The 'Departures' stat for the hub '%s': %f", hubName, HubStat));
            HubsStats hubStat = new HubsStats(hubName, "Departures", HubStat);
            hubsStatsList.add(hubStat);
            // Arrivals
            logger.debug(String.format("Get 'Arrivals' for the hub '%s'", hubName));
            hubStatWebElement = bot.getSubElement(hubWebElement, APIXpath.xpathTextHubArrivals);
            HubStat = bot.getIntFromElement(hubStatWebElement).floatValue();
            logger.trace(String.format("The 'Arrivals' stat for the hub '%s': %f", hubName, HubStat));
            hubStat = new HubsStats(hubName, "Arrivals", HubStat);
            hubsStatsList.add(hubStat);
            // PAX Departed
            logger.debug(String.format("Get 'PAX Departed' for the hub '%s'", hubName));
            hubStatWebElement = bot.getSubElement(hubWebElement, APIXpath.xpathTextHubPaxDeparted);
            HubStat = bot.getIntFromElement(hubStatWebElement).floatValue();
            logger.trace(String.format("The 'PAX Departed' stat for the hub '%s': %f", hubName, HubStat));
            hubStat = new HubsStats(hubName, "PAX Departed", HubStat);
            hubsStatsList.add(hubStat);
            // PAX Arrived
            logger.debug(String.format("Get 'PAX Arrived' for the hub '%s'", hubName));
            hubStatWebElement = bot.getSubElement(hubWebElement, APIXpath.xpathTextHubPaxArrived);
            HubStat = bot.getIntFromElement(hubStatWebElement).floatValue();
            logger.trace(String.format("The 'PAX Arrived' stat for the hub '%s': %f", hubName, HubStat));
            hubStat = new HubsStats(hubName, "PAX Arrived", HubStat);
            hubsStatsList.add(hubStat);
        }

        bot.clickButton(APIXpath.xpathButtonPopupClose);
    }

    private final void checkOverview() throws Exception {
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

        return;
    }

    private final void checkFuel() throws Exception {
        logger.trace("Getting 'FuelLimit'...");

        bot.clickButton(APIXpath.xpathAllFuelElementsMap.get("common").get("xpathButtonFuelMenu"));
        bot.clickButton(APIXpath.xpathAllFuelElementsMap.get("common").get("xpathButtonFuelTab"));

        metricsMap.put(
                "FuelLimit",
                bot.getIntFromElement(APIXpath.xpathAllFuelElementsMap
                        .get("fuel").get("xpathTextMaxCapacity")).floatValue());

        logger.trace("Getting 'Co2Limit'...");

        bot.clickButton(APIXpath.xpathAllFuelElementsMap.get("common").get("xpathButtonCO2Tab"));

        metricsMap.put(
                "Co2Limit",
                bot.getIntFromElement(APIXpath.xpathAllFuelElementsMap
                        .get("co2").get("xpathTextMaxCapacity")).floatValue());

        bot.clickButton(APIXpath.xpathButtonPopupClose);

        return;
    }

    private final void collect() throws Exception {
        logger.debug("Collecting AM4 metrics...");
        bot.startBot();

        while (true) {
            logger.debug("Collecting metrics...");
            bot.refreshPage();
            this.checkMoney();
            this.checkHubs();
            this.checkOverview();
            this.checkFuel();

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
            try {
                bot.quit();
            } catch (Exception e1) {
                e1.printStackTrace();
            }
            System.exit(1);
        }
    }

    public final HashMap<String, Float> getMetrics() {
        return MetricsCollector.metricsMap;
    }

    public final HashMap<String, HashMap<String, Float>> getComplexMetrics() {
        return MetricsCollector.complexMetricsMap;
    }

    public final List<HubsStats> getHubsMetricsList() {
        return MetricsCollector.hubsStatsList;
    }

    public final AtomicBoolean isUpdated() {
        return MetricsCollector.metricsCollected;
    }
}
