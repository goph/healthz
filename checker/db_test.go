package healthz

import (
	"testing"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
