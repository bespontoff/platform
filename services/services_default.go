package services

import (
	"platform/config"
	"platform/logging"
	"platform/templates"
)

func RegisterDefaultServices() {
	err := AddSingleton(func() (c config.Configuration) {
		c, loadErr := config.Load("config.json")
		if loadErr != nil {
			panic(loadErr)
		}
		return
	})
	err = AddSingleton(func(cfg config.Configuration) logging.Logger {
		return logging.NewDefaultLogger(cfg)
	})
	if err != nil {
		panic(err)
	}
	err = AddSingleton(func(c config.Configuration) templates.TemplateExecutor {
		_ = templates.LoadTemplates(c)
		return &templates.LayoutTemplateProcessor{}
	})
	if err != nil {
		panic(err)
	}
}
