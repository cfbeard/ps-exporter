package config

import (
    "flag"
)

type configArgs struct {
	mock    bool
	port    int
	process bool
}

func ParseConfigArgs() error {
	config := configArgs{}

	flag.BoolVar(&config.process, "ps-exporter", PsCollector.Get(), PsCollector.Description)
	flag.IntVar(&config.port, "port", WebPort.Get(), WebPort.Description)

	flag.Parse()

	PsCollector.Set(config.process)
	WebPort.Set(config.port)

	return nil
}
