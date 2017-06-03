package healthz

import (
	"net"
	"time"
)

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
