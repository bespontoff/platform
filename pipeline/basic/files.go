package basic

import (
	"net/http"
	"platform/config"
	"platform/pipeline"
	"platform/services"
	"strings"
)

type StaticFileComponent struct {
	urlPrefix     string
	stdLibHandler http.Handler
}

func (sfc *StaticFileComponent) Init() {
	var cfg config.Configuration
	_ = services.GetService(&cfg)
	sfc.urlPrefix = cfg.GetStringDefault("files:urlprefix", "/files/")
	path, ok := cfg.GetString("files:path")
	if ok {
		sfc.stdLibHandler = http.StripPrefix(sfc.urlPrefix, http.FileServer(http.Dir(path)))
	} else {
		panic("cannot load file configuration settings")
	}
}

func (sfc *StaticFileComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	if !strings.EqualFold(ctx.Request.URL.Path, sfc.urlPrefix) && strings.HasPrefix(ctx.Request.URL.Path, sfc.urlPrefix) {
		sfc.stdLibHandler.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	} else {
		next(ctx)
	}
}
