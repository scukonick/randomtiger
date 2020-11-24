package enlarger

import (
	"context"

	"github.com/scukonick/randomtiger/internal/app/db/models"
)

type Storage interface {
	GetTiger(ctx context.Context, chatID, userID int64) (*models.Tiger, error)
	EnlargeStripes(ctx context.Context, id, stripes int64, username string) error
}
