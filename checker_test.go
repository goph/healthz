package healthz

import (
	"database/sql"
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

func (hc *HealthCheckerMock) Ping() error {
	args := hc.Called()

	return args.Error(0)
}

func TestStatusHealthChecker_Type(t *testing.T) {
	healthChecker := &StatusHealthChecker{
		Status: true,
	}

	assert.Equal(t, "Status", healthChecker.Type())
}

func TestStatusHealthChecker_Ping(t *testing.T) {
	healthChecker := &StatusHealthChecker{
		Status: true,
	}

	assert.NoError(t, healthChecker.Ping())
}

func TestStatusHealthChecker_Ping_Fail(t *testing.T) {
	healthChecker := &StatusHealthChecker{
		Status: false,
	}

	err := healthChecker.Ping()

	assert.Error(t, err)
	assert.Equal(t, ErrHealthCheckFailed, err)
}

func TestDbHealthChecker_Type(t *testing.T) {
	db, _ := sql.Open("mysql", "obviously_wrong")

	healthChecker := &DbHealthChecker{
		db: db,
	}

	assert.Equal(t, "DatabasePing", healthChecker.Type())
}

// func TestDbHealthChecker_Ping(t *testing.T) {
// 	db, _ := sql.Open("mysql", "obviously_wrong")

// 	healthChecker := &DbHealthChecker{
// 		db: db,
// 	}

// 	//assert.NoError(t, healthChecker.Ping())
// }

func TestDbHealthChecker_Ping_Fail(t *testing.T) {
	db, err := sql.Open("mysql", "user:password@/dbname")

	require.NoError(t, err)

	healthChecker := &DbHealthChecker{
		db: db,
	}

	err = healthChecker.Ping()

	assert.Error(t, err)
}
