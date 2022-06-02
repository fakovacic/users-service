package users

import (
	"os"

	"github.com/rs/zerolog"
)

type ContextKey string

const (
	ServiceName string = "users"

	ContextKeyRequestID ContextKey = "requestID"
)

func NewConfig(env string) *Config {
	return &Config{
		Env: env,
		Log: newLogger(),
	}
}

type Config struct {
	Log zerolog.Logger
	Env string
}

func newLogger() zerolog.Logger {
	zerolog.TimestampFieldName = "timestamp"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	return zerolog.New(os.Stdout).
		With().Timestamp().Logger()
}
