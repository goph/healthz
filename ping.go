package healthz

// Pinger is a commonly used interface to check if a connection is alive (used in sql.DB for example).
type Pinger interface {
	Ping() error
}

// PingChecker checks if a Pinger responds without an error.
type PingChecker struct {
	pinger Pinger
}

// NewPingChecker creates a new PingChecker with a Pinger.
func NewPingChecker(pinger Pinger) *PingChecker {
	return &PingChecker{pinger}
}

// Check implements the Checker interface and checks a resource status by pinging it.
func (c *PingChecker) Check() error {
	return c.pinger.Ping()
}
