package healthz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProbe(t *testing.T) {
	healthChecker := new(HealthCheckerMock)

	probe := NewProbe(healthChecker)

	assert.Equal(t, healthChecker, probe.checkers[0])
}

func TestProbe_Check(t *testing.T) {
	healthChecker := new(HealthCheckerMock)

	healthChecker.On("Ping").Return(nil)

	probe := NewProbe(healthChecker)

	assert.NoError(t, probe.Check())
	healthChecker.AssertExpectations(t)
}

func TestProbe_Check_Fail(t *testing.T) {
	healthChecker1 := new(HealthCheckerMock)
	healthChecker2 := new(HealthCheckerMock)

	healthChecker1.On("Ping").Return(nil)
	healthChecker2.On("Ping").Return(ErrHealthCheckFailed)

	probe := NewProbe(healthChecker1, healthChecker2)

	err := probe.Check()

	assert.Error(t, err)
	assert.Equal(t, ErrHealthCheckFailed, err)
	healthChecker1.AssertExpectations(t)
	healthChecker2.AssertExpectations(t)
}
