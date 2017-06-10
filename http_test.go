package healthz

import (
	"testing"

	"net/http"
	"net/http/httptest"
	"time"
)

func TestHTTPChecker_Check(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer ts.Close()

	checker := NewHTTPChecker(ts.URL)

	assertCheckerSuccessful(t, checker)
}

func TestHTTPChecker_Check_Fail(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("error"))
	}))
	defer ts.Close()

	checker := NewHTTPChecker(ts.URL)

	assertCheckerFailed(t, checker)
}

func TestHTTPChecker_Check_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Nanosecond)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer ts.Close()

	checker := NewHTTPChecker(ts.URL, WithHTTPTimeout(15*time.Millisecond))

	assertCheckerSuccessful(t, checker)
}

func TestHTTPChecker_Check_Timeout_Fail(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Nanosecond)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer ts.Close()

	checker := NewHTTPChecker(ts.URL, WithHTTPTimeout(3*time.Nanosecond))

	err := checker.Check()

	if err == nil {
		t.Fatal("Expected error, none received")
	}
}

func TestHTTPChecker_Check_Method(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer ts.Close()

	checker := NewHTTPChecker(ts.URL, WithHTTPMethod(http.MethodHead))

	assertCheckerSuccessful(t, checker)
}
