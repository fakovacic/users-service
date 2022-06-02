package config

import (
	"os"

	"github.com/fakovacic/users-service/internal/users"
)

func NewConfig() (*users.Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	return users.NewConfig(env), nil
}
