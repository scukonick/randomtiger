package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/scukonick/randomtiger/internal/app/db/models"
)

func (s *Storage) GetTiger(ctx context.Context, chatID, userID int64) (*models.Tiger, error) {
	q := `SELECT id, chat_id, user_id, stripes, username, 
	created_at, updated_at, enlarged_at
	FROM tigers WHERE chat_id = $1 and user_id = $2`

	resp := &models.Tiger{}
	err := s.conn.GetContext(ctx, resp, q, chatID, userID)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("select failed: %w", err)
	}

	return resp, nil
}
