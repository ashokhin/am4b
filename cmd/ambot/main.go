package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
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
	APP_NAME      string = "ambot"
	EXPORTER_NAME string = "ambot_exporter"
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

	prometheusRegistry := prometheus.NewRegistry()
	prometheusRegistry.MustRegister(versionCollector.NewCollector(EXPORTER_NAME))
	prometheusRegistry.MustRegister(collectors.NewGoCollector())

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
	// create cron job
	c.AddFunc(bot.Conf.CronSchedule, func() {
		slog.Warn("start job", "start_time", time.Now().UTC())

		if err := bot.Run(ctx); err != nil {
			slog.Warn("error in Bot.Run", "error", err)

			bot.PrometheusMetrics.Up.Set(0)
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

	if err := http.ListenAndServe(*webAddr, nil); err != nil {
		slog.Error("error in http server", "error", err)

		return
	}
}
