package com.ashokhin.am4bot;

import java.util.Map;
import java.util.Optional;
import java.util.concurrent.atomic.AtomicBoolean;

import org.apache.commons.cli.CommandLine;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import com.ashokhin.am4bot.bot.Bot;
import com.ashokhin.am4bot.bot.MetricsCollector;
import com.ashokhin.am4bot.utils.CliArgumentParser;

import io.prometheus.metrics.core.metrics.Gauge;
import io.prometheus.metrics.exporter.httpserver.HTTPServer;
import io.prometheus.metrics.instrumentation.jvm.JvmMetrics;
import io.prometheus.metrics.model.registry.PrometheusRegistry;

public class PrometheusExporter {
    private static final Logger logger = LogManager.getLogger(PrometheusExporter.class);
    private static final String PROMETHEUS_MAMESPACE = "am4";
    private static final int REGISTRY_UPDATE_INTERVAL_SEC = 10;
    private static MetricsCollector metricsCollector;
    private static AtomicBoolean metricsRegistired = new AtomicBoolean(false);
    private static CommandLine cmd;
    private static String baseUrl = "https://www.airlinemanager.com/";
    private static String username;
    private static String password;
    // am4 metrics
    private static Map<String, Float> metrics;
    private static PrometheusRegistry registry = new PrometheusRegistry();
    // prometheus metrics
    private static Gauge acFleetSize;
    private static Gauge acHangarCapacity;
    private static Gauge acStatus;
    private static Gauge acRoutes;
    private static Gauge companyFuelHolding;
    private static Gauge companyFuelLimit;
    private static Gauge companyHubs;
    private static Gauge companyMoney;
    private static Gauge companyReputation;
    private static Gauge companyShareValue;
    private static Gauge marketFuelPrice;
    private static Gauge statsCargoTransported;
    private static Gauge statsFlightsOperated;
    private static Gauge statsPassengersTransported;

