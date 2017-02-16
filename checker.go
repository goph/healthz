package healthz

import (
	"database/sql"
	"errors"
	"net/http"
	"sync"
)

// ErrCheckFailed is a generic error returned when a check fails
var ErrCheckFailed = errors.New("Check failed")

// Checker is the interface for checking different resources
type Checker interface {
	// Returns nil if the check passes
	Check() error
}

// Checkers is a Checker collection responsible for executing a series of checks and decide if the resource is up or not
type Checkers struct {
	checkers []Checker
	mu       *sync.Mutex
}

// NewCheckers is a shortcut for easily creating a new Checker collection
func NewCheckers(checkers ...Checker) *Checkers {
	return &Checkers{
		checkers: checkers,
		mu:       new(sync.Mutex),
	}
}

// Check implements the Checker interface and executes the underlying checks
// Note that since we have no information about what may become a Checker, this cannot be called concurrently
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

// Status is an enum type representing health status
type Status int

// Possibly values are Health and Unhealthy
const (
	Healthy Status = iota
	Unhealthy
)

// StatusChecker checks the status based on an internal state
type StatusChecker struct {
	status Status
	mu     *sync.Mutex
}

// NewStatusChecker creates a new StatusChecker with an initial state
func NewStatusChecker(status Status) *StatusChecker {
	return &StatusChecker{
		status: status,
		mu:     &sync.Mutex{},
	}
}

// Check implements the Checker interface and checks the internal state
// Returns an error if the state is false
func (c *StatusChecker) Check() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.status == Healthy {
		return nil
	}

	return ErrCheckFailed
}

// SetStatus sets the internal state
func (c *StatusChecker) SetStatus(status Status) {
	c.mu.Lock()
	c.status = status
	c.mu.Unlock()
}

// DbChecker checks if a database is available through builtin database/sql
type DbChecker struct {
	db *sql.DB
}

// NewDbChecker creates a new DBChecker with a DB connection
func NewDbChecker(db *sql.DB) *DbChecker {
	return &DbChecker{
		db: db,
	}
}

// Check implements the Checker interface and checks the database status by pinging it
func (c *DbChecker) Check() error {
	return c.db.Ping()
}

// HTTPChecker checks if an HTTP endpoint is available
type HTTPChecker struct {
	url string
}

// NewHTTPChecker creates a new HTTPChecker with a URL
func NewHTTPChecker(url string) *HTTPChecker {
	return &HTTPChecker{
		url: url,
	}
}

// Check implements the Checker interface and checks the HTTP endpoint status
func (c *HTTPChecker) Check() error {
	resp, err := http.Get(c.url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return ErrCheckFailed
	}

	return nil
}
