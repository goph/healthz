package healthz

import "errors"

// ErrHealthCheckFailed is a generic error returned when a check fails
var ErrHealthCheckFailed = errors.New("Health check failed")

// HealthChecker is responsible for checking certain resources
type HealthChecker interface {
	Type() string
	Ping() error
}

// StatusHealthChecker checks the status based on an internal
type StatusHealthChecker struct {
	Status bool
}

// Type returns the name of the status checker
func (s *StatusHealthChecker) Type() string {
	return "Status"
}

// Ping checks the internal status and returns an error if it is false
func (s *StatusHealthChecker) Ping() error {
	if s.Status {
		return nil
	}

	return ErrHealthCheckFailed
}
