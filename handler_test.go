package healthz

import (
	"net/http/httptest"
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
)

func TestHealthService_HealthStatus(t *testing.T) {
	checker := new(AlwaysSuccessChecker)
	livenessChecker := NewCheckers(checker)
	readinessChecker := new(Checkers)

	service := NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", service.HealthStatus)

	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHealthService_HealthStatus_Fail(t *testing.T) {
	checker := new(AlwaysFailureChecker)
	livenessChecker := NewCheckers(checker)
	readinessChecker := new(Checkers)

	service := NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", service.HealthStatus)

	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func TestHealthService_ReadinessStatus(t *testing.T) {
	checker := new(AlwaysSuccessChecker)
	livenessChecker := new(Checkers)
	readinessChecker := NewCheckers(checker)

	service := NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("/readiness", service.ReadinessStatus)

	req := httptest.NewRequest("GET", "/readiness", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHealthService_ReadinessStatus_Fail(t *testing.T) {
	checker := new(AlwaysFailureChecker)
	livenessChecker := new(Checkers)
	readinessChecker := NewCheckers(checker)

	service := NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("/readiness", service.ReadinessStatus)

	req := httptest.NewRequest("GET", "/readiness", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}
