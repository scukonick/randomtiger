package db

import (
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	conn *sqlx.DB
}

func NewStorage(conn *sqlx.DB) *Storage {
	return &Storage{
		conn: conn,
	}
}
