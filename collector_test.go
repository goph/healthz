package healthz

import (
	"testing"
)

func TestCollector_RegisterChecker(t *testing.T) {
	checker := &AlwaysSuccessChecker{}
	collector := make(Collector)

	collector.RegisterChecker("test", checker)

	handler := collector.Handler("test")

	testHandler(handler, true, t)
}

func TestCollector_Handler_NotFound_Success(t *testing.T) {
	collector := make(Collector)

	handler := collector.Handler("not_found")

	testHandler(handler, true, t)
}
