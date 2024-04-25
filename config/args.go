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

	flag.BoolVar(&config.process, "ps-exporter", PsExporter.Get(), PsExporter.Description)
	flag.IntVar(&config.port, "port", WebPort.Get(), WebPort.Description)

	flag.Parse()

	PsExporter.Set(config.process)
	WebPort.Set(config.port)

	return nil
}
