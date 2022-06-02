package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/julienschmidt/httprouter"
)

type UserCreateRequest struct {
	User *users.User `json:"user"`
}

type UserCreateResponse struct {
	User *users.User `json:"user"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request, par httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.writeError(w, r, ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})

		return
	}

	var req UserCreateRequest

	err = json.Unmarshal(body, &req)
	if err != nil {
		h.writeError(w, r, ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})

		return
	}

	user, err := h.service.Create(r.Context(), req.User)
	if err != nil {
		e, ok := err.(errors.Error)
		if ok {
			h.writeError(w, r, ErrorResponse{
				Message: e.Error(),
				Status:  e.Status,
			})

			return
		}

		h.writeError(w, r, ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})

		return
	}

	h.writeResponse(w, r, &UserCreateResponse{
		User: user,
	})
}
