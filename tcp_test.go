package healthz_test

import (
	"testing"

	"net"
	"time"

	"github.com/goph/healthz"
)

func TestTCPChecker_Check(t *testing.T) {
	addr := "127.0.0.1:54321"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		t.Fatalf("Received unexpected error: %+v", err)
	}
	defer lis.Close()

	checker := healthz.NewTCPChecker(addr)

	assertCheckerSuccessful(t, checker)
}

func TestTCPChecker_Check_Fail(t *testing.T) {
	addr := "127.0.0.1:54321"

	checker := healthz.NewTCPChecker(addr)

	err := checker.Check()

	if err == nil {
		t.Fatal("Expected error, none received")
	}
}

func TestTCPChecker_Check_Timeout(t *testing.T) {
	addr := "127.0.0.1:54321"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		t.Fatalf("Received unexpected error: %+v", err)
	}
	defer lis.Close()

	checker := healthz.NewTCPChecker(addr, healthz.WithTCPTimeout(15*time.Millisecond))

	assertCheckerSuccessful(t, checker)
}

func TestTCPChecker_Check_Timeout_Fail(t *testing.T) {
	addr := "127.0.0.1:54321"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		t.Fatalf("Received unexpected error: %+v", err)
	}
	lis.Close()

	checker := healthz.NewTCPChecker(addr, healthz.WithTCPTimeout(3*time.Nanosecond))

	err = checker.Check()

	if err == nil {
		t.Fatal("Expected error, none received")
	}
}
