package bot

import (
	"context"
	"log/slog"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/ashokhin/am4bot/internal/utils"
	"github.com/chromedp/chromedp"
)

// auth performs authentication on the target website using credentials from the bot configuration.
func (b *Bot) auth(ctx context.Context) error {
	slog.Debug("auth", "url", b.Conf.Url, "user", b.Conf.User)

	if err := chromedp.Run(ctx,
		chromedp.Navigate(b.Conf.Url),
		chromedp.Click(model.BUTTON_PLAY_NOW, chromedp.ByQuery),
		chromedp.Click(model.BUTTON_LOGIN, chromedp.ByQuery),
		chromedp.WaitReady(model.TEXT_FIELD_LOGIN, chromedp.ByQuery),
		chromedp.SendKeys(model.TEXT_FIELD_LOGIN, b.Conf.User, chromedp.ByQuery),
		chromedp.SendKeys(model.TEXT_FIELD_PASSWORD, b.Conf.GetPassword(), chromedp.ByQuery),
		chromedp.Click(model.BUTTON_AUTH, chromedp.ByQuery),
		chromedp.WaitNotVisible(model.OVERLAY_LOADING, chromedp.ByQuery),
		utils.RefreshPage(),
	); err != nil {
		slog.Warn("error in bot.auth", "error", err)

		return err
	}

	return nil
}
