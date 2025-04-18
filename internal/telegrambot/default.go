package telegrambot

import (
	"botinfotime/internal/loging"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	startMessage string = "/hello - тестовый запрос \n/timenow - узнать даты и время на запись у Стефании \n/startsurvey - запустить опросы \n/stoptsurvey - остановить запросы \n/startinfo - запускает опрос для получения свободных слотов у Стефании \n/stoptinfo - останавливает опрос для получения свободных слотов у Стефании "
	AdminChatID  int64  = 224268678
)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   startMessage,
	})
	loging.LogMessage(update.Message.Chat.ID, bot.EscapeMarkdown(update.Message.From.FirstName), startMessage)
}

func SendErrMessage(ctx context.Context, b *bot.Bot, err error) {
	message := err.Error()
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: AdminChatID,
		Text:   message,
	})
	loging.LogMessage(AdminChatID, "admin", message)
}

func SendStartMessage(ctx context.Context, b *bot.Bot) {
	message := "Бот запущен"
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: AdminChatID,
		Text:   message,
	})
	loging.LogMessage(AdminChatID, "admin", message)
}
