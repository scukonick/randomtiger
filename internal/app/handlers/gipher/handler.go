package gipher

import (
	"context"
	"log"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/scukonick/giphy"
)

type Handler struct {
	bot         *tgbotapi.BotAPI
	giphyClient *giphy.Client
}

func NewHandler(bot *tgbotapi.BotAPI, giphyClient *giphy.Client) *Handler {
	return &Handler{
		bot:         bot,
		giphyClient: giphyClient,
	}
}

func (h *Handler) Handle(msg *tgbotapi.Message) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	chatID := msg.Chat.ID

	reply := tgbotapi.NewMessage(chatID, msg.Text)
	reply.ReplyToMessageID = msg.MessageID

	searchQ := giphy.SearchQuery{
		Q:     []string{"tiger"},
		Limit: 0,
	}

	s, err := h.giphyClient.Search(searchQ)
	if err != nil {
		log.Printf("oops: %+v", err)
		reply.Text = "something went terribly wrong"
		_, _ = h.bot.Send(reply)
		return
	}

	count := s.Pagination.TotalCount
	if count == 0 {
		reply.Text = "Нет больше тигров"
		_, _ = h.bot.Send(reply)
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
		log.Printf("oops: %+v", err)
		reply.Text = "something went terribly wrong"
		_, _ = h.bot.Send(reply)
		return
	}

	if len(s.Data) == 0 {
		reply.Text = "Нет больше тигров"
		_, _ = h.bot.Send(reply)
		return
	}

	reply.Text = s.Data[0].URL
	_, _ = h.bot.Send(reply)
}
