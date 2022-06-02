package users

import (
	"context"

	"github.com/fakovacic/users-service/internal/users/errors"
)

func (s *service) Delete(ctx context.Context, id string) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "delete user")
	}

	return nil
}
