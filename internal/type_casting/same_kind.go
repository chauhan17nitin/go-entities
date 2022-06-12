package typecasting

import "reflect"

func CastSameKind(field, value *reflect.Value) {
	if field.Type().Kind() != value.Type().Kind() {
		panic("false call to function")
	}

	if field.Type().Kind() == reflect.Struct {
		// if it's struct we need to find some other way out
		// now we need to treansform them into one another
		return
	}

	if field.Type().Kind() == reflect.Slice {
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
