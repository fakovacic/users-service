package postgres_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/fakovacic/users-service/internal/users/store/postgres"
	"github.com/matryer/is"
)

func TestDelete(t *testing.T) {
	cases := []struct {
		it            string
		id            string
		sqlError      error
		expectedError string
	}{
		{
			it: "it delete user",
			id: "mock-uuid",
		},
		{
			it:            "it error on sql error",
			id:            "mock-uuid",
			sqlError:      errors.Wrap("mock-error"),
			expectedError: "delete user: mock-error",
		},
	}
	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM users WHERE id = $1`)).
				WithArgs(tc.id).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tc.sqlError)

			service := postgres.New(db)

			err = service.Delete(context.Background(), tc.id)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
		})
	}
}
