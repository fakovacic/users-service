package users

import (
	"context"

	"github.com/fakovacic/users-service/internal/users/errors"
)

func (s *service) Create(ctx context.Context, m *User) (*User, error) {
	m.ID = s.uuidFunc().String()
	m.CreatedAt = s.timeFunc()
	m.UpdatedAt = s.timeFunc()

	err := s.store.Create(ctx, m)
	if err != nil {
		return nil, errors.Wrapf(err, "create user")
	}

	return m, nil
}
