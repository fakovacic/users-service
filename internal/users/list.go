package users

import (
	"context"

	"github.com/fakovacic/users-service/internal/users/errors"
)

func (s *service) List(ctx context.Context, m *Meta) (*Meta, []*User, error) {
	count, list, err := s.store.List(ctx, m)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "get users list")
	}

	m.Count = count
	m.CalcPages()

	return m, list, nil
}
