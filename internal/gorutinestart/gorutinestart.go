package gorutinestart

import (
	"botinfotime/internal/telegrambot"
	"context"
	"github.com/go-telegram/bot"
	"time"
)

type Gorutinestart struct {
	ctx             context.Context
	b               *bot.Bot
	chatID          int64
	firstName       string
	listSurrveyUser chan []int64
}

func GorutineStart(ctx context.Context, b *bot.Bot) {
	go func() {
		for {
			select {
			default:
				if telegrambot.StartSurvey == true {
					telegrambot.SendBroadcastMessage(ctx, b, telegrambot.ListChatId)
				}
			case <-telegrambot.StopChan:
				// Останавливаем горутину
				return
			}

			time.Sleep(10 * time.Second)
		}
	}()
}
