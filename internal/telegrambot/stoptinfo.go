package telegrambot

import (
	"botinfotime/internal/loging"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StopInfoHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID

	// Если канал не существует, значит отправка не запущена
	if stopChan == nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Опрос свободного времени у Стефании не был запущен",
		})
		loging.LogMessage(chatID, update.Message.From.FirstName, "Опрос свободного времени у Стефании не был запущен")
		return
	}

	// Закрываем канал, чтобы остановить горутину
	close(stopChan)
	stopChan = nil

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Опрос свободного времени у Стефании остановлен",
	})
	loging.LogMessage(chatID, update.Message.From.FirstName, "Опрос свободного времени у Стефании остановлен")
}
