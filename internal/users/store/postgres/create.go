package postgres

import (
	"context"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/lib/pq"
)

func (s *store) Create(ctx context.Context, model *users.User) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO users (id, first_name, last_name, nickname, password, email, country, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		model.ID,
		model.FirstName,
		model.LastName,
		model.Nickname,
		model.Password,
		model.Email,
		model.Country,
		model.CreatedAt,
		model.UpdatedAt,
	)
	if err != nil {
		p, ok := err.(*pq.Error)
		if ok {
			err = errors.Wrapf(err, " database error: %s", p.Code.Class().Name())
		}

		return errors.Wrapf(err, "create user")
	}

	return nil
}
