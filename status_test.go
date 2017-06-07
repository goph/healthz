package healthz_test

import (
	"testing"

	"github.com/sagikazarmark/healthz"
)

func TestStatusChecker_Check(t *testing.T) {
	checker := healthz.NewStatusChecker(healthz.Healthy)

	assertCheckerSuccessful(t, checker)
}

func TestStatusChecker_Check_Fail(t *testing.T) {
	checker := healthz.NewStatusChecker(healthz.Unhealthy)

	assertCheckerFailed(t, checker)
}

func TestStatusChecker_SetStatus(t *testing.T) {
	checker := healthz.NewStatusChecker(healthz.Unhealthy)

	checker.SetStatus(healthz.Healthy)

	assertCheckerSuccessful(t, checker)
}
