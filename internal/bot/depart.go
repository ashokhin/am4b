package bot

import (
	"context"
	"log/slog"
	"math"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/ashokhin/am4bot/internal/utils"
	"github.com/chromedp/chromedp"
)

func (b *Bot) depart(ctx context.Context) error {
	slog.Info("depart all available aircraft")

	aircraftReadyForDepart := b.getReadyForDepart(ctx)
	maxRetries := math.Round(float64(aircraftReadyForDepart)/20) + 1

	for (aircraftReadyForDepart > 0) && (maxRetries > 0) {
		var availableAfterDepart int

		slog.Debug("depart available aircraft", "ready to depart", aircraftReadyForDepart, "depart retries", maxRetries)

		utils.DoClickElement(ctx, model.BUTTON_FI_DEPART_ALL)
		availableAfterDepart = b.getReadyForDepart(ctx)

		slog.Info("aircraft departed", "count", (aircraftReadyForDepart - availableAfterDepart))

		aircraftReadyForDepart = availableAfterDepart

		maxRetries--

		// try to buy fuel after each depart iteration
		b.fuel(ctx)
	}

	return nil
}

func (b *Bot) getReadyForDepart(ctx context.Context) int {
	var readyForDepart int

	if err := chromedp.Run(ctx,
		utils.GetIntFromElement(model.TEXT_FI_DEPART_AMOUNT, &readyForDepart),
	); err != nil {
		return 0
	}

	return readyForDepart
}
