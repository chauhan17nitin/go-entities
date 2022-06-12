package typecasting

import "reflect"

var allowedInts = map[reflect.Kind]struct{}{
	reflect.Int:   {},
	reflect.Int8:  {},
	reflect.Int16: {},
	reflect.Int32: {},
	reflect.Int64: {},
}

var allowedIntCasts = map[reflect.Kind]struct{}{
	reflect.Int:   {},
	reflect.Int8:  {},
	reflect.Int16: {},
	reflect.Int32: {},
	reflect.Int64: {},

	reflect.Float32: {},
	reflect.Float64: {},
}

func CastDifferentKind(field, value *reflect.Value) {
	if field.Type().Kind() == value.Type().Kind() {
		CastSameKind(field, value)
		return
	}

	if _, ok := allowedInts[field.Type().Kind()]; ok {

	}

}

func intCasting(field, value *reflect.Value) {
	if _, ok := allowedInts[field.Type().Kind()]; !ok {
		panic("Can not cast non int type to int")
	}

	if _, ok := allowedIntCasts[value.Type().Kind()]; !ok {
		panic("Can not cast non int type to int")
	}

	if field.Type().Kind() == value.Type().Kind() {
		field.Set(*value)
	}
}
