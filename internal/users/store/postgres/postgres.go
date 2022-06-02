package postgres

import (
	"database/sql"

	"github.com/fakovacic/users-service/internal/users"
)

func New(db *sql.DB) users.Store {
	s := &store{
		db: db,
	}

	return s
}

type store struct {
	db *sql.DB
}

func (s *store) DB() *sql.DB {
	return s.db
}
