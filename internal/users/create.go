package users

import (
	"context"

	"github.com/fakovacic/users-service/internal/users/errors"
)

func ValidateCreate(ctx context.Context, m *User) error {
	if m.Email == "" {
		return errors.Wrap("email is required")
	}

	return nil
}

func (s *service) Create(ctx context.Context, m *User) (*User, error) {
	err := ValidateCreate(ctx, m)
	if err != nil {
		return nil, err
	}

	m.ID = s.uuidFunc().String()
	m.CreatedAt = s.timeFunc()
	m.UpdatedAt = s.timeFunc()

	if m.Password != "" {
		m.Password = HashAndSalt(m.Password)
	}

	err = s.store.Create(ctx, m)
	if err != nil {
		return nil, errors.Wrapf(err, "create user")
	}

	return m, nil
}
