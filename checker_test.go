package healthz_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"errors"

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

type PingerMock struct {
	err error
}

func (p *PingerMock) Ping() error {
	return p.err
}

func TestPingChecker_Check(t *testing.T) {
	checker := healthz.NewPingChecker(&PingerMock{})

	assertCheckerSuccessful(t, checker)
}

func TestPingChecker_Check_Fail(t *testing.T) {
	checker := healthz.NewPingChecker(&PingerMock{errors.New("ping failed")})

	err := checker.Check()

	if err == nil {
		t.Fatal("Expected error, none received")
	}
}

func TestHTTPChecker_Check(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer ts.Close()

	checker := healthz.NewHTTPChecker(ts.URL)

	assertCheckerSuccessful(t, checker)
}

func TestHTTPChecker_Check_Fail(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("error"))
	}))
	defer ts.Close()

	checker := healthz.NewHTTPChecker(ts.URL)

	assertCheckerFailed(t, checker)
}
