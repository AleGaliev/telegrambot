package telegrambot

import (
	"botinfotime/internal/loging"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StopInfoHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	delete(ListChatId, chatID)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Вас добавили в подписку о новом времени для Стефании",
	})
	loging.LogMessage(chatID, update.Message.From.FirstName, "Вас добавили в подписку о новом времени для Стефании.")
}
