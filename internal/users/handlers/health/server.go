package health

import (
	"net/http"
	"sync"

	"github.com/braintree/manners"
)

var (
	healthzStatus = http.StatusOK
	mu            sync.RWMutex
)

// StartServer starts health server
func StartServer(healthAddr string) *manners.GracefulServer {
	hmux := http.NewServeMux()
	hmux.HandleFunc("/health", Health)

	healthServer := manners.NewServer()
	healthServer.Addr = healthAddr
	healthServer.Handler = hmux

	return healthServer
}
