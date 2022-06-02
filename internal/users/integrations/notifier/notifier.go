package notifier

import (
	"github.com/fakovacic/users-service/internal/users"
)

func New(c *users.Config) users.Notifier {
	return &notifier{
		config: c,
	}
}

type notifier struct {
	config *users.Config
}
