package healthz

// AlwaysSuccessChecker always returns success as the check result
type AlwaysSuccessChecker struct{}

// Check implements the Checker interface
func (c *AlwaysSuccessChecker) Check() error {
	return nil
}

// AlwaysFailureChecker always returns failure as the check result
type AlwaysFailureChecker struct{}

// Check implements the Checker interface
func (c *AlwaysFailureChecker) Check() error {
	return ErrCheckFailed
}
