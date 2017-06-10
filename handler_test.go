package healthz

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Handler_Success(t *testing.T) {
	checker := &AlwaysSuccessChecker{}
	handler := NewHandler(checker)

	testHandler(handler, true, t)
}

func TestHandler_Handler_Failure(t *testing.T) {
	checker := &AlwaysFailureChecker{}
	handler := NewHandler(checker)

	testHandler(handler, false, t)
}

func testHandler(handler http.Handler, success bool, t *testing.T) {
	ts := httptest.NewServer(handler)
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
