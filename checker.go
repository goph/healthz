package healthz

import (
	"database/sql"
	"errors"
	"net/http"
	"sync"
)

// ErrHealthCheckFailed is a generic error returned when a check fails
var ErrHealthCheckFailed = errors.New("Health check failed")

// HealthChecker is responsible for checking certain resources
type HealthChecker interface {
	Type() string
	Ping() error
}

// StatusHealthChecker checks the status based on an internal state
type StatusHealthChecker struct {
	status bool
	mu     *sync.Mutex
}

// NewStatusHealthChecker creates a new status health checker with an initial status
func NewStatusHealthChecker(status bool) *StatusHealthChecker {
	return &StatusHealthChecker{
		status: status,
		mu:     &sync.Mutex{},
	}
}

// Type returns the name of the status checker
func (c *StatusHealthChecker) Type() string {
	return "Status"
}

// Ping checks the internal status and returns an error if it is false
func (c *StatusHealthChecker) Ping() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.status {
		return nil
	}

	return ErrHealthCheckFailed
}

// SetStatus sets the internal status
func (c *StatusHealthChecker) SetStatus(status bool) {
	c.mu.Lock()
	c.status = status
	c.mu.Unlock()
}

// DbHealthChecker checks if a database is available through builtin database/sql
type DbHealthChecker struct {
	db *sql.DB
}

// NewDbHealthChecker creates a new DB health checker with a connection
func NewDbHealthChecker(db *sql.DB) *DbHealthChecker {
	return &DbHealthChecker{
		db: db,
	}
}

// Type returns the name of the database checker
func (c *DbHealthChecker) Type() string {
	return "DatabasePing"
}

// Ping checks the database status by pinging it
func (c *DbHealthChecker) Ping() error {
	return c.db.Ping()
}

// HTTPHealthChecker checks if an HTTP service is available
type HTTPHealthChecker struct {
	url string
}

// NewHTTPHealthChecker creates a new HTTP health checker with an URL
func NewHTTPHealthChecker(url string) *HTTPHealthChecker {
	return &HTTPHealthChecker{
		url: url,
	}
}

// Type returns the name of the database checker
func (c *HTTPHealthChecker) Type() string {
	return "HTTPPing"
}

// Ping checks the database status by pinging it
func (c *HTTPHealthChecker) Ping() error {
	resp, err := http.Get(c.url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return ErrHealthCheckFailed
	}

	return nil
}
