/*
Package healthz provides tools for service health checks

The easiest way to setup Kubernetes-like liveness and readiness checks:

	liveness  := healthz.NewTCPChecker(":80")
	readiness := healthz.NewStatusChecker()

	healthService := healthz.HealthService{
		healthz.LivenessCheck: liveness,
		healthz.ReadinessCheck: readiness,
	}

	mux := http.NewServeMux()
	mux.Handle("/healthz", healthService.Handler(healthz.LivenessCheck))
	mux.Handle("/readiness", healthService.Handler(healthz.ReadinessCheck))

	http.ListenAndServe(":8081", mux)

Setup the health service using a collector:

	collector := &healthz.Collector{}

	liveness := healthz.NewTCPChecker(":80")
	collector.RegisterChecker(healthz.LivenessCheck, liveness)

	readiness := healthz.NewStatusChecker()
	collector.RegisterChecker(healthz.ReadinessCheck, readiness)

	healthService := collector.NewHealthService()
	mux := http.NewServeMux()

	mux.HandleFunc("/liveness", healthService.HealthStatus)
	mux.HandleFunc("/readiness", healthService.ReadinessStatus)
	http.ListenAndServe(":8081", mux)
*/
package healthz
