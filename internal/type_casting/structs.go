package typecasting

import (
	"reflect"
)

func CastStructs(field, value *reflect.Value) {
	if (field.Type().Kind() != reflect.Struct) || (value.Type().Kind() != reflect.Struct) {
		panic("Input and Output Type must be struct")
	}

	outputType := field.Type()

	for i := 0; i < field.NumField(); i++ {
		innerField := field.Field(i)
		innerValue := value.FieldByName(outputType.Field(i).Name)

		// If the Field is not present in input then continue
		if innerValue.Kind() == reflect.Invalid {
			continue
		}

		CastField(&innerField, &innerValue)
	}
}

func CastSliceofStructs(field, value *reflect.Value) interface{} {
	if field.Type().Kind() != reflect.Struct || value.Type().Kind() != reflect.Slice {
		panic("Unsupported Input and output types")
	}

	outputType := reflect.SliceOf(field.Type())
	output := reflect.MakeSlice(outputType, value.Len(), value.Len())

	// fmt.Println(output)
	CastSlices(&output, value)

	return output.Interface()
}
