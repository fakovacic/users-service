package postgres_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/fakovacic/users-service/internal/users/store/postgres"
	"github.com/matryer/is"
)

func TestGet(t *testing.T) {
	type row struct {
		id         string
		first_name string
		last_name  string
		nickname   string
		password   string
		email      string
		country    string
		created_at time.Time
		updated_at time.Time
	}

	cases := []struct {
		it             string
		id             string
		r              *row
		sqlError       error
		rowError       error
		expectedError  string
		expectedResult *users.User
	}{
		{
			it: "it returns user",
			id: "mock-id",
			r: &row{
				id:         "mock-id",
				first_name: "mock-first-name",
				last_name:  "mock-last-name",
				nickname:   "mock-nickname",
				password:   "mock-password",
				email:      "mock-email",
				country:    "mock-country",
			},
			expectedResult: &users.User{
				ID:        "mock-id",
				FirstName: "mock-first-name",
				LastName:  "mock-last-name",
				Nickname:  "mock-nickname",
				Password:  "mock-password",
				Email:     "mock-email",
				Country:   "mock-country",
			},
		},
		{
			it:            "it returns error on row error",
			rowError:      errors.Wrap("expected 7 destination arguments in Scan, not 9"),
			expectedError: "get user: sql: expected 7 destination arguments in Scan, not 9",
		},
		{
			it:            "it returns error db error",
			sqlError:      errors.Wrap("mock-error"),
			expectedError: "get user: mock-error",
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

			query := mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, first_name, last_name, nickname, password, email, country, created_at, updated_at FROM users WHERE id=$1`))

			switch {
			case tc.r != nil:
				mockRow := sqlmock.NewRows([]string{
					"id",
					"first_name",
					"last_name",
					"nickname",
					"password",
					"email",
					"country",
					"created_at",
					"updated_at",
				}).AddRow(
					tc.r.id,
					tc.r.first_name,
					tc.r.last_name,
					tc.r.nickname,
					tc.r.password,
					tc.r.email,
					tc.r.country,
					tc.r.created_at,
					tc.r.updated_at,
				)
				query.WillReturnRows(mockRow)
			case tc.rowError != nil:
				mockRow := sqlmock.NewRows([]string{
					"id",
					"first_name",
					"last_name",
					"nickname",
					"password",
					"email",
					"country",
				}).AddRow("", "", "", "", "", "", "").RowError(1, tc.rowError)
				query.WillReturnRows(mockRow)
			default:
				query.WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"first_name",
					"last_name",
					"nickname",
					"password",
					"email",
					"country",
					"created_at",
					"updated_at",
				}).AddRow("", "", "", "", "", "", "", time.Now(), time.Now()))
				query.WillReturnError(tc.sqlError)
			}

			service := postgres.New(db)

			res, err := service.Get(context.Background(), tc.id)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(res, tc.expectedResult)
		})
	}
}
