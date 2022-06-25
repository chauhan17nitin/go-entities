package goentities

import (
	"reflect"
)

func castMap(field, value *reflect.Value) {
	if value.Type().Kind() != reflect.Map {
		panic("can not cast non map field to map")
	}

	if value.Len() == 0 {
		// no benefit in casting 0 length map
		return
	}

	if field.Type() == value.Type() {
		field.Set(*value)
	}

	fieldKey := reflect.New(field.Type().Key()).Elem()
	fieldKeyValue := reflect.New(field.Type().Elem()).Elem()

	dummyMap := reflect.MakeMap(field.Type())

	for _, key := range value.MapKeys() {
		keyVal := value.MapIndex(key)
		castField(&fieldKeyValue, &keyVal)
		castField(&fieldKey, &key)
		dummyMap.SetMapIndex(fieldKey, fieldKeyValue)
	}

	field.Set(dummyMap)
}
