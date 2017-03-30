package healthz

import "net/http"

// NewHealthServiceHandler creates an http.Handler from user configured Checkers
// The handler is a standard http.ServeMux
func NewHealthServiceHandler(livenessChecker Checker, readinessChecker Checker) http.Handler {
	healthService := NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", healthService.HealthStatus)
	mux.HandleFunc("/readiness", healthService.ReadinessStatus)

	return mux
}

// HealthStatus checks if the application is healthy
func (s *HealthService) HealthStatus(w http.ResponseWriter, r *http.Request) {
	s.checkStatus(w, r, s.livenessChecker)
}

// ReadinessStatus checks if the app is ready for accepting request (eg. database is available as well)
func (s *HealthService) ReadinessStatus(w http.ResponseWriter, r *http.Request) {
	s.checkStatus(w, r, s.readinessChecker)
}

// Since both health check rely on Checkers, common logic for them is here
func (s *HealthService) checkStatus(w http.ResponseWriter, r *http.Request, c Checker) {
	// If the checker fails, we return an error
	if err := c.Check(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("error"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
