package goentities

import (
	"reflect"
	"strings"
)

func castStructs(field, value *reflect.Value) {
	if (field.Type().Kind() != reflect.Struct) || (value.Type().Kind() != reflect.Struct) {
		panic("Input and Output Type must be struct")
	}

	outputType := field.Type()

	for i := 0; i < field.NumField(); i++ {
		defer fieldFuncCasting(field, i)
		innerField := field.Field(i)
		// in future we can add multiple key support here
		// comma separated multiple fields whichever we found first in the other struct use that
		key := strings.Split(outputType.Field(i).Tag.Get("entity"), ",")[0]
		if key == "" {
			continue
		}
		innerValue := value.FieldByName(key)

		// If the Field is not present in input then continue
		if innerValue.Kind() == reflect.Invalid {
			continue
		}

		castField(&innerField, &innerValue)
	}
}

func fieldFuncCasting(structField *reflect.Value, index int) {
	innerField := structField.Field(index)
	if !innerField.IsValid() {
		return
	}

	outputType := structField.Type()

	methodName := strings.Split(outputType.Field(index).Tag.Get("method"), ",")[0]
	if methodName == "" {
		return
	}

	method := structField.MethodByName(methodName)
	if !method.IsValid() {
		// method not present invalid method
		return
	}

	methodOutput := method.Call([]reflect.Value{})
	if len(methodOutput) == 0 {
		return
	}

	innerField.Set(methodOutput[0])
}

func castSliceofStructs(field, value *reflect.Value) interface{} {
	if field.Type().Kind() != reflect.Struct || value.Type().Kind() != reflect.Slice {
		panic("Unsupported Input and output types")
	}

	outputType := reflect.SliceOf(field.Type())
	output := reflect.MakeSlice(outputType, value.Len(), value.Len())

	castSlices(&output, value)

	return output.Interface()
}
