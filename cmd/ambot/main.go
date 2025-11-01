package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/ashokhin/am4bot/internal/bot"
	"github.com/ashokhin/am4bot/internal/config"

	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	versionCollector "github.com/prometheus/client_golang/prometheus/collectors/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promslog"
	"github.com/prometheus/common/promslog/flag"
	"github.com/prometheus/common/version"
	"github.com/robfig/cron/v3"
)

const (
	APP_NAME           string = "ambot"
	EXPORTER_NAME      string = "ambot_exporter"
	EXPORTER_NAMESPACE string = "am4"
)

var (
	configFile   = kingpin.Flag("app.config", "YAML file with configuration.").Short('c').Default("config.yaml").String()
	webAddr      = kingpin.Flag("web.listen-address", "Addresses on which to expose metrics and web interface.").Default(":9150").String()
	webTelemetry = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
)

func main() {
	var err error
	var conf *config.Config

	promslogConfig := &promslog.Config{}
	flag.AddFlags(kingpin.CommandLine, promslogConfig)
	kingpin.Version(version.Print(APP_NAME))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	logger := promslog.New(promslogConfig)
	slog.SetDefault(logger)

	slog.Info(fmt.Sprintf("starting application %s", APP_NAME), "version", version.Info())
	slog.Info("build context", "build_context", version.BuildContext())

	// load configuration
	confPath, _ := filepath.Abs(*configFile)

	if conf, err = config.New(confPath); err != nil {
		slog.Error("config loading error", "error", err)

		return
	}

	// The CLI's "log.level" and config's "log_level" by default are both "info"
	// if they are not -- check further
	if promslogConfig.Level.String() != conf.LogLevel {
		// If CLI's "log.level" is not default (info) then prioritize CLI's value
		if promslogConfig.Level.String() != "info" {
			slog.Info("set log level from CLI", "log.level", promslogConfig.Level.String())

		} else { // else - set "log_level" from config
			slog.Info("set log level from config", "log_level", conf.LogLevel)

			promslogConfig.Level.Set(conf.LogLevel)
		}
	}

	// The CLI's "web.listen-address" and config's "prometheus_address" by default are both ":9150"
	// if they are not -- check further
	if *webAddr != conf.PrometheusAddress {
		// If CLI's "web.listen-address" is not default (:9150) then prioritize CLI's value
		if *webAddr != ":9150" {
			slog.Info("set Prometheus address from CLI", "address", *webAddr)

		} else { // else - set "prometheus_address" from config
			slog.Info("set Prometheus address from config", "address", conf.PrometheusAddress)

			*webAddr = conf.PrometheusAddress
		}
	}

	// create Prometheus registry
	prometheusRegistry := prometheus.NewRegistry()
	prometheusRegistry.MustRegister(versionCollector.NewCollector(EXPORTER_NAMESPACE))
	prometheusRegistry.MustRegister(collectors.NewGoCollector())

	// create Bot object with loaded configuration
	bot := bot.New(conf, prometheusRegistry)

	handler := promhttp.HandlerFor(
		prometheusRegistry,
		promhttp.HandlerOpts{
			Registry: prometheusRegistry,
		})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start it once in the blocking mode (not inside a goroutine)
	// for collecting initial Prometheus metrics
	if err := bot.Run(ctx); err != nil {
		slog.Warn("error in Bot.Run", "error", err)

		bot.PrometheusMetrics.Up.Set(0)
	} else {
		bot.PrometheusMetrics.Up.Set(1)
	}

	// now start it inside "cronjob" (goroutine with schedule)

	// create cron object
	c := cron.New()
	// create cron job with schedule from configuration
	c.AddFunc(bot.Conf.CronSchedule, func() {
		slog.Warn("start job", "start_time", time.Now().UTC())

		if err := bot.Run(ctx); err != nil {
			slog.Error("error in Bot.Run", "error", err)

			os.Exit(1)
		} else {
			bot.PrometheusMetrics.Up.Set(1)
		}

		slog.Warn("job done", "end_time", time.Now().UTC(), "next_run", c.Entry(1).Next.UTC())
	})

	// start cron object, schedule jobs
	c.Start()

	slog.Info("job scheduled", "next_run", c.Entry(1).Next.UTC())

	// register handler for the webTelemetry page
	http.Handle(*webTelemetry, handler)

	slog.Info(fmt.Sprintf("starting Prometheus exporter %s", EXPORTER_NAME), "address", *webAddr, "location", *webTelemetry)

	// start HTTP server for Prometheus scraping
	if err := http.ListenAndServe(*webAddr, nil); err != nil {
		slog.Error("error in http server", "error", err)

		return
	}
}
