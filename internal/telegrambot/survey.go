package telegrambot

import (
	"botinfotime/internal/app"
	"botinfotime/internal/config"
	"botinfotime/internal/keydb"
	"botinfotime/internal/loging"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	userRights  []int64 = []int64{224268678, 211348295}
	StartSurvey bool    = true
	StopChan            = make(chan struct{})
)

func StartSurveyHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	// если нет прав на запуск опросника
	if !checkUserRights(chatID) {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "У вас нет прав на запуск этой команды",
		})
		loging.LogMessage(chatID, update.Message.From.FirstName, "У вас нет прав на запуск этой команды")
		return
	}
	// Если канал уже существует, значит отправка уже запущена
	if StartSurvey == true {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Получение информации о времени у Стефании уже запущено.",
		})
		loging.LogMessage(chatID, update.Message.From.FirstName, "Получение информации о времени у Стефании уже запущено.")
		return
	}

	// Создаем канал для управления горутиной
	StopChan = make(chan struct{})
	StartSurvey = true
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Запущен опрос свободного времени с Стефании с периодичностью 10 минут .",
	})
	loging.LogMessage(chatID, update.Message.From.FirstName, "Запущен опрос свободного времени с Стефании с периодичностью ]0 минут .")
}

func StopSurveyHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	// если нет прав на запуск опросника
	if !checkUserRights(chatID) {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "У вас нет прав на запуск этой команды",
		})
		loging.LogMessage(chatID, update.Message.From.FirstName, "У вас нет прав на запуск этой команды")
		return
	}
	// Если канал уже существует, значит отправка уже запущена
	if StartSurvey == false {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Опрос свободного времени у Стефании остановлен",
		})
		loging.LogMessage(chatID, update.Message.From.FirstName, "Опрос свободного времени у Стефании остановлен")
		return
	}

	// Создаем канал для управления горутиной
	close(StopChan)
	StopChan = nil
	StartSurvey = false
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Опрос свободного времени у Стефании остановлен",
	})
	loging.LogMessage(chatID, update.Message.From.FirstName, "Опрос свободного времени у Стефании остановлен")
}

func checkUserRights(chatID int64) bool {
	check := false
	for _, user := range userRights {
		if user == chatID {
			check = true
		}
	}
	return check
}

func SendBroadcastMessage(ctx context.Context, b *bot.Bot, appConfig config.AppConfig) {
	message, _, err := app.RunGetChangeTime(appConfig)
	if err != nil {
		SendErrMessage(ctx, b, err)
	}

	if message != "" {
		KeyDbConfig := keydb.InitKeyDb(appConfig.KeyDB)
		listSurrveyUser, err := KeyDbConfig.GetValueKeyDb(ctx, "ListChat")
		if err != nil {
			SendErrMessage(ctx, b, err)
		}
		for _, survey := range listSurrveyUser.Surveys {
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: survey.ChatId,
				Text:   message,
			})
			if err != nil {
				fmt.Printf("Ошибка при отправке сообщения: %v\n", err)
			} else {
				loging.LogMessage(survey.ChatId, survey.FirstName, message)
			}
		}
	}
}
