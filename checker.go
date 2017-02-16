package healthz

import (
	"database/sql"
	"errors"
	"net/http"
	"sync"
)

// ErrCheckFailed is a generic error returned when a check fails
var ErrCheckFailed = errors.New("Health check failed")

// Checker is responsible for checking certain resources
type Checker interface {
	Type() string
	Check() error
}

// StatusChecker checks the status based on an internal state
type StatusChecker struct {
	status bool
	mu     *sync.Mutex
}

// NewStatusChecker creates a new status checker with an initial status
func NewStatusChecker(status bool) *StatusChecker {
	return &StatusChecker{
		status: status,
		mu:     &sync.Mutex{},
	}
}

// Type returns the name of the status checker
func (c *StatusChecker) Type() string {
	return "Status"
}

// Check checks the internal status and returns an error if it is false
func (c *StatusChecker) Check() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.status {
		return nil
	}

	return ErrCheckFailed
}

// SetStatus sets the internal status
func (c *StatusChecker) SetStatus(status bool) {
	c.mu.Lock()
	c.status = status
	c.mu.Unlock()
}

// DbChecker checks if a database is available through builtin database/sql
type DbChecker struct {
	db *sql.DB
}

// NewDbChecker creates a new DB checker with a connection
func NewDbChecker(db *sql.DB) *DbChecker {
	return &DbChecker{
		db: db,
	}
}

// Type returns the name of the database checker
func (c *DbChecker) Type() string {
	return "DatabasePing"
}

// Check checks the database status by pinging it
func (c *DbChecker) Check() error {
	return c.db.Ping()
}

// HTTPChecker checks if an HTTP service is available
type HTTPChecker struct {
	url string
}

// NewHTTPChecker creates a new HTTP checker with an URL
func NewHTTPChecker(url string) *HTTPChecker {
	return &HTTPChecker{
		url: url,
	}
}

// Type returns the name of the HTTP checker
func (c *HTTPChecker) Type() string {
	return "HTTPPing"
}

// Check checks the HTTP service status
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
