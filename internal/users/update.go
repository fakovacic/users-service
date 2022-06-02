package users

import (
	"context"

	"github.com/fakovacic/users-service/internal/users/errors"
)

func (s *service) Update(ctx context.Context, m *User, fields []string) (*User, error) {
	if len(fields) == 0 {
		return m, nil
	}

	for i := range fields {
		val, ok := UserUpdateFields[fields[i]]
		if !ok {
			return nil, errors.BadRequest("field not exist")
		}

		if !val {
			return nil, errors.BadRequest("field cannot be updated")
		}
	}

	fields = append(fields, UserField.UpdatedAt)

	m.UpdatedAt = s.timeFunc()

	model, err := s.store.Update(ctx, m, fields)
	if err != nil {
		return nil, errors.Wrapf(err, "update user")
	}

	return model, nil
}
