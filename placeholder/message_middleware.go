package placeholder

import (
	"errors"
	"io"
	"platform/config"
	"platform/pipeline"
	"platform/services"
)

type SimpleMessageComponent struct{}

func (c *SimpleMessageComponent) Init() {}

func (c *SimpleMessageComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	var cfg config.Configuration
	_ = services.GetService(&cfg)
	msg, ok := cfg.GetString("main:message")
	if ok {
		_, _ = io.WriteString(ctx.ResponseWriter, msg)
	} else {
		ctx.Error(errors.New("cannot find config setting"))
	}
	next(ctx)
}
