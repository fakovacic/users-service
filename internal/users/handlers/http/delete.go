package http

import (
	"net/http"

	"github.com/fakovacic/users-service/internal/users/errors"
	"github.com/julienschmidt/httprouter"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, par httprouter.Params) {
	err := h.service.Delete(r.Context(), par.ByName("id"))
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

	h.writeResponse(w, r, http.StatusOK)
}
