package typecasting

import "reflect"

func basicSliceValidations(field, value *reflect.Value) {
	if (field.Type().Kind() != reflect.Slice) || (value.Type().Kind() != reflect.Slice) {
		panic("Input and Output Type must be slices")
	}

	if field.Len() != value.Len() || field.Len() <= 0 {
		panic("Slices lengths are unequal")
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

	return
}

func CastUintSlices(field, value *reflect.Value) {
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

	return
}

func CastFloatSlices(field, value *reflect.Value) {
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

	return
}

func CastStringSlices(field, value *reflect.Value) {
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

	return
}
