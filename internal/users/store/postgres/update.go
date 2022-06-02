package postgres

import (
	"context"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/lib/pq"
)

func (s *store) Update(ctx context.Context, id string, model *users.User) error {
	_, err := s.db.ExecContext(ctx, `UPDATE users 
		SET 
			first_name = $1, 
			last_name = $2, 
			nickname = $3, 
			email = $4, 
			country = $5, 
			updated_at = $6
		WHERE id = $7`,
		model.FirstName,
		model.LastName,
		model.Nickname,
		model.Email,
		model.Country,
		model.UpdatedAt,
		id,
	)
	if err != nil {
		p, ok := err.(*pq.Error)
		if ok {
			err = errors.Wrapf(err, " database error: %s", p.Code.Class().Name())
		}

		return errors.Wrapf(err, "update user")
	}

	return nil
}
