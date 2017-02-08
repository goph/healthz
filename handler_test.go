package healthz

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthService_HealthStatus(t *testing.T) {
	statusChecker := &StatusHealthChecker{true}
	livenessProbe := NewProbe(statusChecker)
	readinessProbe := new(Probe)

	service := NewHealthService(livenessProbe, readinessProbe)

	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	service.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestHealthService_ReadinessStatus(t *testing.T) {
	statusChecker := &StatusHealthChecker{true}
	livenessProbe := new(Probe)
	readinessProbe := NewProbe(statusChecker)

	service := NewHealthService(livenessProbe, readinessProbe)

	req := httptest.NewRequest("GET", "/readiness", nil)
	w := httptest.NewRecorder()

	service.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
