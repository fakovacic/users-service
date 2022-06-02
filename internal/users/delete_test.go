package users_test

import (
	"context"
	"testing"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/mocks"
	"github.com/matryer/is"
)

func TestDelete(t *testing.T) {
	cases := []struct {
		it string
		id string

		// Store
		userDeleteInput string
		userDeleteError error

		expectedError string
	}{
		{
			it:              "it delete user",
			id:              "mock-id",
			userDeleteInput: "mock-id",
		},
		{
			it:              "it return error on store Delete",
			id:              "mock-id",
			userDeleteInput: "mock-id",
			expectedError:   "delete user: mock-error",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			store := &mocks.StoreMock{
				DeleteFunc: func(ctx context.Context, id string) error {
					checkIs.Equal(id, tc.userDeleteInput)

					return tc.userDeleteError
				},
			}

			service := users.New(
				users.NewConfig(""),
				store,
				GenTime,
				GenUUID,
			)

			err := service.Delete(context.Background(), tc.id)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
		})
	}
}
