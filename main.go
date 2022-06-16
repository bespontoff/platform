package main

import (
	"platform/config"
	"platform/logging"
	"platform/services"
)

func main() {
	var cfg config.Configuration
	var logger logging.Logger

	services.RegisterDefaultServices()
	if err := services.GetService(&cfg); err != nil {
		panic(err)
	}
	if err := services.GetService(&logger); err != nil {
		panic(err)
	}

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
