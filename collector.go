package healthz

import "net/http"

// Collector is a global context structure to accept checkers from all kinds of sources.
// It aggregates them and returns a HealthService.
type Collector map[string][]Checker

// RegisterChecker registers a new checker for a specific check.
func (c Collector) RegisterChecker(check string, checker Checker) {
	c[check] = append(c[check], checker)
}

// Handler returns an http.Handler for a check.
// If a check is not found the returned handler will always return success.
func (c Collector) Handler(check string) http.Handler {
	checkers, ok := c[check]
	if !ok {
		return NewHandler(&AlwaysSuccessChecker{})
	}

	if len(checkers) == 1 {
		return NewHandler(checkers[0])
	}

	return NewHandler(NewCompositeChecker(checkers...))
}
