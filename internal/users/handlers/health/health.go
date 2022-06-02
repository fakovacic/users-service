package health

import "net/http"

// Health responds to health check requests.
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(GetHealthStatus())
}

func GetHealthStatus() int {
	mu.RLock()
	defer mu.RUnlock()

	return healthzStatus
}

func SetHealthStatus(status int) {
	mu.Lock()
	healthzStatus = status
	mu.Unlock()
}
