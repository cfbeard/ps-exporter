package main

import (
    "github.com/cfbeard/ps-exporter/api"
    "github.com/cfbeard/ps-exporter/collector"
    "github.com/cfbeard/ps-exporter/config"
    "github.com/cfbeard/ps-exporter/log"
)

func main() {
    err := config.ParseConfigArgs()
    if err != nil {
        log.Log.Fatalf("Failed to parse command-line args: %v", err)
    }

    collector.MustRegisterCollectors()

    err = api.StartServer(config.WebPort.Get())
    if err != nil {
        log.Log.Fatalf("Failed to start server: %v", err)
    }
}
