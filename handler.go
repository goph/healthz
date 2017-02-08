package healthz

import "net/http"

// HealthService is the HTTP handler implementation which wraps the Probe and Checker logic
type HealthService struct {
	LivenessProbe  *Probe
	ReadinessProbe *Probe
}

// NewHealthService creates a new HealthService from user configured Probes
func NewHealthService(livenessProbe *Probe, readinessProbe *Probe) http.Handler {
	healthService := &HealthService{
		LivenessProbe:  livenessProbe,
		ReadinessProbe: readinessProbe,
	}

	hmux := http.NewServeMux()

	hmux.HandleFunc("/healthz", healthService.HealthStatus)
	hmux.HandleFunc("/readiness", healthService.ReadinessStatus)

	return hmux
}

// HealthStatus checks if the application is healthy
func (hs *HealthService) HealthStatus(w http.ResponseWriter, r *http.Request) {
	hs.checkStatus(w, r, hs.LivenessProbe)
}

// ReadinessStatus checks if the app is ready for accepting request (eg. database is available as well)
func (hs *HealthService) ReadinessStatus(w http.ResponseWriter, r *http.Request) {
	hs.checkStatus(w, r, hs.ReadinessProbe)
}

func (hs *HealthService) checkStatus(w http.ResponseWriter, r *http.Request, p *Probe) {
	if err := p.Check(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("error"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
