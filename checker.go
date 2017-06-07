package healthz

import (
	"errors"
)

// ErrCheckFailed is a generic error which MAY be returned when a check fails.
var ErrCheckFailed = errors.New("Check failed")

// Checker is the interface for checking different resources.
type Checker interface {
	// Check returns nil if the check passes.
	Check() error
}

// CheckFunc is a convenience type to create functions that implement the Checker interface.
type CheckFunc func() error

// Check implements the Checker interface and allows any func() error signatured method to be passed as a Checker.
func (f CheckFunc) Check() error {
	return f()
}
