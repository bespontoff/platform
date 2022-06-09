package logging

import "strings"

type LogLevel int

const (
	Trace LogLevel = iota
	Debug
	Information
	Warning
	Fatal
	None
)

type Logger interface {
	Trace(string)
	Tracef(string, ...interface{})

	Debug(string)
	Debugf(string, ...interface{})

	Info(string)
	Infof(string, ...interface{})

	Warn(string)
	Warnf(string, ...interface{})

	Panic(string)
	Panicf(string, ...interface{})
}

func LogLevelFromString(value string) (level LogLevel) {
	switch strings.ToLower(value) {
	case "debug":
		level = Debug
	case "info", "information":
		level = Information
	case "warning", "warn":
		level = Warning
	case "fatal", "error":
		level = Fatal
	case "none":
		level = None
	default:
		level = Debug
	}
	return
}
