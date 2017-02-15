package healthz

import "net/http"

// NewHealthServiceHandler creates an http.Handler from user configured Probes
// The handler is a standard http.ServeMux
func NewHealthServiceHandler(livenessProbe *Probe, readinessProbe *Probe) http.Handler {
	healthService := NewHealthService(livenessProbe, readinessProbe)
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", healthService.HealthStatus)
	mux.HandleFunc("/readiness", healthService.ReadinessStatus)

	return mux
}

// HealthStatus checks if the application is healthy
func (s *HealthService) HealthStatus(w http.ResponseWriter, r *http.Request) {
	s.checkStatus(w, r, s.livenessProbe)
}

// ReadinessStatus checks if the app is ready for accepting request (eg. database is available as well)
func (s *HealthService) ReadinessStatus(w http.ResponseWriter, r *http.Request) {
	s.checkStatus(w, r, s.readinessProbe)
}

// Since both health check relies on probes, common logic for them is here
func (s *HealthService) checkStatus(w http.ResponseWriter, r *http.Request, p *Probe) {
	// If the probe fails, we return an error
	if err := p.Check(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("error"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
