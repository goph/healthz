package healthz

import (
	"net/http/httptest"
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
)

func TestHealthService_HealthStatus(t *testing.T) {
	checker := new(CheckerMock)
	livenessChecker := NewCheckers(checker)
	readinessChecker := new(Checkers)

	checker.On("Check").Return(nil)

	service := NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", service.HealthStatus)

	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	checker.AssertExpectations(t)
}

func TestHealthService_HealthStatus_Fail(t *testing.T) {
	checker := new(CheckerMock)
	livenessChecker := NewCheckers(checker)
	readinessChecker := new(Checkers)

	checker.On("Check").Return(ErrCheckFailed)

	service := NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", service.HealthStatus)

	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	checker.AssertExpectations(t)
}

func TestHealthService_ReadinessStatus(t *testing.T) {
	checker := new(CheckerMock)
	livenessChecker := new(Checkers)
	readinessChecker := NewCheckers(checker)

	checker.On("Check").Return(nil)

	service := NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("/readiness", service.ReadinessStatus)

	req := httptest.NewRequest("GET", "/readiness", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	checker.AssertExpectations(t)
}

func TestHealthService_ReadinessStatus_Fail(t *testing.T) {
	checker := new(CheckerMock)
	livenessChecker := new(Checkers)
	readinessChecker := NewCheckers(checker)

	checker.On("Check").Return(ErrCheckFailed)

	service := NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("/readiness", service.ReadinessStatus)

	req := httptest.NewRequest("GET", "/readiness", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	checker.AssertExpectations(t)
}
