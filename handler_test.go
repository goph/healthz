package healthz

import (
	"net/http/httptest"
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
)

func TestHealthService_HealthStatus(t *testing.T) {
	healthChecker := new(HealthCheckerMock)
	livenessProbe := NewProbe(healthChecker)
	readinessProbe := new(Probe)

	healthChecker.On("Ping").Return(nil)

	service := NewHealthService(livenessProbe, readinessProbe)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", service.HealthStatus)

	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	healthChecker.AssertExpectations(t)
}

func TestHealthService_HealthStatus_Fail(t *testing.T) {
	healthChecker := new(HealthCheckerMock)
	livenessProbe := NewProbe(healthChecker)
	readinessProbe := new(Probe)

	healthChecker.On("Ping").Return(ErrHealthCheckFailed)

	service := NewHealthService(livenessProbe, readinessProbe)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", service.HealthStatus)

	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	healthChecker.AssertExpectations(t)
}

func TestHealthService_ReadinessStatus(t *testing.T) {
	healthChecker := new(HealthCheckerMock)
	livenessProbe := new(Probe)
	readinessProbe := NewProbe(healthChecker)

	healthChecker.On("Ping").Return(nil)

	service := NewHealthService(livenessProbe, readinessProbe)
	mux := http.NewServeMux()
	mux.HandleFunc("/readiness", service.ReadinessStatus)

	req := httptest.NewRequest("GET", "/readiness", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	healthChecker.AssertExpectations(t)
}

func TestHealthService_ReadinessStatus_Fail(t *testing.T) {
	healthChecker := new(HealthCheckerMock)
	livenessProbe := new(Probe)
	readinessProbe := NewProbe(healthChecker)

	healthChecker.On("Ping").Return(ErrHealthCheckFailed)

	service := NewHealthService(livenessProbe, readinessProbe)
	mux := http.NewServeMux()
	mux.HandleFunc("/readiness", service.ReadinessStatus)

	req := httptest.NewRequest("GET", "/readiness", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	healthChecker.AssertExpectations(t)
}
