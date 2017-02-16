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

type CheckerMock struct {
	mock.Mock
}

func (c *CheckerMock) Check() error {
	args := c.Called()

	return args.Error(0)
}

func TestStatusChecker_Check(t *testing.T) {
	checker := NewStatusChecker(true)

	assert.NoError(t, checker.Check())
}

func TestStatusChecker_Check_Fail(t *testing.T) {
	checker := NewStatusChecker(false)

	err := checker.Check()

	assert.Error(t, err)
	assert.Equal(t, ErrCheckFailed, err)
}

func TestStatusChecker_SetStatus(t *testing.T) {
	checker := NewStatusChecker(false)

	checker.SetStatus(true)

	assert.NoError(t, checker.Check())
}

// func TestDbChecker_Check(t *testing.T) {
// 	db, _ := sql.Open("mysql", "obviously_wrong")

// 	checker := &DbChecker{
// 		db: db,
// 	}

// 	//assert.NoError(t, checker.Check())
// }

func TestDbChecker_Check_Fail(t *testing.T) {
	db, err := sql.Open("mysql", "user:password@/dbname")

	require.NoError(t, err)

	checker := NewDbChecker(db)

	err = checker.Check()

	assert.Error(t, err)
}

func TestHTTPChecker_Check(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer ts.Close()

	checker := NewHTTPChecker(ts.URL)

	assert.NoError(t, checker.Check())
}

func TestHTTPChecker_Check_Fail(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("error"))
	}))
	defer ts.Close()

	checker := NewHTTPChecker(ts.URL)

	err := checker.Check()

	assert.Error(t, err)
	assert.Equal(t, ErrCheckFailed, err)
}
