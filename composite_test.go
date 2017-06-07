package healthz_test

import (
	"testing"

	"github.com/sagikazarmark/healthz"
)

func TestCompositeChecker_Check(t *testing.T) {
	checker1 := new(healthz.AlwaysSuccessChecker)
	checker2 := new(healthz.AlwaysSuccessChecker)

	checkers := healthz.NewCompositeChecker(checker1, checker2)

	assertCheckerSuccessful(t, checkers)
}

func TestCompositeChecker_Check_Fail(t *testing.T) {
	checker1 := new(healthz.AlwaysSuccessChecker)
	checker2 := new(healthz.AlwaysFailureChecker)

	checkers := healthz.NewCompositeChecker(checker1, checker2)

	assertCheckerFailed(t, checkers)
}
