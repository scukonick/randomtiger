package app

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Router struct {
	bot    *tgbotapi.BotAPI
	routes map[string]msgHandler
}

func NewRouter(bot *tgbotapi.BotAPI) *Router {
	return &Router{
		bot:    bot,
		routes: make(map[string]msgHandler, 10),
	}
}

type msgHandler interface {
	Handle(message *tgbotapi.Message)
}

func (r *Router) AddCmdRoute(msg string, handler msgHandler) {
	r.routes[msg] = handler
}

func (r *Router) Route(update tgbotapi.Update) {
	if update.Message == nil { // ignore any non-Message Updates
		return
	}
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	log.Printf("%v", update.Message.Chat)

	if update.Message.IsCommand() {
		cmd := update.Message.Command()
		h, ok := r.routes[cmd]
		if !ok {
			log.Printf("unknown command: %s", cmd)
			return
		}
		h.Handle(update.Message)
		return
	}

}
