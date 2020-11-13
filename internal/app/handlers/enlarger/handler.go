package enlarger

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/scukonick/randomtiger/internal/app/db/models"
)

const enlargeCD = 3 * time.Hour

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
	userID := int64(msg.From.ID)

	reply := tgbotapi.NewMessage(chatID, msg.Text)
	reply.ReplyToMessageID = msg.MessageID

	tiger, err := h.storage.GetTiger(ctx, chatID, userID)
	switch {
	case errors.Is(err, models.ErrNotFound):
		log.Printf("err not found")
		reply.Text = "У тебя нет тигра, создай его командой /get"
		_, err = h.bot.Send(reply)
		if err != nil {
			log.Printf("failed to send: %+v", err)
		}
		return
	case err != nil:
		h.sendErr(reply, err)
		return
	}

	now := time.Now()
	when := tiger.EnlargedAt.Add(enlargeCD)
	if now.Before(when) {
		diff := when.Sub(now)
		reply.Text = fmt.Sprintf(
			"Ты недавно растил тигра, надо подождать %d ч. %d м.",
			int(diff.Hours()), int(diff.Minutes())%60)
		_, err := h.bot.Send(reply)
		if err != nil {
			log.Printf("failed to send: %+v", err)
		}
		return
	}

	stripes := tiger.Stripes
	addStripes := rand.Int63n(10) - 3
	stripes += addStripes
	if stripes <= 0 {
		stripes = 1
	}

	err = h.storage.UpdateStripes(ctx, tiger.ID, stripes)
	if err != nil {
		h.sendErr(reply, err)
		return
	}

	text := fmt.Sprintf("Количество полосок на твоем тигре выросло, теперь их %d", stripes)
	if addStripes == 0 {
		text = fmt.Sprintf("Количество полосок не изменилось, их %d", stripes)
	} else if addStripes < 0 {
		text = fmt.Sprintf("Количество полосок уменьшилось, теперь  их %d", stripes)
	}

	reply.Text = text

	_, err = h.bot.Send(reply)
	if err != nil {
		log.Printf("something went wrong: %+v", err)
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
