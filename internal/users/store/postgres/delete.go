package postgres

import (
	"context"

	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/lib/pq"
)

func (s *store) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM users WHERE id = $1`,
		id,
	)
	if err != nil {
		p, ok := err.(*pq.Error)
		if ok {
			err = errors.Wrapf(err, " database error: %s", p.Code.Class().Name())
		}

		return errors.Wrapf(err, "delete user")
	}

	return nil
}
