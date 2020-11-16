package getter

import (
	"context"

	"github.com/scukonick/randomtiger/internal/app/db/models"
)

type Storage interface {
	GetTiger(ctx context.Context, chatID, userID int64) (*models.Tiger, error)
	CreateTiger(ctx context.Context, chatID, userID, stripes int64, username string) error
}
