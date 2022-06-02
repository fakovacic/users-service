package users

import (
	"context"

	"github.com/fakovacic/users-service/internal/users/errors"
)

func (s *service) Update(ctx context.Context, id string, m *User, fields []string) (*User, error) {
	if id == "" {
		return nil, errors.BadRequest("id is empty")
	}

	if len(fields) == 0 {
		return nil, errors.BadRequest("update fields empty")
	}

	for i := range fields {
		val, ok := UserUpdateFields[fields[i]]
		if !ok {
			return nil, errors.BadRequest("field '%s' not exist", fields[i])
		}

		if !val {
			return nil, errors.BadRequest("field '%s' cannot be updated", fields[i])
		}
	}

	model, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "get user")
	}

	for i := range fields {
		switch fields[i] {
		case UserField.FirstName:
			model.FirstName = m.FirstName
		case UserField.LastName:
			model.LastName = m.LastName
		case UserField.Nickname:
			model.Nickname = m.Nickname
		case UserField.Password:
			model.Password = m.Password
		case UserField.Email:
			model.Email = m.Email
		case UserField.Country:
			model.Country = m.Country
		}
	}

	model.UpdatedAt = s.timeFunc()

	err = s.store.Update(ctx, id, model)
	if err != nil {
		return nil, errors.Wrapf(err, "update user")
	}

	return model, nil
}
