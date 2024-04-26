package log

import (
    stdlog "log"
    "os"
)

type Logger interface {
    Debugf(format string, v ...any)
    Infof(format string, v ...any)
    Errorf(format string, v ...any)
    Fatalf(format string, v ...any)
    Panicf(format string, v ...any)
}

type logger struct {
    *stdlog.Logger
}

var Log = NewLoggerWrapper(stdlog.New(os.Stderr, "[ exporter ] ", 0))

func NewLoggerWrapper(l *stdlog.Logger) Logger {
    return &logger{Logger: l}
}

func (l *logger) Debugf(format string, v ...interface{}) {
    l.Logger.Printf("DEBUG: "+format, v...)
}

func (l *logger) Infof(format string, v ...interface{}) {
    l.Logger.Printf("INFO: "+format, v...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
    l.Logger.Printf("ERROR: "+format, v...)
}

func (l *logger) Panicf(format string, v ...interface{}) {
    l.Logger.Panicf(format, v...)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
    l.Logger.Fatalf(format, v...)
}



