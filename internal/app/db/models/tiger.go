package models

import "time"

type Tiger struct {
	ID         int64     `db:"id"`
	ChatID     int64     `db:"chat_id"`
	UserID     int64     `db:"user_id"`
	Stripes    int64     `db:"stripes"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	EnlargedAt time.Time `db:"enlarged_at"`
}
