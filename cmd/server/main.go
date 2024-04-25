package main

import (
    "github.com/cfbeard/ps-exporter/api"
    "github.com/cfbeard/ps-exporter/config"
    "github.com/cfbeard/ps-exporter/log"
)

func main() {

    err := api.StartServer(config.WebPort.Get())
    if err != nil {
        log.Log.Fatalf("Failed to start server: %v", err)
    }
    log.Log.Infof("Server starter on port: %d", config.WebPort.Get())
}