    private static Bot bot;

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
            throw e;
        }

        return result;
    }

    private static void registerMetrics() throws Exception {
        // Avoid of try to register multiple times
        if (metricsRegistired.get()) {

            return;
        }

        while (!bot.hasNewMetrics().get()) {
            logger.debug("Waiting for AM4 metrics...");

            Thread.sleep(10000);
        }

        logger.debug("Registering AM4 metrics...");

        acFleetSize = Gauge.builder()
                .name(String.format("%s_ac_fleet_size", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("aircraft fleet size")
                .register(registry);

        acHangarCapacity = Gauge.builder()
                .name(String.format("%s_ac_hangar_capacity", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("aircraft hangar capacity")
                .register(registry);

        acRoutes = Gauge.builder()
                .name(String.format("%s_ac_routes", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("aircraft routes created")
                .register(registry);

        acStatus = Gauge.builder()
                .name(String.format("%s_ac_status", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("aircraft status by type")
                .labelNames("type")
                .register(registry);

        companyFuelHolding = Gauge.builder()
                .name(String.format("%s_company_fuel_holding", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("company fuel holding by fuel type")
                .labelNames("type")
                .register(registry);

        companyFuelLimit = Gauge.builder()
                .name(String.format("%s_company_fuel_limit", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("company fuel holding limit by fuel type")
                .labelNames("type")
                .register(registry);

        companyHubs = Gauge.builder()
                .name(String.format("%s_company_hubs", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("company hubs")
                .register(registry);

        companyMoney = Gauge.builder()
                .name(String.format("%s_company_money", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("company money by account")
                .labelNames("account")
                .register(registry);

        companyReputation = Gauge.builder()
                .name(String.format("%s_company_reputation", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("company reputation by type")
                .labelNames("type")
                .register(registry);

        companyShareValue = Gauge.builder()
                .name(String.format("%s_company_share_value", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("company share value")
                .register(registry);

        marketFuelPrice = Gauge.builder()
                .name(String.format("%s_market_fuel_price", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("market fuel price by fuel type")
                .labelNames("type")
                .register(registry);

        statsCargoTransported = Gauge.builder()
                .name(String.format("%s_stats_cargo_transported", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("statistics cargo trasported by load type")
                .labelNames("type")
                .register(registry);

        statsFlightsOperated = Gauge.builder()
                .name(String.format("%s_stats_flights_operated", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("statistics flights operated")
                .register(registry);

        statsPassengersTransported = Gauge.builder()
                .name(String.format("%s_stats_passengers_transported", PrometheusExporter.PROMETHEUS_MAMESPACE))
                .help("statistics assengers trasported by class type")
                .labelNames("type")
                .register(registry);

        metricsRegistired = new AtomicBoolean(true);

        logger.debug("AM4 metrics registered.");
    }

    private static void updateRegistry() throws Exception {
        while (true) {
            if (bot.hasNewMetrics().getAndSet(false)) {
                logger.debug("Updating registry...");

                metrics = PrometheusExporter.metricsCollector.getMetrics();
                acFleetSize.set(metrics.get("FleetSize"));
                acHangarCapacity.set(metrics.get("HangarCapacity"));
                acRoutes.set(metrics.get("Routes"));
                acStatus.labelValues("in_flight").set(metrics.get("ACInflight"));
                acStatus.labelValues("pending_delivery").set(metrics.get("PendingDelivery"));
                acStatus.labelValues("pending_maintenance").set(metrics.get("PendingMaintenance"));
                companyFuelHolding.labelValues("co2").set(metrics.get("Co2Quota"));
                companyFuelHolding.labelValues("fuel").set(metrics.get("FuelHolding"));
                companyFuelLimit.labelValues("co2").set(metrics.get("Co2Limit"));
                companyFuelLimit.labelValues("fuel").set(metrics.get("FuelLimit"));
                companyHubs.set(metrics.get("Hubs"));
                companyMoney.labelValues("airline").set(metrics.get("AirlineAccountMoney"));
                companyReputation.labelValues("cargo").set(metrics.get("CargoReputation"));
                companyReputation.labelValues("airline").set(metrics.get("AirlineReputation"));
                companyShareValue.set(metrics.get("ShareValue"));
                marketFuelPrice.labelValues("co2").set(metrics.get("Co2Cost"));
                marketFuelPrice.labelValues("fuel").set(metrics.get("FuelCost"));
                statsCargoTransported.labelValues("heavy").set(metrics.get("HeavyLoad"));
                statsCargoTransported.labelValues("large").set(metrics.get("LargeLoad"));
                statsFlightsOperated.set(metrics.get("FlightsOperated"));
                statsPassengersTransported.labelValues("economy").set(metrics.get("YClassPax"));
                statsPassengersTransported.labelValues("business").set(metrics.get("JClassPax"));
                statsPassengersTransported.labelValues("first").set(metrics.get("FClassPax"));

                logger.debug("Registry updated.");
            }

            logger.trace(String.format("Sleeping for %d sec.", PrometheusExporter.REGISTRY_UPDATE_INTERVAL_SEC));
            Thread.sleep(10000);
        }
    }

    public static void main(String[] args) throws Exception {
        JvmMetrics.builder().register(registry);

        HTTPServer server = HTTPServer.builder()
                .registry(registry)
                .port(9400)
                .buildAndStart();

        logger.info(String.format("HTTPServer listening on port http://localhost:%d/metrics", server.getPort()));

        logger.info("Initializing AM4 exporter...");

        if (args.length > 0) {
            cmd = CliArgumentParser.parseArgs(args);
        }

        logger.debug("Processing arguments and environment variables...");
        username = getCmdArgOrDefault("username", System.getenv("AM4_USERNAME"));

        if (username == null) {
            throw new IllegalArgumentException(
                    "Argument '--username' or environment variable 'AM4_USERNAME' should be set!");

        }

        password = getCmdArgOrDefault("password", System.getenv("AM4_PASSWORD"));

        if (password == null) {
            throw new IllegalArgumentException(
                    "Argument '--password' or environment variable 'AM4_PASSWORD' should be set!");
        }

        baseUrl = getCmdArgOrDefault("url", baseUrl);

        bot = new Bot(baseUrl, username, password);

        try {
            // start parallel thread of MetricsCollector which will collect metrics every
            // 5 min.
            PrometheusExporter.metricsCollector = bot.startMetricsCollector();
            metrics = PrometheusExporter.metricsCollector.getMetrics();
        } catch (InterruptedException e) {
            logger.debug("Prometheus execution interrupted");
            bot.quit();
        }

        PrometheusExporter.registerMetrics();
        logger.info("AM4 metrics collection process started");
        PrometheusExporter.updateRegistry();
    }
}
