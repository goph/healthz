package healthz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProbe(t *testing.T) {
	checker := new(CheckerMock)

	probe := NewProbe(checker)

	assert.Equal(t, checker, probe.checkers[0])
}

func TestProbe_Check(t *testing.T) {
	checker := new(CheckerMock)

	checker.On("Check").Return(nil)

	probe := NewProbe(checker)

	assert.NoError(t, probe.Check())
	checker.AssertExpectations(t)
}

func TestProbe_Check_Fail(t *testing.T) {
	checker1 := new(CheckerMock)
	checker2 := new(CheckerMock)

	checker1.On("Check").Return(nil)
	checker2.On("Check").Return(ErrCheckFailed)

	probe := NewProbe(checker1, checker2)

	err := probe.Check()

	assert.Error(t, err)
	assert.Equal(t, ErrCheckFailed, err)
	checker1.AssertExpectations(t)
	checker2.AssertExpectations(t)
}
