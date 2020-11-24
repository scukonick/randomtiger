package db

import (
	"context"
	"fmt"
)

func (s *Storage) EnlargeStripes(ctx context.Context, id, stripes int64, username string) error {
	q := `UPDATE tigers SET 
                  stripes = $2, 
                  username = $3,
                  updated_at = NOW(),
                  enlarged_at = NOW()
        WHERE id = $1`
	_, err := s.conn.ExecContext(ctx, q, id, stripes, username)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	return nil
}
