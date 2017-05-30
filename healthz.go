package healthz

import "net/http"

// HealthService wraps Checkers and exposes them as HTTP handlers.
type HealthService map[string]Checker

// Handler returns an http.Handler for a check.
// If a check is not found the returned handler will always return success.
func (s HealthService) Handler(check string) http.Handler {
	checker, ok := s[check]
	if !ok {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		})
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the checker fails, we return an error
		if err := checker.Check(); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("error"))

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}
