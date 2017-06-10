package healthz

import (
	"testing"
)

func TestCollector_RegisterChecker(t *testing.T) {
	checker := &AlwaysSuccessChecker{}
	collector := make(Collector)

	collector.RegisterChecker("test", checker)

	healthService := collector.NewHealthService()

	testHealthService(healthService, true, t)
}
