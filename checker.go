package healthz

import (
	"errors"
	"net"
	"net/http"
	"sync"
	"time"
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
	mu     *sync.Mutex
}

// NewStatusChecker creates a new StatusChecker with an initial state.
func NewStatusChecker(status Status) *StatusChecker {
	return &StatusChecker{
		status: status,
		mu:     &sync.Mutex{},
	}
}

// Check implements the Checker interface and checks the internal state.
// Returns an error if the value of state is false.
func (c *StatusChecker) Check() error {
	c.mu.Lock()
	defer c.mu.Unlock()

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

// Pinger is a commonly used interface to check if a connection is alive (used in sql.DB for example).
type Pinger interface {
	Ping() error
}

// PingChecker checks if a Pinger responds without an error.
type PingChecker struct {
	pinger Pinger
}

// NewPingChecker creates a new PingChecker with a Pinger.
func NewPingChecker(pinger Pinger) *PingChecker {
	return &PingChecker{pinger}
}

// Check implements the Checker interface and checks a resource status by pinging it.
func (c *PingChecker) Check() error {
	return c.pinger.Ping()
}

// HTTPChecker checks if an HTTP endpoint is available and successfully responds.
type HTTPChecker struct {
	url     string
	timeout time.Duration
	method  string
}

// HTTPCheckerOption configures how we check the HTTP endpoint.
type HTTPCheckerOption func(*HTTPChecker)

// WithHTTPTimeout returns an HTTPCheckerOption that specifies the timeout for HTTP requests.
//
// Setting a timeout is highly recommended, but it needs to be carefully chosen to avoid false results.
func WithHTTPTimeout(timeout time.Duration) HTTPCheckerOption {
	return func(c *HTTPChecker) {
		c.timeout = timeout
	}
}

// WithHTTPMethod returns an HTTPCheckerOption that specifies the method for HTTP requests.
//
// The default method is "GET" which works in most of the cases, but another popular choice is "HEAD".
func WithHTTPMethod(method string) HTTPCheckerOption {
	return func(c *HTTPChecker) {
		c.method = method
	}
}

// NewHTTPChecker creates a new HTTPChecker with a URL and optional configuration.
//
// Example:
// 		checker := healthz.NewHTTPChecker("http://example.com", healthz.WithHTTPTimeout(3*time.Second))
func NewHTTPChecker(url string, opts ...HTTPCheckerOption) *HTTPChecker {
	checker := &HTTPChecker{
		url:    url,
		method: http.MethodGet,
	}

	for _, opt := range opts {
		opt(checker)
	}

	return checker
}

// Check implements the Checker interface and checks the HTTP endpoint status.
func (c *HTTPChecker) Check() error {
	client := &http.Client{
		Timeout: c.timeout,
	}

	req, err := http.NewRequest(c.method, c.url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return ErrCheckFailed
	}

	return nil
}

// TCPChecker checks if something is listening on a TCP address.
type TCPChecker struct {
	addr    string
	timeout time.Duration
}

// TCPCheckerOption configures how we check the TCP address.
type TCPCheckerOption func(*TCPChecker)

// WithTCPTimeout returns a TCPCheckerOption that specifies the timeout for TCP requests.
//
// Setting a timeout is highly recommended, but it needs to be carefully chosen to avoid false results.
func WithTCPTimeout(timeout time.Duration) TCPCheckerOption {
	return func(c *TCPChecker) {
		c.timeout = timeout
	}
}

// NewTCPChecker creates a new TCPChecker with an address and optional configuration.
// Example:
// 		checker := healthz.NewTCPChecker(":80", healthz.WithTCPTimeout(3*time.Second))
func NewTCPChecker(addr string, opts ...TCPCheckerOption) *TCPChecker {
	checker := &TCPChecker{
		addr: addr,
	}

	for _, opt := range opts {
		opt(checker)
	}

	return checker
}

// Check implements the Checker interface and checks the TCP address status.
func (c *TCPChecker) Check() error {
	dialer := net.Dialer{
		Timeout: c.timeout,
	}

	conn, err := dialer.Dial("tcp", c.addr)

	if err != nil {
		return err
	}

	conn.Close()

	return nil
}
