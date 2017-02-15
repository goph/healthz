// Package healthz provides tools for service health checks
package healthz

// HealthService is the mplementation which wraps the Probe and Checker logic
type HealthService struct {
	livenessProbe  *Probe
	readinessProbe *Probe
}

// NewHealthService creates a new HealthService from user configured Probes
func NewHealthService(livenessProbe *Probe, readinessProbe *Probe) *HealthService {
	return &HealthService{
		livenessProbe:  livenessProbe,
		readinessProbe: readinessProbe,
	}
}
