package healthz

import "sync"

// Probe is responsible for executing a series of checks and decide if the service is up or not
type Probe struct {
	checkers []HealthChecker
	mu       *sync.Mutex
}

// NewProbe is a shortcut for easily creating a new Probe
func NewProbe(checkers ...HealthChecker) *Probe {
	return &Probe{
		checkers: checkers,
		mu:       &sync.Mutex{},
	}
}

// Check executes the underlying health checks
func (p *Probe) Check() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, c := range p.checkers {
		if err := c.Ping(); err != nil {
			return err
		}
	}

	return nil
}
