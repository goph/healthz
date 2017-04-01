package healthz

// HealthService is a helper struct which wraps two Checkers and exposes them as HTTP handlers.
type HealthService struct {
	livenessChecker  Checker
	readinessChecker Checker
}

// NewHealthService creates a new HealthService from user configured Checkers.
func NewHealthService(livenessChecker Checker, readinessChecker Checker) *HealthService {
	return &HealthService{
		livenessChecker:  livenessChecker,
		readinessChecker: readinessChecker,
	}
}
