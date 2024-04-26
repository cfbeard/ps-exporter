package api

import (
    "fmt"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

func StartServer(port int) error {
    http.Handle("/metrics", promhttp.Handler())

    return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
