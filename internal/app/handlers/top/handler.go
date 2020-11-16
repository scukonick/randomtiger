package top

import (
	"context"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Handler struct {
	bot     *tgbotapi.BotAPI
	storage Storage
}

func NewHandler(bot *tgbotapi.BotAPI, storage Storage) *Handler {
	return &Handler{
		bot:     bot,
		storage: storage,
	}
}

func (h *Handler) Handle(msg *tgbotapi.Message) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	chatID := msg.Chat.ID

	reply := tgbotapi.NewMessage(chatID, msg.Text)
	reply.ReplyToMessageID = msg.MessageID

	tigers, err := h.storage.GetTopTigers(ctx, chatID, 10)
	if err != nil {
		h.sendErr(reply, err)
		return
	}

	resp := "Топ 10 тигров:\n\n"
	for i, t := range tigers {
		pos := i + 1
		resp += fmt.Sprintf("%d) %s - %d полосок\n", pos, t.Username, t.Stripes)
	}

	reply.Text = resp

	_, err = h.bot.Send(reply)
	if err != nil {
		log.Printf("failed to send reply: %+v", err)
	}
}

func (h *Handler) sendErr(reply tgbotapi.MessageConfig, err error) {
	log.Printf("sending error: %+v", err)
	reply.Text = "something went wrong"
	_, err = h.bot.Send(reply)
	if err != nil {
		log.Printf("failed to send err: %+v", err)
	}
}
