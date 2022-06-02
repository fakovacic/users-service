package middleware

import (
	"net/http"

	"github.com/fakovacic/users-service/internal/users"
)

func Logger(c *users.Config, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := c.Log.With().
			Str("url", r.URL.String()).
			Str("method", r.Method).
			Logger()
		l.Info().Msg("http request")
		h.ServeHTTP(w, r)
	})
}
