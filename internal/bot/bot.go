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

type Bot struct {
	Conf              *config.Config
	chromeOpts        []chromedp.ExecAllocatorOption
	AccountBalance    float64
	PrometheusMetrics metrics.Metrics
}

func New(conf *config.Config, registry *prometheus.Registry) Bot {
	metrics := metrics.New()
	metrics.RegisterMetrics(registry)
	metrics.StartTime.SetToCurrentTime()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.WindowSize(1920, 1080),
		// Change to 'false' for displaying Chrome window
		chromedp.Flag("headless", true),
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

func (b *Bot) Run(ctx context.Context) error {
	timeStart := time.Now()

	slog.Debug("create context with the 2 minutes timeout")
	ctx, cancelToCtx := context.WithTimeout(ctx, 2*time.Minute)
	defer cancelToCtx()

	ctx, cancelChrExecCtx := chromedp.NewExecAllocator(ctx, b.chromeOpts...)
	defer cancelChrExecCtx()

	ctx, cancelChrCtx := chromedp.NewContext(
		ctx,
		chromedp.WithLogf(log.Printf),
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancelChrCtx()

	slog.Debug("run bot", "start_time", timeStart.UTC())
	slog.Info("authentication")

	if err := b.auth(ctx); err != nil {
		slog.Warn("error in Bot.Run > Bot.auth", "error", err)

		return err
	}

	if err := b.Money(ctx); err != nil {
		slog.Warn("error in Bot.Run > Bot.Money", "error", err)

		return err
	}

	for _, serviceName := range b.Conf.Services {
		switch serviceName {
		case "company_stats":
			if err := b.companyStats(ctx); err != nil {
				slog.Warn("error in Bot.Run > Bot.companyStats", "error", err)

				return err
			}

		case "staff_morale":
			if err := b.staffMorale(ctx); err != nil {
				slog.Warn("error in Bot.Run > Bot.staffMorale", "error", err)

				return err
			}

		case "hubs":
			if err := b.hubs(ctx); err != nil {
				slog.Warn("error in Bot.Run > Bot.hubs", "error", err)

				return err
			}

		case "buy_fuel":
			if err := b.fuel(ctx); err != nil {
				slog.Warn("error in Bot.Run > Bot.fuel", "error", err)

				return err
			}

		case "marketing_companies":
			if err := b.marketingCompanies(ctx); err != nil {
				slog.Warn("error in Bot.Run > Bot.marketingCompanies", "error", err)

				return err
			}

		case "ac_maintenance":
			if err := b.maintenance(ctx); err != nil {
				slog.Warn("error in Bot.Run > Bot.maintenance", "error", err)

				return err
			}
		case "depart":
			if err := b.depart(ctx); err != nil {
				slog.Warn("error in Bot.Run > Bot.depart", "error", err)

				return err
			}

		default:
			slog.Warn("unknown service", "service", serviceName,
				"available_services",
				[]string{"company_stats, staff_morale, hubs, buy_fuel, marketing_companies, ac_maintenance", "depart"})
		}
	}

	duration := time.Since(timeStart)

	slog.Info("run complete", "elapsed_time", fmt.Sprint(duration))

	b.PrometheusMetrics.Duration.Set(duration.Seconds())

	return nil
}
