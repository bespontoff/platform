package pipeline

import (
	"net/http"
	"platform/services"
	"reflect"
)

type RequestPipeline func(*ComponentContext)

var emptyPipeline RequestPipeline = func(*ComponentContext) {}

func (pl RequestPipeline) ProcessRequest(req *http.Request, resp http.ResponseWriter) error {
	ctx := ComponentContext{
		Request:        req,
		ResponseWriter: resp,
	}
	pl(&ctx)
	return ctx.error
}

func CreatePipeline(components ...interface{}) RequestPipeline {
	f := emptyPipeline
	for i := len(components) - 1; i >= 0; i-- {
		currentComponent := components[i]
		_ = services.Populate(currentComponent)
		nextFunc := f
		if servComp, ok := currentComponent.(ServicesMiddlewareComponent); ok {
			f = createServiceDependentFunction(currentComponent, nextFunc)
			servComp.Init()
		} else if stdComp, ok := currentComponent.(MiddlewareComponent); ok {
			f = func(context *ComponentContext) {
				if context.error == nil {
					stdComp.ProcessRequest(context, nextFunc)
				}
			}
			stdComp.Init()
		} else {
			panic("Value is not a middleware component")
		}
	}
	return f
}

func createServiceDependentFunction(component interface{}, nextFunc RequestPipeline) RequestPipeline {
	method := reflect.ValueOf(component).MethodByName("ProcessRequestWithServices")
	if method.IsValid() {
		return func(ctx *ComponentContext) {
			if ctx.error == nil {
				_, err := services.CallForContext(ctx.Request.Context(), method.Interface(), ctx, nextFunc)
				if err != nil {
					ctx.Error(err)
				}
			}
		}
	} else {
		panic("No ProcessRequestWithServices method defined")
	}
}
