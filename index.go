package goentities

import (
	"reflect"
)

func Present(input interface{}, output interface{}) interface{} {
	// if both of them are not of struct type then you need to panic
	inputValue := reflect.ValueOf(input)
	dummyOutput := reflect.New(reflect.Indirect(reflect.ValueOf(output)).Type()).Elem()

	if (inputValue.Type().Kind() != reflect.Struct) || dummyOutput.Kind() != reflect.Struct {
		if inputValue.Type().Kind() == reflect.Slice && dummyOutput.Kind() == reflect.Struct {
			return castSliceofStructs(&dummyOutput, &inputValue)
		}
		panic("invalid input and output formats")
	}

	castStructs(&dummyOutput, &inputValue)

	return dummyOutput.Interface()
}
