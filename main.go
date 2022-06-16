package main

import (
	"platform/config"
	"platform/logging"
	"platform/services"
)

func main() {
	services.RegisterDefaultServices()
	if _, err := services.Call(writeMessage); err != nil {
		panic(err)
	}
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
