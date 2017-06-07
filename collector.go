package healthz

// Collector is a global context structure to accept checkers from all kinds of sources.
// It aggregates them and returns a HealthService.
type Collector map[string][]Checker

// RegisterChecker registers a new checker for a specific check.
func (c Collector) RegisterChecker(check string, checker Checker) {
	c[check] = append(c[check], checker)
}

// NewHealthService aggregates the checkers and returns a new HealthService.
func (c Collector) NewHealthService() HealthService {
	healthService := make(HealthService)

	for t, checkers := range c {
		if len(checkers) == 1 {
			healthService[t] = checkers[0]
		} else {
			healthService[t] = NewCompositeChecker(checkers...)
		}
	}

	return healthService
}
