package bot

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/ashokhin/am4bot/internal/config"
	"github.com/ashokhin/am4bot/internal/metrics"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/chromedp/chromedp"
)

// Bot represents the automation bot with its configuration and state.
type Bot struct {
	Conf              *config.Config
	chromeOpts        []chromedp.ExecAllocatorOption
	AccountBalance    float64
	PrometheusMetrics metrics.Metrics
}

// New creates a new Bot instance with the provided configuration and Prometheus registry.
func New(conf *config.Config, registry *prometheus.Registry) Bot {
	metrics := metrics.New()
	metrics.RegisterMetrics(registry)
	metrics.StartTime.SetToCurrentTime()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.WindowSize(1920, 1080),
		// Change to 'false' for displaying Chrome window
		chromedp.Flag("headless", conf.ChromeHeadless),
		chromedp.Flag("start-maximized", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36"),
	)

	return Bot{
		Conf:              conf,
		chromeOpts:        opts,
		PrometheusMetrics: *metrics,
	}
}

// Run executes the bot's main workflow, including authentication and service tasks.
func (b *Bot) Run(ctx context.Context) error {
	timeStart := time.Now()

	slog.Debug("create context with timeout", "timeout_seconds", b.Conf.TimeoutSeconds)
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(b.Conf.TimeoutSeconds)*time.Second)
	defer cancel()

	allocatorCtx, cancel := chromedp.NewExecAllocator(timeoutCtx, b.chromeOpts...)
	defer cancel()

	taskCtx, cancel := chromedp.NewContext(
		allocatorCtx,
		chromedp.WithLogf(log.Printf),
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	slog.Debug("run bot", "start_time", timeStart.UTC())
	slog.Info("authentication")

	// perform authentication
	if err := b.auth(taskCtx); err != nil {
		slog.Warn("error in Bot.Run > Bot.auth", "error", err)

		return err
	}

	// perform money check
	if err := b.Money(taskCtx); err != nil {
		slog.Warn("error in Bot.Run > Bot.Money", "error", err)

		return err
	}

	// iterate over configured services and execute them
	for _, serviceName := range b.Conf.Services {
		switch serviceName {
		case "company_stats":
			if err := b.companyStats(taskCtx); err != nil {
				slog.Warn("error in Bot.Run > Bot.companyStats", "error", err)

				return err
			}

		case "alliance_stats":
			if err := b.allianceStats(taskCtx); err != nil {
				slog.Warn("error in Bot.Run > Bot.allianceStats", "error", err)

				return err
			}

		case "staff_morale":
			if err := b.staffMorale(taskCtx); err != nil {
				slog.Warn("error in Bot.Run > Bot.staffMorale", "error", err)

				return err
			}

		case "hubs":
			if err := b.hubs(taskCtx); err != nil {
				slog.Warn("error in Bot.Run > Bot.hubs", "error", err)

				return err
			}

		case "buy_fuel":
			if err := b.fuel(taskCtx); err != nil {
				slog.Warn("error in Bot.Run > Bot.fuel", "error", err)

				return err
			}

		case "marketing":
			if err := b.marketingCompanies(taskCtx); err != nil {
				slog.Warn("error in Bot.Run > Bot.marketingCompanies", "error", err)

				return err
			}

		case "ac_maintenance":
			if err := b.maintenance(taskCtx); err != nil {
				slog.Warn("error in Bot.Run > Bot.maintenance", "error", err)

				return err
			}
		case "depart":
			if err := b.depart(taskCtx); err != nil {
				slog.Warn("error in Bot.Run > Bot.depart", "error", err)

				return err
			}

		default:
			slog.Warn("unknown service", "service", serviceName,
				"available_services",
				[]string{"company_stats", "staff_morale", "alliance_stats", "hubs", "buy_fuel", "depart", "marketing", "ac_maintenance"})
		}
	}

	// calculate total duration for Prometheus metric
	duration := time.Since(timeStart)

	slog.Info("run complete", "elapsed_time", fmt.Sprint(duration))

	b.PrometheusMetrics.Duration.Set(duration.Seconds())

	return nil
}
