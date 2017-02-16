package healthz_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sagikazarmark/healthz"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckers_Check(t *testing.T) {
	checker1 := new(healthz.AlwaysSuccessChecker)
	checker2 := new(healthz.AlwaysSuccessChecker)

	checkers := healthz.NewCheckers(checker1, checker2)

	assert.NoError(t, checkers.Check())
}

func TestCheckers_Check_Fail(t *testing.T) {
	checker1 := new(healthz.AlwaysSuccessChecker)
	checker2 := new(healthz.AlwaysFailureChecker)

	checkers := healthz.NewCheckers(checker1, checker2)

	err := checkers.Check()

	assert.Error(t, err)
	assert.Equal(t, healthz.ErrCheckFailed, err)
}

func TestStatusChecker_Check(t *testing.T) {
	checker := healthz.NewStatusChecker(healthz.Healthy)

	assert.NoError(t, checker.Check())
}

func TestStatusChecker_Check_Fail(t *testing.T) {
	checker := healthz.NewStatusChecker(healthz.Unhealthy)

	err := checker.Check()

	assert.Error(t, err)
	assert.Equal(t, healthz.ErrCheckFailed, err)
}

func TestStatusChecker_SetStatus(t *testing.T) {
	checker := healthz.NewStatusChecker(healthz.Unhealthy)

	checker.SetStatus(healthz.Healthy)

	assert.NoError(t, checker.Check())
}

// func TestDbChecker_Check(t *testing.T) {
// 	db, _ := sql.Open("mysql", "obviously_wrong")

// 	checker := healthz.NewDbChecker(db)

// 	//assert.NoError(t, checker.Check())
// }

func TestDbChecker_Check_Fail(t *testing.T) {
	db, err := sql.Open("mysql", "user:password@/dbname")

	require.NoError(t, err)

	checker := healthz.NewDbChecker(db)

	err = checker.Check()

	assert.Error(t, err)
}

func TestHTTPChecker_Check(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer ts.Close()

	checker := healthz.NewHTTPChecker(ts.URL)

	assert.NoError(t, checker.Check())
}

func TestHTTPChecker_Check_Fail(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("error"))
	}))
	defer ts.Close()

	checker := healthz.NewHTTPChecker(ts.URL)

	err := checker.Check()

	assert.Error(t, err)
	assert.Equal(t, healthz.ErrCheckFailed, err)
}
