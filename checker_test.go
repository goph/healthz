package healthz_test

import (
	"testing"

	"github.com/sagikazarmark/healthz"
)

func assertCheckerSuccessful(t *testing.T, checker healthz.Checker) {
	err := checker.Check()

	if err != nil {
		t.Fatalf("Received unexpected error: %+v", err)
	}
}

func assertCheckerFailed(t *testing.T, checker healthz.Checker) {
	err := checker.Check()

	if err != healthz.ErrCheckFailed {
		t.Fatal("Expected ErrCheckFailed, none received")
	}
}

func TestCheckerFunc_Check(t *testing.T) {
	checker := healthz.CheckFunc(func() error {
		return nil
	})

	assertCheckerSuccessful(t, checker)
}

func TestCheckerFunc_Check_Fail(t *testing.T) {
	checker := healthz.CheckFunc(func() error {
		return healthz.ErrCheckFailed
	})

	assertCheckerFailed(t, checker)
}
