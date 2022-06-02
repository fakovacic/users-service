package http

import (
	"net/http"

	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/julienschmidt/httprouter"
)

type UsersListResponse struct {
	Meta  *users.Meta   `json:"meta"`
	Users []*users.User `json:"users"`
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	meta := users.NewMeta()

	err := meta.ParseRequest(r)
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

	meta, users, err := h.service.List(r.Context(), meta)
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

	h.writeResponse(w, r, &UsersListResponse{
		Meta:  meta,
		Users: users,
	})
}
