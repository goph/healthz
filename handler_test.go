package healthz_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sagikazarmark/healthz"
)

func assertStatusCodesEqual(t *testing.T, got int, want int) {
	if got != want {
		t.Fatalf("Expected status code %d, received %d", got, want)
	}
}

func TestNewHealthServiceHandler(t *testing.T) {
	checker := new(healthz.AlwaysSuccessChecker)
	livenessChecker := healthz.NewCheckers(checker)
	readinessChecker := healthz.NewCheckers(checker)

	mux := healthz.NewHealthServiceHandler(livenessChecker, readinessChecker)

	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assertStatusCodesEqual(t, http.StatusOK, w.Code)

	req = httptest.NewRequest("GET", "/readiness", nil)
	w = httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assertStatusCodesEqual(t, http.StatusOK, w.Code)
}

func TestHealthService_HealthStatus(t *testing.T) {
	checker := new(healthz.AlwaysSuccessChecker)
	livenessChecker := healthz.NewCheckers(checker)
	readinessChecker := new(healthz.Checkers)

	service := healthz.NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", service.HealthStatus)

	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assertStatusCodesEqual(t, http.StatusOK, w.Code)
}

func TestHealthService_HealthStatus_Fail(t *testing.T) {
	checker := new(healthz.AlwaysFailureChecker)
	livenessChecker := healthz.NewCheckers(checker)
	readinessChecker := new(healthz.Checkers)

	service := healthz.NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", service.HealthStatus)

	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assertStatusCodesEqual(t, http.StatusServiceUnavailable, w.Code)
}

func TestHealthService_ReadinessStatus(t *testing.T) {
	checker := new(healthz.AlwaysSuccessChecker)
	livenessChecker := new(healthz.Checkers)
	readinessChecker := healthz.NewCheckers(checker)

	service := healthz.NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("/readiness", service.ReadinessStatus)

	req := httptest.NewRequest("GET", "/readiness", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assertStatusCodesEqual(t, http.StatusOK, w.Code)
}

func TestHealthService_ReadinessStatus_Fail(t *testing.T) {
	checker := new(healthz.AlwaysFailureChecker)
	livenessChecker := new(healthz.Checkers)
	readinessChecker := healthz.NewCheckers(checker)

	service := healthz.NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("/readiness", service.ReadinessStatus)

	req := httptest.NewRequest("GET", "/readiness", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assertStatusCodesEqual(t, http.StatusServiceUnavailable, w.Code)
}
