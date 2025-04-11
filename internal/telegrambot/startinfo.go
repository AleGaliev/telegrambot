package telegrambot

import (
	"botinfotime/internal/keydb"
	"botinfotime/internal/loging"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartInfoHandler(ctx context.Context, b *bot.Bot, update *models.Update, keyDBAdrr string) {
	chatID := update.Message.Chat.ID
	KeyDbConfig := keydb.InitKeyDb(keyDBAdrr)
	listSurrveyUser, err := KeyDbConfig.GetValueKeyDb(ctx, "ListChat")
	if err != nil {
		SendErrMessage(ctx, b, err)
	}
	listSurrveyUser.AddByChatId(keydb.Survey{
		ChatId:    chatID,
		FirstName: update.Message.From.FirstName,
	})

	if err = KeyDbConfig.PushValueKeyDb(ctx, "ListChat", listSurrveyUser); err != nil {
		SendErrMessage(ctx, b, err)
		loging.LogMessage(0000000, "admin", "Не записана переменная в KeyDB")
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Вас добавили в подписку о новом времени для Стефании",
		})
		loging.LogMessage(chatID, update.Message.From.FirstName, "Вас добавили в подписку о новом времени для Стефании.")
	}
}
