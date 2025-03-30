package telegrambot

import (
	"botinfotime/internal/app"
	"botinfotime/internal/loging"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	userRights []int64   = []int64{224268678}
	StopChan   chan bool = make(chan bool)
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
	if stopChan != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Получение информации о времени у Стефании уже запущено.",
		})
		loging.LogMessage(chatID, update.Message.From.FirstName, "Получение информации о времени у Стефании уже запущено.")
		return
	}

	// Создаем канал для управления горутиной
	stopChan = make(chan struct{})

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Запущен опрос свободного времени с Стефании с переодичностью 20 минут .",
	})
	loging.LogMessage(chatID, update.Message.From.FirstName, "Запущен опрос свободного времени с Стефании с переодичностью 20 минут .")
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

func SendBroadcastMessage(ctx context.Context, b *bot.Bot, listSurrveyUser map[int64]string) {
	var message string
	OldTime, message = app.AppRun(OldTime)
	//if message != "" {
	for chatID, firstName := range listSurrveyUser {
		loging.LogMessage(chatID, firstName, "прогон пошел")
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   message,
		})
		if err != nil {
			fmt.Printf("Ошибка при отправке сообщения: %v\n", err)
		} else {
			loging.LogMessage(chatID, firstName, message)
		}
	}
	//}
}
