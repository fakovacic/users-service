package postgres

import (
	"context"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/lib/pq"
)

func (s *store) List(ctx context.Context, meta *users.Meta) (int64, []*users.User, error) {
	var count int64

	err := s.db.QueryRowContext(ctx, `SELECT COUNT(id) FROM users`).Scan(&count)
	if err != nil {
		p, ok := err.(*pq.Error)
		if ok {
			err = errors.Wrapf(err, " database error: %s", p.Code.Class().Name())
		}

		return 0, nil, errors.Wrapf(err, "get user list count")
	}

	if count == 0 {
		var list []*users.User

		return 0, list, nil
	}

	var list []*users.User

	offset := (meta.Page * meta.Limit) - meta.Limit

	rows, err := s.db.QueryContext(ctx,
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
		 OFFSET $1
		 LIMIT $2`,
		offset,
		meta.Limit,
	)
	if err != nil {
		p, ok := err.(*pq.Error)
		if ok {
			err = errors.Wrapf(err, " database error: %s", p.Code.Class().Name())
		}

		return 0, nil, errors.Wrapf(err, "get user list")
	}

	defer rows.Close()

	if rows.Err() != nil {
		return 0, nil, errors.InternalWrap(rows.Err(), "rows error")
	}

	for rows.Next() {
		var usr *users.User

		err = rows.Scan(
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
			return 0, nil, errors.Wrapf(err, "scan user")
		}

		list = append(list, usr)
	}

	return count, list, nil
}
