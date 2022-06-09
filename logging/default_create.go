package logging

import (
	"log"
	"os"
	"platform/config"
)

func NewDefaultLogger(cfg config.Configuration) Logger {
	//var level LogLevel
	//if configLevel, found := cfg.GetString("logging:level"); found {
	//	level = LogLevelFromString(configLevel)
	//}else {
	//	level = Debug
	//}
	level := LogLevelFromString(cfg.GetStringDefault("logging:level", "debug"))
	flags := log.LstdFlags
	return &DefaultLogger{
		minLevel: level,
		loggers: map[LogLevel]*log.Logger{
			Trace:       log.New(os.Stdout, "[TRACE] ", flags),
			Debug:       log.New(os.Stdout, "[DEBUG] ", flags),
			Information: log.New(os.Stdout, "[INFO] ", flags),
			Warning:     log.New(os.Stdout, "[WARN] ", flags),
			Fatal:       log.New(os.Stdout, "[FATAL] ", flags),
		},
		triggerPanic: true,
	}
}
