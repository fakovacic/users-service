package users_test

import (
	"context"
	"testing"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/fakovacic/users-service/internal/users/mocks"
	"github.com/matryer/is"
)

func TestList(t *testing.T) {
	cases := []struct {
		it   string
		meta *users.Meta

		// Store
		userListInput        *users.Meta
		userListResultCount  int64
		userListResultModels []*users.User
		userListError        error

		expectedError      string
		expectedResultMeta *users.Meta
		expectedResultList []*users.User
	}{
		{
			it: "it return list of users",
			meta: &users.Meta{
				Page:  1,
				Limit: 10,
			},
			userListInput: &users.Meta{
				Page:  1,
				Limit: 10,
			},
			userListResultCount: 1,
			userListResultModels: []*users.User{
				{
					ID: "mock-id",
				},
			},
			expectedResultMeta: &users.Meta{
				Page:  1,
				Limit: 10,
				Count: 1,
				Pages: 1,
			},
			expectedResultList: []*users.User{
				{
					ID: "mock-id",
				},
			},
		},
		{
			it: "it return error on store List",
			meta: &users.Meta{
				Page:  1,
				Limit: 10,
			},
			userListInput: &users.Meta{
				Page:  1,
				Limit: 10,
			},
			userListError: errors.Wrap("mock-error"),
			expectedError: "get users list: mock-error",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			store := &mocks.StoreMock{
				ListFunc: func(ctx context.Context, meta *users.Meta) (int64, []*users.User, error) {
					checkIs.Equal(meta, tc.userListInput)

					return tc.userListResultCount, tc.userListResultModels, tc.userListError
				},
			}

			service := users.New(
				users.NewConfig(""),
				store,
				nil,
				nil,
			)

			meta, res, err := service.List(context.Background(), tc.meta)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}

			checkIs.Equal(meta, tc.expectedResultMeta)
			checkIs.Equal(res, tc.expectedResultList)
		})
	}
}
