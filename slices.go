package goentities

import (
	"reflect"
)

func basicSliceValidations(field, value *reflect.Value) {
	if (field.Type().Kind() != reflect.Slice) || (value.Type().Kind() != reflect.Slice) {
		panic("Input and Output Type must be slices")
	}

	if field.Len() != value.Len() || field.Len() <= 0 {
		panic("Slices lengths are unequal")
	}

}

func castSlices(field, value *reflect.Value) {
	if (field.Type().Kind() != reflect.Slice) || (value.Type().Kind() != reflect.Slice) {
		panic("Input and Output Type must be slice")
	}

	if field.Type() == value.Type() {
		field.Set(*value)

		return
	}

	if value.Len() <= 0 {
		return
	}

	var x reflect.Value
	same := false

	if field.Len() != value.Len() {
		x = reflect.MakeSlice(field.Type(), value.Len(), value.Len())
	} else {
		x = *field
		same = true
	}

	sliceType := x.Index(0).Type().Kind()

	if sliceType == reflect.Slice {
		// in future we will do
		return
	}

	if sliceType == reflect.Array {
		// in future to handle
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

	if sliceType == reflect.Struct {
		castStructSlices(&x, value)
	}

	if _, ok := allowedInts[sliceType]; ok {
		CastIntSlices(&x, value)
	}

	if _, ok := allowedUints[sliceType]; ok {
		castUintSlices(&x, value)
	}

	if _, ok := allowedFloats[sliceType]; ok {
		castFloatSlices(&x, value)
	}

	if sliceType == reflect.String {
		castStringSlices(&x, value)
	}

	if sliceType == reflect.Interface {
		CastInterfaceSlice(&x, value)
	}

	if !same {
		field.Set(x)
	}
}

func CastInterfaceSlice(field, value *reflect.Value) {
	basicSliceValidations(field, value)

	if field.Index(0).Type().Kind() != reflect.Interface {
		panic("field type is not slice of interface")
	}

	for i := 0; i < value.Len(); i++ {
		elemValue := value.Index(i)
		elemField := field.Index(i)
		elemField.Set(elemValue)
	}
}

func CastIntSlices(field, value *reflect.Value) {
	basicSliceValidations(field, value)

	if _, ok := allowedInts[field.Index(0).Type().Kind()]; !ok {
		panic("Can not cast non int type to int")
	}

	if _, ok := allowedInts[value.Index(0).Type().Kind()]; !ok {
		panic("Can not cast non int type to int")
	}

	if field.Type() == value.Type() {
		field.Set(*value)

		return
	}

	for i := 0; i < value.Len(); i++ {
		elemValue := value.Index(i)
		elemField := field.Index(i)

		elemField.SetInt(elemValue.Int())
	}
}

func castUintSlices(field, value *reflect.Value) {
	basicSliceValidations(field, value)

	if _, ok := allowedFloats[field.Index(0).Type().Kind()]; !ok {
		panic("Can not cast non uint type to uint")
	}

	if _, ok := allowedFloats[value.Index(0).Type().Kind()]; !ok {
		panic("Can not cast non uint type to uint")
	}

	if field.Type() == value.Type() {
		field.Set(*value)

		return
	}

	for i := 0; i < value.Len(); i++ {
		elemValue := value.Index(i)
		elemField := field.Index(i)

		elemField.SetUint(elemValue.Uint())
	}
}

func castFloatSlices(field, value *reflect.Value) {
	basicSliceValidations(field, value)

	if _, ok := allowedFloats[field.Index(0).Type().Kind()]; !ok {
		panic("Can not cast non float type to float")
	}

	if _, ok := allowedFloats[value.Index(0).Type().Kind()]; !ok {
		panic("Can not cast non float type to float")
	}

	if field.Type() == value.Type() {
		field.Set(*value)

		return
	}

	for i := 0; i < value.Len(); i++ {
		elemValue := value.Index(i)
		elemField := field.Index(i)

		elemField.SetFloat(elemValue.Float())
	}
}

func castStringSlices(field, value *reflect.Value) {
	basicSliceValidations(field, value)

	if field.Index(0).Type().Kind() != reflect.String {
		panic("Can not cast non string type to string")
	}

	if value.Index(0).Type().Kind() != reflect.String {
		panic("Can not cast non string type to string")
	}

	for i := 0; i < value.Len(); i++ {
		elemValue := value.Index(i)
		elemField := field.Index(i)

		elemField.SetString(elemValue.String())
	}
}

func castStructSlices(field, value *reflect.Value) {
	basicSliceValidations(field, value)

	if field.Index(0).Type().Kind() != reflect.Struct {
		panic("Can not cast non string type to string")
	}

	if value.Index(0).Type().Kind() != reflect.Struct {
		panic("Can not cast non string type to string")
	}

	for i := 0; i < value.Len(); i++ {
		elemValue := value.Index(i)
		elemField := field.Index(i)

		castStructs(&elemField, &elemValue)
	}
}
