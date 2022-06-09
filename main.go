package main

import (
	"platform/config"
	"platform/logging"
)

func main() {
	var cfg config.Configuration
	var err error

	if cfg, err = config.Load("config.json"); err != nil {
		panic(err)
	}
	logger := logging.NewDefaultLogger(cfg)

	writeMessage(logger, cfg)
}

func writeMessage(logger logging.Logger, cfg config.Configuration) {
	section, ok := cfg.GetSection("main")
	if !ok {
		logger.Warn("Cannot find main section in config file.")
	}
	if message, ok := section.GetString("message"); ok {
		logger.Info(message)
	} else {
		logger.Warn("Cannot find message in main section of config file.")
	}
}
