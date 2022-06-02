package http_test

import (
	"bytes"
	"context"
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

func TestCreate(t *testing.T) {
	cases := []struct {
		it string

		requestBody string

		createInputUser *users.User
		createResponse  *users.User
		createError     error

		expectedError  string
		expectedResult string
		expectedStatus int
	}{
		{
			it:          "it create user",
			requestBody: `{"user":{"email":"mock-email"}}`,

			createInputUser: &users.User{
				Email: "mock-email",
			},

			createResponse: &users.User{
				ID:        "mock-id",
				FirstName: "mock-first-name",
				LastName:  "mock-last-name",
				Nickname:  "mock-nickname",
				Password:  "mock-password",
				Email:     "mock-email",
				Country:   "mock-country",
			},

			expectedResult: `{"user":{"id":"mock-id","first_name":"mock-first-name","last_name":"mock-last-name","nickname":"mock-nickname","password":"mock-password","email":"mock-email","country":"mock-country","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}}`,
			expectedStatus: http.StatusOK,
		},
		{
			it:          "it return error on service Create",
			requestBody: `{"user":{"email":"mock-email"}}`,

			createInputUser: &users.User{
				Email: "mock-email",
			},

			createError:    errors.Wrap("mock-error"),
			expectedResult: `{"message":"mock-error","status":500}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			service := &mocks.ServiceMock{
				CreateFunc: func(contextMoqParam context.Context, user *users.User) (*users.User, error) {
					checkIs.Equal(user, tc.createInputUser)

					return tc.createResponse, tc.createError
				},
			}

			req, err := http.NewRequest(
				http.MethodPost,
				"/",
				bytes.NewReader([]byte(tc.requestBody)),
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

			router.POST("/", h.Create)
			router.ServeHTTP(rr, req)

			checkIs.Equal(rr.Body.String(), tc.expectedResult)
			checkIs.Equal(rr.Code, tc.expectedStatus)
		})
	}
}
