package db

import (
	"context"
	"fmt"
)

func (s *Storage) CreateTiger(ctx context.Context, chatID, userID, stripes int64) error {
	q := `INSERT INTO tigers 
    (chat_id, user_id, stripes)
    VALUES ($1, $2, $3)`

	_, err := s.conn.ExecContext(ctx, q, chatID, userID, stripes)
	if err != nil {
		return fmt.Errorf("failed to insert to db: %w", err)
	}

	return nil
}
