package healthz

import (
	"testing"

	"errors"
)

type PingerMock struct {
	err error
}

func (p *PingerMock) Ping() error {
	return p.err
}

func TestPingChecker_Check(t *testing.T) {
	checker := NewPingChecker(&PingerMock{})

	assertCheckerSuccessful(t, checker)
}

func TestPingChecker_Check_Fail(t *testing.T) {
	checker := NewPingChecker(&PingerMock{errors.New("ping failed")})

	err := checker.Check()

	if err == nil {
		t.Fatal("Expected error, none received")
	}
}
