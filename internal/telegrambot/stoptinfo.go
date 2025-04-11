package telegrambot

import (
	"botinfotime/internal/keydb"
	"botinfotime/internal/loging"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StopInfoHandler(ctx context.Context, b *bot.Bot, update *models.Update, keyDBAdrr string) {
	chatID := update.Message.Chat.ID
	KeyDbConfig := keydb.InitKeyDb(keyDBAdrr)
	listSurrveyUser, err := KeyDbConfig.GetValueKeyDb(ctx, "ListChat")
	if err != nil {
		SendErrMessage(ctx, b, err)
	}

	listSurrveyUser.RemoveByChatId(chatID)
	if err = KeyDbConfig.PushValueKeyDb(ctx, "ListChat", listSurrveyUser); err != nil {
		SendErrMessage(ctx, b, err)
		loging.LogMessage(0000000, "", "Не записана переменная в KeyDB")
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Вас удалили с подписки времени для Стефании",
		})
		loging.LogMessage(chatID, update.Message.From.FirstName, "Вас удалили с подписки времени для Стефании.")
	}
}
