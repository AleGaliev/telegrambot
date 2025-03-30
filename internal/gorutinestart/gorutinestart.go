package gorutinestart

import (
	"botinfotime/internal/telegrambot"
	"context"
	"fmt"
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
		ticker := time.NewTicker(10 * time.Second) // Таймер на 10 секунд
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				listSurrveyChatID := <-telegrambot.DataChan
				fmt.Println(listSurrveyChatID)
				telegrambot.SendBroadcastMessage(ctx, b, listSurrveyChatID)
			case <-telegrambot.StopChan:
				// Останавливаем горутину
				return
			}
		}
	}()
}
