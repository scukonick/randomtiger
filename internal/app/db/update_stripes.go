package db

import (
	"context"
	"fmt"
)

func (s *Storage) UpdateStripes(ctx context.Context, id, stripes int64) error {
	q := `UPDATE tigers SET 
                  stripes = $2, 
                  updated_at = NOW(),
                  enlarged_at = NOW()
        WHERE id = $1`
	_, err := s.conn.ExecContext(ctx, q, id, stripes)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	return nil
}
