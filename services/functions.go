package services

import (
	"context"
	"errors"
	"reflect"
)

func Call(target interface{}, otherArgs ...interface{}) ([]interface{}, error) {
	return CallForContext(context.Background(), target, otherArgs...)
}

func CallForContext(c context.Context, target interface{}, otherArgs ...interface{}) (results []interface{}, err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Func {
		resultValues := invokeFunction(c, targetValue, otherArgs...)
		results = make([]interface{}, len(resultValues))
		for i := 0; i < len(resultValues); i++ {
			results[i] = resultValues[i].Interface()
		}
	} else {
		err = errors.New("only functions can be invoked")
	}
	return
}
