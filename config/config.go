package config

import (
	"encoding/json"
	"os"
	"platform/logging"
)

type Configuration interface {
	GetString(name string) (value string, found bool)
	GetInt(name string) (value int, found bool)
	GetFloat(name string) (value float64, found bool)
	GetBool(name string) (value bool, found bool)

	GetStringDefault(name string, defaultValue string) (value string)
	GetIntDefault(name string, defaultValue int) (value int)
	GetFloatDefault(name string, defaultValue float64) (value float64)
	GetBoolDefault(name string, defaultValue bool) (value bool)

	GetSection(name string) (section Configuration, found bool)
}

var logger logging.Logger

func init() {
	logger = logging.NewDefaultLogger(logging.Debug)
}

func Load(filename string) (config Configuration, err error) {
	file, err := os.Open(filename)
	if err != nil {
		logger.Panicf("Cant open config file. %v", filename)
	}
	decoder := json.NewDecoder(file)
	m := map[string]interface{}{}
	if err = decoder.Decode(&m); err == nil {
		config = &DefaultConfig{configData: m}
	}
	return
}
