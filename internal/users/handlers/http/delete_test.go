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

func TestDelete(t *testing.T) {
	cases := []struct {
		it string

		id string

		deleteInputID string
		deleteError   error

		expectedError  string
		expectedResult string
		expectedStatus int
	}{
		{
			it: "it delete user",
			id: "mock-id",

			deleteInputID:  "mock-id",
			expectedStatus: http.StatusOK,
			expectedResult: "200",
		},
		{
			it: "it return error on service Delete",
			id: "mock-id",

			deleteInputID: "mock-id",

			deleteError:    errors.Wrap("mock-error"),
			expectedResult: `{"message":"mock-error","status":500}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			service := &mocks.ServiceMock{
				DeleteFunc: func(contextMoqParam context.Context, id string) error {
					checkIs.Equal(id, tc.deleteInputID)

					return tc.deleteError
				},
			}

			req, err := http.NewRequest(
				http.MethodDelete,
				fmt.Sprintf("/%s", tc.id),
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

			router.DELETE("/:id", h.Delete)
			router.ServeHTTP(rr, req)

			checkIs.Equal(rr.Code, tc.expectedStatus)

			checkIs.Equal(rr.Body.String(), tc.expectedResult)

		})
	}
}
