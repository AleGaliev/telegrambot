package telegrambot

import (
	"botinfotime/internal/loging"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	startMessage string = "/hello - тестовый запрос \n/timenow - узнать даты и время на запись у Стефании \n/startsurvey - запустить опросы \n/startinfo - запускает опрос для получения свободных слотов у Стефании \n/stoptinfo - останавливает опрос для получения свободных слотов у Стефании "
)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   startMessage,
	})
	loging.LogMessage(update.Message.Chat.ID, bot.EscapeMarkdown(update.Message.From.FirstName), startMessage)
}
