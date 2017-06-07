package healthz

import "sync"

// CompositeChecker is responsible for executing a series of checks and decide if the resource is up or not.
type CompositeChecker struct {
	checkers []Checker
	mu       *sync.Mutex
}

// NewCompositeChecker is a shortcut for easily creating a new Checker collection.
func NewCompositeChecker(checkers ...Checker) *CompositeChecker {
	return &CompositeChecker{
		checkers: checkers,
		mu:       new(sync.Mutex),
	}
}

// Check implements the Checker interface and executes the underlying checks.
//
// Note that since we have no information about what may become a Checker, this cannot be called concurrently.
func (c *CompositeChecker) Check() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, checker := range c.checkers {
		if err := checker.Check(); err != nil {
			return err
		}
	}

	return nil
}
