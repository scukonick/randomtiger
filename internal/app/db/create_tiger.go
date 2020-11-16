package db

import (
	"context"
	"fmt"
)

func (s *Storage) CreateTiger(ctx context.Context, chatID, userID, stripes int64, username string) error {
	q := `INSERT INTO tigers 
    (chat_id, user_id, stripes, username)
    VALUES ($1, $2, $3, $4)`

	_, err := s.conn.ExecContext(ctx, q, chatID, userID, stripes, username)
	if err != nil {
		return fmt.Errorf("failed to insert to db: %w", err)
	}

	return nil
}
