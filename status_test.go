package healthz

import (
	"testing"
)

func TestStatusChecker_Check(t *testing.T) {
	checker := NewStatusChecker(Healthy)

	assertCheckerSuccessful(t, checker)
}

func TestStatusChecker_Check_Fail(t *testing.T) {
	checker := NewStatusChecker(Unhealthy)

	assertCheckerFailed(t, checker)
}

func TestStatusChecker_SetStatus(t *testing.T) {
	checker := NewStatusChecker(Unhealthy)

	checker.SetStatus(Healthy)

	assertCheckerSuccessful(t, checker)
}
