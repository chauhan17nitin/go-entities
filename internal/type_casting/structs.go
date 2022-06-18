package typecasting

import "reflect"

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

		if innerField.Type().Kind() == innerValue.Type().Kind() {
			CastSameKind(&innerField, &innerValue)
			continue
		} else {
			// check how to cast of different types if there is any possibility
			CastDifferentKind(&innerField, &innerValue)
		}
	}
}
