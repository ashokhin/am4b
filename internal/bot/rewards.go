package bot

import (
	"context"
	"log/slog"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/ashokhin/am4bot/internal/utils"

	"github.com/chromedp/chromedp"
)

// claimRewards checks for available biweekly gifts and claims them.
func (b *Bot) claimRewards(ctx context.Context) error {
	var hasRewards bool

	slog.Info("check duty free rewards")

	// check "Free Reward" icon for the "Bonus" menu
	hasRewards = utils.IsElementVisible(ctx, model.ICON_FREE_REWARDS)

	slog.Debug("rewards available", "has_rewards", hasRewards)

	if !hasRewards {
		slog.Info("no rewards available to claim")

		return nil
	}

	slog.Debug("open pop-up window", "window", "bonus")

	// open "Bonus" pop-up
	if err := chromedp.Run(ctx,
		utils.ClickElement(model.BUTTON_MAIN_BONUS),
	); err != nil {
		slog.Warn("error in Bot.claimRewards > open bonus window", "error", err)

		return err
	}

	defer utils.DoClickElement(ctx, model.BUTTON_COMMON_CLOSE_POPUP)

	slog.Debug("navigate to Duty Free tab and claim rewards")

	if err := chromedp.Run(ctx,
		utils.ClickElement(model.BUTTON_BONUS_DUTY_FREE_TAB),
		chromedp.WaitReady(model.BUTTON_BONUS_CLAIM_GIFT, chromedp.ByQuery),
		utils.ClickElement(model.BUTTON_BONUS_CLAIM_GIFT),
	); err != nil {
		slog.Warn("error in Bot.claimRewards > claim gifts", "error", err)

		return err
	}

	slog.Info("rewards successfully claimed")

	return nil
}
