package services

import (
	"context"
	"errors"
	"reflect"
)

func GetServiceForContext(c context.Context, target interface{}) (err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr && targetValue.Elem().CanSet() {
		err = resolveServiceFromValue(c, targetValue)
	} else {
		err = errors.New("type cannot be used as target")
	}
	return
}

func GetService(target interface{}) error {
	return GetServiceForContext(context.Background(), target)
}
