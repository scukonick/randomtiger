package getter

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/scukonick/randomtiger/internal/app/db/models"

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

func (g *Handler) Handle(msg *tgbotapi.Message) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	chatID := msg.Chat.ID
	userID := int64(msg.From.ID)

	reply := tgbotapi.NewMessage(chatID, msg.Text)
	reply.ReplyToMessageID = msg.MessageID

	tiger, err := g.storage.GetTiger(ctx, chatID, userID)
	switch {
	case err == nil:
		reply.Text = fmt.Sprintf("У тебя уже есть тигр, кол-во полосок: %d", tiger.Stripes)
	case errors.Is(err, models.ErrNotFound):
		stripes := rand.Int63n(5) + 1
		err = g.storage.CreateTiger(ctx, chatID, userID, stripes)
		if err != nil {
			reply.Text = "Что-то пошло не так, сорян"
			break
		}
		reply.Text = fmt.Sprintf("Ты взял себе тигра, количество полосок: %d", stripes)
	default: // err != nil
		reply.Text = "Что-то пошло не так, сорри"
	}

	_, err = g.bot.Send(reply)
	if err != nil {
		log.Printf("err: %+v", err)
	}
}
