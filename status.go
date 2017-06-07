package healthz

import "sync"

// Status is an enum type representing health status.
type Status int

// Possible values are Health and Unhealthy.
const (
	Healthy Status = iota
	Unhealthy
)

// StatusChecker checks the status based on an internal state.
type StatusChecker struct {
	status Status
	mu     *sync.RWMutex
}

// NewStatusChecker creates a new StatusChecker with an initial state.
func NewStatusChecker(status Status) *StatusChecker {
	return &StatusChecker{
		status: status,
		mu:     &sync.RWMutex{},
	}
}

// Check implements the Checker interface and checks the internal state.
// Returns an error if the value of state is false.
func (c *StatusChecker) Check() error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.status == Healthy {
		return nil
	}

	return ErrCheckFailed
}

// SetStatus sets the internal state of the checker.
func (c *StatusChecker) SetStatus(status Status) {
	c.mu.Lock()
	c.status = status
	c.mu.Unlock()
}
