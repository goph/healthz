package healthz

import (
	"testing"
)

func TestCompositeChecker_Check(t *testing.T) {
	checker1 := new(AlwaysSuccessChecker)
	checker2 := new(AlwaysSuccessChecker)

	checkers := NewCompositeChecker(checker1, checker2)

	assertCheckerSuccessful(t, checkers)
}

func TestCompositeChecker_Check_Fail(t *testing.T) {
	checker1 := new(AlwaysSuccessChecker)
	checker2 := new(AlwaysFailureChecker)

	checkers := NewCompositeChecker(checker1, checker2)

	assertCheckerFailed(t, checkers)
}
