package logging

import (
	"fmt"
	"log"
)

type DefaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel]*log.Logger
	triggerPanic bool
}

func (l DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

func (l DefaultLogger) write(level LogLevel, message string) {
	if l.minLevel <= level {
		err := l.loggers[level].Output(2, message)
		if err != nil {
			panic(err)
		}
	}
}

func (l DefaultLogger) Trace(msg string) {
	l.write(Trace, msg)
}

func (l DefaultLogger) Tracef(tmpl string, vals ...interface{}) {
	l.write(Trace, fmt.Sprintf(tmpl, vals...))
}

func (l DefaultLogger) Debug(msg string) {
	l.write(Debug, msg)
}

func (l DefaultLogger) Debugf(tmpl string, vals ...interface{}) {
	l.write(Debug, fmt.Sprintf(tmpl, vals...))
}

func (l DefaultLogger) Warn(msg string) {
	l.write(Warning, msg)
}

func (l DefaultLogger) Warnf(tmpl string, vals ...interface{}) {
	l.write(Warning, fmt.Sprintf(tmpl, vals...))
}

func (l DefaultLogger) Info(msg string) {
	l.write(Information, msg)
}

func (l DefaultLogger) Infof(tmpl string, vals ...interface{}) {
	l.write(Information, fmt.Sprintf(tmpl, vals...))
}

func (l DefaultLogger) Panic(msg string) {
	l.write(Fatal, msg)
	if l.triggerPanic {
		panic(msg)
	}
}

func (l DefaultLogger) Panicf(tmpl string, vals ...interface{}) {
	l.write(Fatal, fmt.Sprintf(tmpl, vals...))
	if l.triggerPanic {
		panic(fmt.Sprintf(tmpl, vals...))
	}
}
