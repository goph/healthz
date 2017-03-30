/*
Package healthz provides tools for service health checks

The easiest way to setup Kubernetes-like liveness and readiness checks:

	liveness  := healthz.NewTCPChecker(":80")
	readiness := healthz.NewStatusChecker()

	// Exposes "/healthz" for liveness, "/readiness" for readiness checks
	handler := healthz.NewHealthServiceHandler(liveness, readiness)
	http.ListenAndServe(":8081", handler)

Setup the health service with custom paths:

	liveness  := healthz.NewTCPChecker(":80")
	readiness := healthz.NewStatusChecker()

	healthService := NewHealthService(livenessChecker, readinessChecker)
	mux := http.NewServeMux()

	mux.HandleFunc("/liveness", healthService.HealthStatus)
	mux.HandleFunc("/readiness", healthService.ReadinessStatus)
	http.ListenAndServe(":8081", mux)
*/
package healthz
