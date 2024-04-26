package api

import (
    "fmt"
    "github.com/cfbeard/ps-exporter/config"
    "github.com/cfbeard/ps-exporter/log"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

func StartServer(port int) error {
    http.Handle("/metrics", promhttp.Handler())

    log.Log.Infof("Server starter on port: %d", config.WebPort.Get())
    return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
