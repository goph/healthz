package healthz_test

import (
	"testing"

	"github.com/goph/healthz"
)

func TestCollector_RegisterChecker(t *testing.T) {
	checker := &healthz.AlwaysSuccessChecker{}
	collector := make(healthz.Collector)

	collector.RegisterChecker("test", checker)

	handler := collector.Handler("test")

	testHandler(handler, true, t)
}

func TestCollector_Handler_NotFound_Success(t *testing.T) {
	collector := make(healthz.Collector)

	handler := collector.Handler("not_found")

	testHandler(handler, true, t)
}
