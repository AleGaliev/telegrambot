package main

import (
	"botinfotime/internal/config"
	"botinfotime/internal/gorutinestart"
	"botinfotime/internal/telegrambot"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"os"
	"os/signal"
)

func main() {
	appConfig := config.NewAppConfig()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	TimeNowHandler := func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		telegrambot.TimeNowHandler(ctx, bot, update, appConfig)
	}
	StartInfoHandler := func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		telegrambot.StartInfoHandler(ctx, bot, update, appConfig.KeyDB)
	}
	StopInfoHandler := func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		telegrambot.StopInfoHandler(ctx, bot, update, appConfig.KeyDB)
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(telegrambot.DefaultHandler),
	}
	b, err := bot.New(os.Getenv("TELEGRAM_TOKEN"), opts...)

	if nil != err {
		panic(err)
	}

	gorutinestart.GorutineStart(ctx, b, appConfig)
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	telegrambot.SendStartMessage(ctx, b)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, telegrambot.HelloHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/startsurvey", bot.MatchTypeExact, telegrambot.StartSurveyHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/stoptsurvey", bot.MatchTypeExact, telegrambot.StopSurveyHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/timenow", bot.MatchTypeExact, TimeNowHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/startinfo", bot.MatchTypeExact, StartInfoHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/stoptinfo", bot.MatchTypeExact, StopInfoHandler)

	b.Start(ctx)
}
