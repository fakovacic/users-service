package store

import (
	"database/sql"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/store/postgres"
)

func NewStore(db *sql.DB) users.Store {
	return postgres.New(db)
}
