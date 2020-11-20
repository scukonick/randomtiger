package gipher

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/scukonick/randomtiger/internal/app/db/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/scukonick/giphy"
)

type Handler struct {
	bot         *tgbotapi.BotAPI
	giphyClient *giphy.Client
	storage     Storage
}

func NewHandler(bot *tgbotapi.BotAPI, giphyClient *giphy.Client, storage Storage) *Handler {
	return &Handler{
		bot:         bot,
		giphyClient: giphyClient,
		storage:     storage,
	}
}

func (h *Handler) Handle(msg *tgbotapi.Message) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	chatID := msg.Chat.ID
	userID := int64(msg.From.ID)
	username := msg.From.UserName

	reply := tgbotapi.NewMessage(chatID, msg.Text)
	reply.ReplyToMessageID = msg.MessageID

	tiger, err := h.storage.GetTiger(ctx, chatID, userID)
	switch {
	case errors.Is(err, models.ErrNotFound):
		reply.Text = "У тебя нет тигра, возьми его командой /get"
		_, err = h.bot.Send(reply)
		if err != nil {
			h.sendErr(reply, err)
		}
		return
	case err != nil:
		h.sendErr(reply, err)
		return
	}

	if tiger.Stripes <= 1 {
		reply.Text = fmt.Sprintf("У твоего тигра слишком мало полосок (%d)", tiger.Stripes)
	}

	searchQ := giphy.SearchQuery{
		Q:     []string{"tiger"},
		Limit: 0,
	}

	s, err := h.giphyClient.Search(searchQ)
	if err != nil {
		h.sendErr(reply, err)
		return
	}

	count := s.Pagination.TotalCount
	if count == 0 {
		reply.Text = "Нет больше тигров"
		_, err = h.bot.Send(reply)
		if err != nil {
			h.sendErr(reply, err)
		}
		return
	}
	if count > 500 {
		count = 500
	}

	rand.Seed(time.Now().UnixNano())
	offset := rand.Intn(count)
	searchQ.Offset = offset
	searchQ.Limit = 1

	s, err = h.giphyClient.Search(searchQ)
	if err != nil {
		h.sendErr(reply, err)
		return
	}

	if len(s.Data) == 0 {
		reply.Text = "Нет больше тигров"
		_, err = h.bot.Send(reply)
		if err != nil {
			h.sendErr(reply, err)
		}
		return
	}

	newStripes := tiger.Stripes - 1
	err = h.storage.UpdateStripes(ctx, tiger.ID, newStripes, username)
	if err != nil {
		h.sendErr(reply, err)
	}
	text := fmt.Sprintf("Ты заплатил 1 полоску за гифку, осталось: %d\n\n", newStripes)
	text += "Powered by GIPHY\n"
	text += s.Data[0].URL
	reply.Text = text

	_, err = h.bot.Send(reply)
	if err != nil {
		h.sendErr(reply, err)
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
