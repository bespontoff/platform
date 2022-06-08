package main

import (
	"platform/logging"
)

func main() {
	logger := logging.NewDefaultLogger(logging.Information)
	writeMessage(logger)
}

func writeMessage(logger logging.Logger) {
	logger.Info("go man")
}
