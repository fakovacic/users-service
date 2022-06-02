package config

import (
	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/handlers/http"
)

func NewHandlers(c *users.Config, service users.Service) http.Handler {
	return *http.New(c, service)
}
