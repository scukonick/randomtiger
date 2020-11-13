package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type RouterInterface interface {
	Route(update tgbotapi.Update)
}

type App struct {
	bot    *tgbotapi.BotAPI
	router RouterInterface
}

func NewApp(bot *tgbotapi.BotAPI, router RouterInterface) *App {
	return &App{
		bot:    bot,
		router: router,
	}
}

func (app *App) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := app.bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}
	for update := range updates {
		app.router.Route(update)
	}
}
