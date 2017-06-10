package healthz

import (
	"testing"
)

func assertCheckerSuccessful(t *testing.T, checker Checker) {
	err := checker.Check()

	if err != nil {
		t.Fatalf("Received unexpected error: %+v", err)
	}
}

func assertCheckerFailed(t *testing.T, checker Checker) {
	err := checker.Check()

	if err != ErrCheckFailed {
		t.Fatal("Expected ErrCheckFailed, none received")
	}
}

func TestCheckerFunc_Check(t *testing.T) {
	checker := CheckFunc(func() error {
		return nil
	})

	assertCheckerSuccessful(t, checker)
}

func TestCheckerFunc_Check_Fail(t *testing.T) {
	checker := CheckFunc(func() error {
		return ErrCheckFailed
	})

	assertCheckerFailed(t, checker)
}
