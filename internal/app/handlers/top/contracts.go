package top

import (
	"context"

	"github.com/scukonick/randomtiger/internal/app/db/models"
)

type Storage interface {
	GetTopTigers(ctx context.Context, chatID, limit int64) ([]models.Tiger, error)
}
