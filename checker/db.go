package healthz

import "database/sql"

// DbHealthChecker checks if a database is available through builtin database/sql
type DbHealthChecker struct {
	db *sql.DB
}

// Type returns the name of the database checker
func (d *DbHealthChecker) Type() string {
	return "DatabasePing"
}

// Ping checks the database status by pinging it
func (d *DbHealthChecker) Ping() error {
	return d.db.Ping()
}
