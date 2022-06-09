package config

import "strings"

type DefaultConfig struct {
	configData map[string]interface{}
}

func (c DefaultConfig) GetString(name string) (value string, found bool) {
	val, found := c.get(name)
	if found {
		value = val.(string)
	}
	return
}

func (c DefaultConfig) GetInt(name string) (value int, found bool) {
	val, found := c.get(name)
	if found {
		value = int(val.(float64))
	}
	return
}

func (c DefaultConfig) GetFloat(name string) (value float64, found bool) {
	val, found := c.get(name)
	if found {
		value = val.(float64)
	}
	return
}

func (c DefaultConfig) GetBool(name string) (value bool, found bool) {
	val, found := c.get(name)
	if found {
		value = val.(bool)
	}
	return
}

func (c DefaultConfig) GetStringDefault(name string, defaultValue string) (value string) {
	val, found := c.get(name)
	if found {
		value = val.(string)
	} else {
		value = defaultValue
	}
	return
}

func (c DefaultConfig) GetIntDefault(name string, defaultValue int) (value int) {
	val, found := c.get(name)
	if found {
		value = int(val.(float64))
	} else {
		value = defaultValue
	}
	return
}

func (c DefaultConfig) GetFloatDefault(name string, defaultValue float64) (value float64) {
	val, found := c.get(name)
	if found {
		value = val.(float64)
	} else {
		value = defaultValue
	}
	return
}

func (c DefaultConfig) GetBoolDefault(name string, defaultValue bool) (value bool) {
	val, found := c.get(name)
	if found {
		value = val.(bool)
	} else {
		value = defaultValue
	}
	return
}

func (c DefaultConfig) get(name string) (result interface{}, found bool) {
	data := c.configData
	for _, key := range strings.Split(name, ":") {
		result, found = data[key]
		if section, ok := result.(map[string]interface{}); ok && found {
			data = section
		} else {
			return
		}
	}
	return
}

func (c DefaultConfig) GetSection(name string) (section Configuration, found bool) {
	value, found := c.get(name)
	if found {
		if sectionData, ok := value.(map[string]interface{}); ok {
			section = &DefaultConfig{configData: sectionData}
		}
	}
	return
}
