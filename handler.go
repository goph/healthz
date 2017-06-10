package healthz

import (
	"net/http"
)

// Handler is responsible for serving the health check over HTTP.
type Handler struct {
	checker Checker
}

// NewHandler returns a new HTTP handler for a checker.
func NewHandler(checker Checker) http.Handler {
	return &Handler{checker}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If the checker fails, we return an error
	if err := h.checker.Check(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("error"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
