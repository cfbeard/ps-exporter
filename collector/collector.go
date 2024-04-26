package collector

import (
    "github.com/prometheus/client_golang/prometheus"

    "github.com/cfbeard/ps-exporter/config"
)

const (
    processStats = "process.stats"
)

var (
    psFeatures = []string{"mem_percent", "cpu_percent"}
)

type collectors struct {
    Implemented map[string]prometheus.Collector
}

var Collectors = &collectors{
    Implemented: make(map[string]prometheus.Collector),
}

func addCollector(key string, factory func() prometheus.Collector) {
    Collectors.Implemented[key] = factory()
}

func MustRegisterCollectors() {
    for key, c := range Collectors.Implemented {
        if config.Exporters[key].Get() {
            prometheus.MustRegister(c)
        }
    }
}

func GetMonitoringTasks() map[string][]string {
    monTasks := map[string][]string {
        processStats: psFeatures,
    }

    return monTasks
}