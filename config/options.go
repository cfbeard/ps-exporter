package config

// Inspired by: https://github.com/edulinq/autograder-server/tree/main/config

var (
	// Switches
	PsExporter = NewBoolOption("ps.exporter", true, "Whether to enable the ps-exporter. Default is true.")

	// Web
	WebPort = NewIntOption("web.port", 9091, "What port to run the exporter on. Default is 9091.")
)

var Exporters = map[string]*BoolOption{
	PsExporter.Key: PsExporter,
}
