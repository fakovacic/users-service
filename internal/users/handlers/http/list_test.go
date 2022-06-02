package http_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/errors"
	handlers "github.com/fakovacic/users-service/internal/users/handlers/http"
	"github.com/fakovacic/users-service/internal/users/mocks"
	"github.com/julienschmidt/httprouter"
	"github.com/matryer/is"
)

func TestList(t *testing.T) {
	cases := []struct {
		it string

		queryParams string

		listInputMeta     *users.Meta
		listResponseMeta  *users.Meta
		listResponseUsers []*users.User
		listError         error

		expectedError  string
		expectedResult string
		expectedStatus int
	}{
		{
			it:          "it return users list",
			queryParams: `?page=2&limit=200`,

			listInputMeta: &users.Meta{
				Page:  2,
				Limit: 200,
			},
			listResponseMeta: &users.Meta{
				Page:  2,
				Limit: 200,
				Pages: 1,
				Count: 1,
			},
			listResponseUsers: []*users.User{
				{
					ID:        "mock-id",
					FirstName: "mock-first-name",
					LastName:  "mock-last-name",
					Nickname:  "mock-nickname",
					Password:  "mock-password",
					Email:     "mock-email",
					Country:   "mock-country",
				},
			},

			expectedResult: `{"meta":{"page":2,"pages":1,"count":1,"limit":200},"users":[{"id":"mock-id","first_name":"mock-first-name","last_name":"mock-last-name","nickname":"mock-nickname","password":"mock-password","email":"mock-email","country":"mock-country","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]}`,
			expectedStatus: http.StatusOK,
		},
		{
			it:          "it return users list with filters",
			queryParams: `?page=2&limit=200&filters.country=XX`,

			listInputMeta: &users.Meta{
				Page:  2,
				Limit: 200,
				Filters: []users.Filters{
					{
						Field: "country",
						Values: []string{
							"XX",
						},
					},
				},
			},
			listResponseMeta: &users.Meta{
				Page:  2,
				Limit: 200,
				Pages: 1,
				Count: 1,
			},
			listResponseUsers: []*users.User{
				{
					ID:        "mock-id",
					FirstName: "mock-first-name",
					LastName:  "mock-last-name",
					Nickname:  "mock-nickname",
					Password:  "mock-password",
					Email:     "mock-email",
					Country:   "mock-country",
				},
			},

			expectedResult: `{"meta":{"page":2,"pages":1,"count":1,"limit":200},"users":[{"id":"mock-id","first_name":"mock-first-name","last_name":"mock-last-name","nickname":"mock-nickname","password":"mock-password","email":"mock-email","country":"mock-country","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]}`,
			expectedStatus: http.StatusOK,
		},
		{
			it: "it return error on service List",

			listInputMeta: &users.Meta{
				Page:  1,
				Limit: 10,
			},
			listError:      errors.Wrap("mock-error"),
			expectedResult: `{"message":"mock-error","status":500}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			service := &mocks.ServiceMock{
				ListFunc: func(contextMoqParam context.Context, meta *users.Meta) (*users.Meta, []*users.User, error) {
					checkIs.Equal(meta, tc.listInputMeta)

					return tc.listResponseMeta, tc.listResponseUsers, tc.listError
				},
			}

			req, err := http.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/%s", tc.queryParams),
				nil,
			)

			req = req.WithContext(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			h := handlers.New(
				users.NewConfig(""),
				service,
			)

			router := httprouter.New()
			rr := httptest.NewRecorder()

			router.GET("/", h.List)
			router.ServeHTTP(rr, req)

			checkIs.Equal(rr.Body.String(), tc.expectedResult)
			checkIs.Equal(rr.Code, tc.expectedStatus)
		})
	}
}
