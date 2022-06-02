package postgres

import (
	"context"
	"database/sql"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/lib/pq"
)

func (s *store) Get(ctx context.Context, id string) (*users.User, error) {
	var usr users.User

	err := s.db.QueryRowContext(ctx,
		`SELECT 
			id, 
			first_name,
			last_name,
			nickname,
			password,
			email,
			country,
			created_at,
			updated_at
		 FROM users 
		 WHERE id=$1`,
		id,
	).Scan(
		&usr.ID,
		&usr.FirstName,
		&usr.LastName,
		&usr.Nickname,
		&usr.Password,
		&usr.Email,
		&usr.Country,
		&usr.CreatedAt,
		&usr.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFound("not found user")
		}

		p, ok := err.(*pq.Error)
		if ok {
			err = errors.Wrapf(err, " database error: %s", p.Code.Class().Name())
		}

		return nil, errors.Wrapf(err, "get user")
	}

	return &usr, nil
}
