package telegrambot

import (
	"botinfotime/internal/loging"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	OldTime    map[string][]string
	ListChatId map[int64]string      = make(map[int64]string)
	DataChan   chan map[int64]string = make(chan map[int64]string)
)

func StartInfoHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	ListChatId[chatID] = update.Message.From.FirstName

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Вас добавили в подписку о новом времени для Стефании",
	})
	loging.LogMessage(chatID, update.Message.From.FirstName, "Вас добавили в подписку о новом времени для Стефании.")
}
