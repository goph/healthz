package healthz_test

import (
	"testing"

	"github.com/sagikazarmark/healthz"
)

func TestCollector_RegisterChecker(t *testing.T) {
	checker := &healthz.AlwaysSuccessChecker{}
	collector := make(healthz.Collector)

	collector.RegisterChecker("test", checker)

	healthService := collector.NewHealthService()

	testHealthService(healthService, true, t)
}
