package healthz_test

import (
	"testing"

	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/goph/healthz"
)

func TestHealthService_Handler_Success(t *testing.T) {
	checker := &healthz.AlwaysSuccessChecker{}

	healthService := healthz.HealthService{
		"test": checker,
	}

	testHealthService(healthService, true, t)
}

func TestHealthService_Handler_NotFound_Success(t *testing.T) {
	healthService := healthz.HealthService{}

	testHealthService(healthService, true, t)
}

func TestHealthService_Handler_Failure(t *testing.T) {
	checker := &healthz.AlwaysFailureChecker{}

	healthService := healthz.HealthService{
		"test": checker,
	}

	testHealthService(healthService, false, t)
}

func testHealthService(healthService healthz.HealthService, success bool, t *testing.T) {
	ts := httptest.NewServer(healthService.Handler("test"))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	if success {
		if string(body) != "ok" || res.StatusCode != http.StatusOK {
			t.Error("health check was unsuccessful, expected success")
		}
	} else {
		if string(body) != "error" || res.StatusCode != http.StatusServiceUnavailable {
			t.Error("health check was successful, expected failure")
		}
	}
}
