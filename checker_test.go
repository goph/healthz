package healthz

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type HealthCheckerMock struct {
	mock.Mock
}

func (hc *HealthCheckerMock) Type() string {
	return "Mock"
}

func (hc *HealthCheckerMock) Ping() error {
	args := hc.Called()

	return args.Error(0)
}

func TestStatusHealthChecker_Type(t *testing.T) {
	healthChecker := &StatusHealthChecker{
		Status: true,
	}

	assert.Equal(t, "Status", healthChecker.Type())
}

func TestStatusHealthChecker_Ping(t *testing.T) {
	healthChecker := &StatusHealthChecker{
		Status: true,
	}

	assert.NoError(t, healthChecker.Ping())
}

func TestStatusHealthChecker_Ping_Fail(t *testing.T) {
	healthChecker := &StatusHealthChecker{
		Status: false,
	}

	err := healthChecker.Ping()

	assert.Error(t, err)
	assert.Equal(t, ErrHealthCheckFailed, err)
}
