package healthz

import (
	"net/http"
	"time"
)

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
