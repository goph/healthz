package healthz

import "net/http"

// HealthService is the HTTP handler implementation which wraps the Probe and Checker logic
type HealthService struct {
	LivenessProbe  *Probe
	ReadinessProbe *Probe
}

// NewHealthService creates a new HealthService from user configured Probes
func NewHealthService(livenessProbe *Probe, readinessProbe *Probe) *HealthService {
	return &HealthService{
		LivenessProbe:  livenessProbe,
		ReadinessProbe: readinessProbe,
	}
}

// RegisterHandlers registers HTTP Endpoints in a Mux
func (s *HealthService) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", s.HealthStatus)
	mux.HandleFunc("/readiness", s.ReadinessStatus)
}

// HealthStatus checks if the application is healthy
func (s *HealthService) HealthStatus(w http.ResponseWriter, r *http.Request) {
	s.checkStatus(w, r, s.LivenessProbe)
}

// ReadinessStatus checks if the app is ready for accepting request (eg. database is available as well)
func (s *HealthService) ReadinessStatus(w http.ResponseWriter, r *http.Request) {
	s.checkStatus(w, r, s.ReadinessProbe)
}

func (s *HealthService) checkStatus(w http.ResponseWriter, r *http.Request, p *Probe) {
	if err := p.Check(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("error"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
