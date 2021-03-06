package goentities

import (
	"reflect"
)

var allowedInts = map[reflect.Kind]struct{}{
	reflect.Int:   {},
	reflect.Int8:  {},
	reflect.Int16: {},
	reflect.Int32: {},
	reflect.Int64: {},
}

var allowedFloats = map[reflect.Kind]struct{}{
	reflect.Float32: {},
	reflect.Float64: {},
}

var allowedUints = map[reflect.Kind]struct{}{
	reflect.Uint:   {},
	reflect.Uint8:  {},
	reflect.Uint16: {},
	reflect.Uint32: {},
	reflect.Uint64: {},
}

func castDifferentKind(field, value *reflect.Value) {

	if field.Type().Kind() == reflect.Ptr {
		pointerFieldCasting(field, value)
		return
	}

	if value.Type().Kind() == reflect.Ptr {
		pointerValueCasting(field, value)
		return
	}

	if field.Type().Kind() == value.Type().Kind() {
		castField(field, value)
		return
	}

	if _, ok := allowedInts[field.Type().Kind()]; ok {
		intCasting(field, value)
		return
	}

	if _, ok := allowedFloats[field.Type().Kind()]; ok {
		floatCasting(field, value)
		return
	}

	if _, ok := allowedUints[field.Type().Kind()]; ok {
		uintCasting(field, value)
		return
	}

	if field.Type().Kind() == reflect.Interface {
		field.Set(*value)
		return
	}

	// here i think we can panic because the person is
	// trying to cast very different kind of fields all together
}

func intCasting(field, value *reflect.Value) {
	if _, ok := allowedInts[field.Type().Kind()]; !ok {
		panic("Can not cast non int type to int")
	}

	if _, ok := allowedInts[value.Type().Kind()]; !ok {
		panic("Can not cast non int type to int")
	}

	field.SetInt(value.Int())
}

func uintCasting(field, value *reflect.Value) {
	if _, ok := allowedUints[field.Type().Kind()]; !ok {
		panic("Can not cast non float type to float")
	}

	if _, ok := allowedUints[field.Type().Kind()]; !ok {
		panic("Can not cast non float type to float")
	}

	field.SetUint(value.Uint())
}

func floatCasting(field, value *reflect.Value) {
	if _, ok := allowedFloats[field.Type().Kind()]; !ok {
		panic("Can not cast non float type to float")
	}

	if _, ok := allowedFloats[field.Type().Kind()]; !ok {
		panic("Can not cast non float type to float")
	}

	field.SetFloat(value.Float())
}
