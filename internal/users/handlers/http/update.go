package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/julienschmidt/httprouter"
)

type UserUpdateRequest struct {
	Fields []string    `json:"fields"`
	User   *users.User `json:"user"`
}

type UserUpdateResponse struct {
	User *users.User `json:"user"`
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request, par httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.writeError(w, r, ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})

		return
	}

	var req UserUpdateRequest

	err = json.Unmarshal(body, &req)
	if err != nil {
		h.writeError(w, r, ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})

		return
	}

	user, err := h.service.Update(r.Context(), par.ByName("id"), req.User, req.Fields)
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

	h.writeResponse(w, r, &UserUpdateResponse{
		User: user,
	})
}
