package bot

import (
	"context"
	"log/slog"
	"math"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/ashokhin/am4bot/internal/utils"
	"github.com/chromedp/chromedp"
)

// depart handles the departure of all available aircraft from the fleet.
func (b *Bot) depart(ctx context.Context) error {
	slog.Info("depart all available aircraft")

	aircraftReadyForDepart := b.getReadyForDepart(ctx)
	// calculate maximum retries, because the "Depart" button may process only 20 aircraft at a time
	// also to avoid infinite loops when aircraft has been grounded
	maxRetries := int(math.Round(float64(aircraftReadyForDepart)/20) + 1)

	for (aircraftReadyForDepart > 0) && (maxRetries > 0) {
		var availableAfterDepart int

		slog.Debug("depart available aircraft", "ready to depart", aircraftReadyForDepart, "depart retries", maxRetries)

		utils.DoClickElement(ctx, model.BUTTON_FI_DEPART_ALL)
		availableAfterDepart = b.getReadyForDepart(ctx)

		slog.Info("aircraft departed", "count", (aircraftReadyForDepart - availableAfterDepart))

		aircraftReadyForDepart = availableAfterDepart

		maxRetries--

		// try to buy fuel after each depart iteration
		if err := b.fuel(ctx); err != nil {
			slog.Error("failed to refuel during depart iteration", "error", err)
		}
	}

	return nil
}

// getReadyForDepart retrieves the number of aircraft ready for departure from the fleet interface.
func (b *Bot) getReadyForDepart(ctx context.Context) int {
	var readyForDepart int

	if err := chromedp.Run(ctx,
		utils.GetIntFromElement(model.TEXT_FI_DEPART_AMOUNT, &readyForDepart),
	); err != nil {
		slog.Debug("the 'Depart' amount element not found, assuming 0 ready for depart", "error", err)

		return 0
	}

	return readyForDepart
}
