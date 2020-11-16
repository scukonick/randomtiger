package db

import (
	"context"
	"fmt"

	"github.com/scukonick/randomtiger/internal/app/db/models"
)

func (s *Storage) GetTopTigers(ctx context.Context, chatID, limit int64) ([]models.Tiger, error) {
	q := `SELECT id, chat_id, user_id, stripes, username, 
       created_at, updated_at, enlarged_at
       FROM tigers WHERE chat_id = $1
		ORDER BY stripes DESC LIMIT $2`

	resp := make([]models.Tiger, 0, limit)

	err := s.conn.SelectContext(ctx, &resp, q, chatID, limit)
	if err != nil {
		return nil, fmt.Errorf("select failed: %w", err)
	}

	return resp, nil
}
