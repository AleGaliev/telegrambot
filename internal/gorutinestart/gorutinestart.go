package gorutinestart

import (
	"botinfotime/internal/config"
	"botinfotime/internal/telegrambot"
	"context"
	"github.com/go-telegram/bot"
	"time"
)

func GorutineStart(ctx context.Context, b *bot.Bot, appConfig config.AppConfig) {
	go func() {
		for {
			select {
			default:
				if telegrambot.StartSurvey == true {
					telegrambot.SendBroadcastMessage(ctx, b, appConfig)
				}
			case <-telegrambot.StopChan:
				// Останавливаем горутину
				return
			}

			time.Sleep(10 * time.Minute)
		}
	}()
}
