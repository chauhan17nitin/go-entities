package typecasting

import (
	"reflect"
)

func CastField(field, value *reflect.Value) {
	if field.Type().Kind() != value.Type().Kind() {
		CastDifferentKind(field, value)

		return
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
		return
	}

	if field.Type().Kind() == reflect.Array {
		// some other way for Array
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
