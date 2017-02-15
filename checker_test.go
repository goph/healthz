package healthz

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type HealthCheckerMock struct {
	mock.Mock
}

func (hc *HealthCheckerMock) Type() string {
	return "Mock"
}

func (hc *HealthCheckerMock) Check() error {
	args := hc.Called()

	return args.Error(0)
}

func TestStatusHealthChecker_Type(t *testing.T) {
	healthChecker := &StatusHealthChecker{}

	assert.Equal(t, "Status", healthChecker.Type())
}

func TestStatusHealthChecker_Check(t *testing.T) {
	healthChecker := NewStatusHealthChecker(true)

	assert.NoError(t, healthChecker.Check())
}

func TestStatusHealthChecker_Check_Fail(t *testing.T) {
	healthChecker := NewStatusHealthChecker(false)

	err := healthChecker.Check()

	assert.Error(t, err)
	assert.Equal(t, ErrCheckFailed, err)
}

func TestStatusHealthChecker_SetStatus(t *testing.T) {
	healthChecker := NewStatusHealthChecker(false)

	healthChecker.SetStatus(true)

	assert.NoError(t, healthChecker.Check())
}

func TestDbHealthChecker_Type(t *testing.T) {
	healthChecker := &DbHealthChecker{}

	assert.Equal(t, "DatabasePing", healthChecker.Type())
}

// func TestDbHealthChecker_Check(t *testing.T) {
// 	db, _ := sql.Open("mysql", "obviously_wrong")

// 	healthChecker := &DbHealthChecker{
// 		db: db,
// 	}

// 	//assert.NoError(t, healthChecker.Check())
// }

func TestDbHealthChecker_Check_Fail(t *testing.T) {
	db, err := sql.Open("mysql", "user:password@/dbname")

	require.NoError(t, err)

	healthChecker := NewDbHealthChecker(db)

	err = healthChecker.Check()

	assert.Error(t, err)
}

func TestHTTPHealthChecker_Type(t *testing.T) {
	healthChecker := &HTTPHealthChecker{}

	assert.Equal(t, "HTTPPing", healthChecker.Type())
}

func TestHTTPHealthChecker_Check(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer ts.Close()

	healthChecker := NewHTTPHealthChecker(ts.URL)

	assert.NoError(t, healthChecker.Check())
}

func TestHTTPHealthChecker_Check_Fail(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("error"))
	}))
	defer ts.Close()

	healthChecker := NewHTTPHealthChecker(ts.URL)

	err := healthChecker.Check()

	assert.Error(t, err)
	assert.Equal(t, ErrCheckFailed, err)
}
