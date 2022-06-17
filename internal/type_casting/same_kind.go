package typecasting

import (
	"reflect"
)

func CastSameKind(field, value *reflect.Value) {
	if field.Type().Kind() != value.Type().Kind() {
		panic("false call to function")
	}

	if field.Type().Kind() == reflect.Struct {
		// if their types are same then we can cast them directly
		if field.Type() == value.Type() {
			field.Set(*value)

			return
		}
		CastStructs(field, value)
		return
	}

	if field.Type().Kind() == reflect.Slice {
		CastSlices(field, value)
		// we need to find some other way for slice
		return
	}

	if field.Type().Kind() == reflect.Array {
		// some other way for Array
		return
	}

	if field.Type().Kind() == reflect.Interface {
		// let's not go for interfaces for now not sure how it will go
		return
	}

	if field.Type().Kind() == reflect.Map {
		// we need to find some other way for map
		return
	}

	if field.Type().Kind() == reflect.Chan {
		// need some other way for channel
		return
	}

	field.Set(*value)
}

func CastSlices(field, value *reflect.Value) {
	if (field.Type().Kind() != reflect.Slice) || (value.Type().Kind() != reflect.Slice) {
		panic("Input and Output Type must be struct")
	}

	if field.Type() == value.Type() {
		field.Set(*value)

		return
	}

	if value.Len() <= 0 {
		return
	}

	x := reflect.MakeSlice(field.Type(), value.Len(), value.Len())

	sliceType := x.Index(0).Type().Kind()

	if sliceType == reflect.Struct {
		// we will have a different strategy for this in future
		return
	}

	if sliceType == reflect.Slice {
		// in future we will do
		return
	}

	if sliceType == reflect.Array {
		// in future to handle
		return
	}

	if sliceType == reflect.Interface {
		// let's not go for interfaces for now not sure how it will go
		return
	}

	if sliceType == reflect.Map {
		// te be handled in future
		return
	}

	if sliceType == reflect.Chan {
		// te be handled in future
		return
	}

	if _, ok := allowedInts[sliceType]; ok {
		CastIntSlices(&x, value)
	}

	if _, ok := allowedUints[sliceType]; ok {
		CastUintSlices(&x, value)
	}

	if _, ok := allowedFloats[sliceType]; ok {
		CastFloatSlices(&x, value)
	}

	if sliceType == reflect.String {
		CastStringSlices(&x, value)
	}

	field.Set(x)

	return
}

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
