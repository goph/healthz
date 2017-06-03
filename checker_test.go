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

func TestCheckers_Check(t *testing.T) {
	checker1 := new(healthz.AlwaysSuccessChecker)
	checker2 := new(healthz.AlwaysSuccessChecker)

	checkers := healthz.NewCheckers(checker1, checker2)

	assertCheckerSuccessful(t, checkers)
}

func TestCheckers_Check_Fail(t *testing.T) {
	checker1 := new(healthz.AlwaysSuccessChecker)
	checker2 := new(healthz.AlwaysFailureChecker)

	checkers := healthz.NewCheckers(checker1, checker2)

	assertCheckerFailed(t, checkers)
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
