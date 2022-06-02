package http_test

import (
	"bytes"
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

func TestUpdate(t *testing.T) {
	cases := []struct {
		it string

		id          string
		requestBody string

		updateInputID     string
		updateInputUser   *users.User
		updateInputFields []string
		updateResponse    *users.User
		updateError       error

		expectedError  string
		expectedResult string
		expectedStatus int
	}{
		{
			it:          "it update user",
			id:          "mock-id",
			requestBody: `{"fields":["email"],"user":{"email":"mock-email"}}`,

			updateInputID: "mock-id",
			updateInputUser: &users.User{
				Email: "mock-email",
			},
			updateInputFields: []string{
				"email",
			},

			updateResponse: &users.User{
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
			it:          "it return error on service Update",
			id:          "mock-id",
			requestBody: `{"fields":["email"],"user":{"email":"mock-email"}}`,

			updateInputID: "mock-id",
			updateInputUser: &users.User{
				Email: "mock-email",
			},
			updateInputFields: []string{
				"email",
			},

			updateError:    errors.Wrap("mock-error"),
			expectedResult: `{"message":"mock-error","status":500}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			service := &mocks.ServiceMock{
				UpdateFunc: func(contextMoqParam context.Context, id string, user *users.User, fields []string) (*users.User, error) {
					checkIs.Equal(id, tc.updateInputID)
					checkIs.Equal(user, tc.updateInputUser)
					checkIs.Equal(fields, tc.updateInputFields)

					return tc.updateResponse, tc.updateError
				},
			}

			req, err := http.NewRequest(
				http.MethodPatch,
				fmt.Sprintf("/%s", tc.id),
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

			router.PATCH("/:id", h.Update)
			router.ServeHTTP(rr, req)

			checkIs.Equal(rr.Body.String(), tc.expectedResult)
			checkIs.Equal(rr.Code, tc.expectedStatus)
		})
	}
}
