package healthz

import (
	"testing"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func TestHealthService_Handler_Success(t *testing.T) {
	checker := &AlwaysSuccessChecker{}

	healthService := HealthService{
		"test": checker,
	}

	testHealthService(healthService, true, t)
}

func TestHealthService_Handler_NotFound_Success(t *testing.T) {
	healthService := HealthService{}

	testHealthService(healthService, true, t)
}

func TestHealthService_Handler_Failure(t *testing.T) {
	checker := &AlwaysFailureChecker{}

	healthService := HealthService{
		"test": checker,
	}

	testHealthService(healthService, false, t)
}

func testHealthService(healthService HealthService, success bool, t *testing.T) {
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
