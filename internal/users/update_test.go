package users_test

import (
	"context"
	"testing"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/fakovacic/users-service/internal/users/mocks"
	"github.com/matryer/is"
)

func TestUpdate(t *testing.T) {
	cases := []struct {
		it           string
		id           string
		model        *users.User
		updateFields []string

		// Store
		userGetInput  string
		userGetResult *users.User
		userGetError  error

		userUpdateInputID    string
		userUpdateInputModel *users.User
		userUpdateError      error

		expectedError  string
		expectedResult *users.User
	}{
		{
			it: "it update and return user",
			id: "mock-id",
			model: &users.User{
				Email: "mock-email",
			},
			updateFields: []string{
				"email",
			},
			userGetInput: "mock-id",
			userGetResult: &users.User{
				ID: "mock-id",
			},
			userUpdateInputID: "mock-id",
			userUpdateInputModel: &users.User{
				ID:        "mock-id",
				Email:     "mock-email",
				UpdatedAt: GenTime(),
			},
			expectedResult: &users.User{
				ID:        "mock-id",
				Email:     "mock-email",
				UpdatedAt: GenTime(),
			},
		},
		{
			it: "it returns error on store Update",
			id: "mock-id",
			model: &users.User{
				Email: "mock-email",
			},
			updateFields: []string{
				"email",
			},

			userGetInput: "mock-id",
			userGetResult: &users.User{
				ID: "mock-id",
			},
			userUpdateInputID: "mock-id",
			userUpdateInputModel: &users.User{
				ID:        "mock-id",
				Email:     "mock-email",
				UpdatedAt: GenTime(),
			},
			userUpdateError: errors.Wrap("mock-error"),
			expectedError:   "update user: mock-error",
		},
		{
			it: "it returns error on store Get",
			id: "mock-id",
			model: &users.User{
				Email: "mock-email",
			},
			updateFields: []string{
				"email",
			},
			userGetInput:  "mock-id",
			userGetError:  errors.Wrap("mock-error"),
			expectedError: "get user: mock-error",
		},
		{
			it: "it returns error on check fields",
			id: "mock-id",
			updateFields: []string{
				"id",
			},
			expectedError: "field 'id' cannot be updated",
		},
		{
			it: "it returns error on field not found",
			id: "mock-id",
			updateFields: []string{
				"mock",
			},
			expectedError: "field 'mock' not exist",
		},
		{
			it:            "it returns error on empty fields",
			id:            "mock-id",
			expectedError: "update fields empty",
		},
		{
			it:            "it returns error on empty id",
			expectedError: "id is empty",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			store := &mocks.StoreMock{
				GetFunc: func(ctx context.Context, id string) (*users.User, error) {
					checkIs.Equal(id, tc.userGetInput)

					return tc.userGetResult, tc.userGetError
				},
				UpdateFunc: func(ctx context.Context, id string, model *users.User) error {
					checkIs.Equal(id, tc.userUpdateInputID)
					checkIs.Equal(model, tc.userUpdateInputModel)

					return tc.userUpdateError
				},
			}

			service := users.New(
				users.NewConfig(""),
				store,
				GenTime,
				nil,
			)

			res, err := service.Update(context.Background(), tc.id, tc.model, tc.updateFields)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(res, tc.expectedResult)
		})
	}
}
