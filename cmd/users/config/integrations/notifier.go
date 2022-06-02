package integrations

import (
	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/integrations/notifier"
)

func NewNotifier(cfg *users.Config) users.Notifier {
	return notifier.New(cfg)
}
