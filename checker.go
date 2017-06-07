package healthz

import (
	"errors"
	"sync"
)

// ErrCheckFailed is a generic error which MAY be returned when a check fails.
var ErrCheckFailed = errors.New("Check failed")

// Checker is the interface for checking different resources.
type Checker interface {
	// Check returns nil if the check passes.
	Check() error
}

// Checkers is a Checker collection responsible for executing a series of checks and decide if the resource is up or not.
type Checkers struct {
	checkers []Checker
	mu       *sync.Mutex
}

// NewCheckers is a shortcut for easily creating a new Checker collection.
func NewCheckers(checkers ...Checker) *Checkers {
	return &Checkers{
		checkers: checkers,
		mu:       new(sync.Mutex),
	}
}

// Check implements the Checker interface and executes the underlying checks.
//
// Note that since we have no information about what may become a Checker, this cannot be called concurrently.
func (c *Checkers) Check() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, checker := range c.checkers {
		if err := checker.Check(); err != nil {
			return err
		}
	}

	return nil
}

// CheckFunc is a convenience type to create functions that implement the Checker interface.
type CheckFunc func() error

// Check implements the Checker interface and allows any func() error signatured method to be passed as a Checker.
func (f CheckFunc) Check() error {
	return f()
}
