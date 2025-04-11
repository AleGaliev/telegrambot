package telegrambot

import (
	"botinfotime/internal/app"
	"botinfotime/internal/config"
	"botinfotime/internal/loging"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func TimeNowHandler(ctx context.Context, b *bot.Bot, update *models.Update, appConfig config.AppConfig) {
	message, _, err := app.RunGetTimeNow(appConfig)
	if err != nil {
		SendErrMessage(ctx, b, err)
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   message,
		//ParseMode: models.ParseModeMarkdown,
	})

	loging.LogMessage(update.Message.Chat.ID, update.Message.From.FirstName, message)
}
