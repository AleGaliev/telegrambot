package main

import (
	"botinfotime/internal/gorutinestart"
	"botinfotime/internal/telegrambot"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(telegrambot.DefaultHandler),
	}
	//b, err := bot.New("7721087218:AAGdZAqI-ThgoqbT-tDd11M3Hj9j6L8j2TY", opts...)
	b, err := bot.New(os.Getenv("TELEGRAM_TOKEN"), opts...)

	if nil != err {
		panic(err)
	}

	gorutinestart.GorutineStart(ctx, b)
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, telegrambot.HelloHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/startsurvey", bot.MatchTypeExact, telegrambot.StartSurveyHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/timenow", bot.MatchTypeExact, telegrambot.TimeNowHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/startinfo", bot.MatchTypeExact, telegrambot.StartInfoHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/stoptinfo", bot.MatchTypeExact, telegrambot.StopInfoHandler)

	b.Start(ctx)
}
