package users_test

import (
	"context"
	"testing"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/fakovacic/users-service/internal/users/mocks"
	"github.com/matryer/is"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		it    string
		model *users.User

		// Store
		userCreateInput  *users.User
		userCreateResult *users.User
		userCreateError  error

		expectedError  string
		expectedResult *users.User
	}{
		{
			it: "it create user",
			model: &users.User{
				Email: "mock-email",
			},
			userCreateInput: &users.User{
				ID:        GenUUID().String(),
				Email:     "mock-email",
				CreatedAt: GenTime(),
				UpdatedAt: GenTime(),
			},
			userCreateResult: &users.User{
				ID:        GenUUID().String(),
				Email:     "mock-email",
				CreatedAt: GenTime(),
				UpdatedAt: GenTime(),
			},
			expectedResult: &users.User{
				ID:        GenUUID().String(),
				Email:     "mock-email",
				CreatedAt: GenTime(),
				UpdatedAt: GenTime(),
			},
		},
		{
			it: "it return error on store Create",
			model: &users.User{
				Email: "mock-email",
			},
			userCreateInput: &users.User{
				ID:        GenUUID().String(),
				Email:     "mock-email",
				CreatedAt: GenTime(),
				UpdatedAt: GenTime(),
			},
			userCreateError: errors.Wrap("mock-error"),
			expectedError:   "create user: mock-error",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			store := &mocks.StoreMock{
				CreateFunc: func(ctx context.Context, model *users.User) (*users.User, error) {
					checkIs.Equal(model, tc.userCreateInput)

					return tc.userCreateResult, tc.userCreateError
				},
			}

			service := users.New(
				users.NewConfig(""),
				store,
				GenTime,
				GenUUID,
			)

			res, err := service.Create(context.Background(), tc.model)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(res, tc.expectedResult)
		})
	}
}
