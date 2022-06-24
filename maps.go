package goentities

import (
	"fmt"
	"reflect"
)

func castMap(field, value *reflect.Value) {
	if value.Len() == 0 {
		// no benefit in casting 0 length map
		return
	}

	if field.Type() == value.Type() {
		field.Set(*value)
	}

	// fmt.Println(field.MapKeys())
	fmt.Println(field.Type().Key().Kind())
	fmt.Println(value.Type().Elem().Kind())

	if _, ok := allowedInts[field.Type().Key().Kind()]; ok {
		castIntKeyMap(field, value)
		return
	}

	// reflect.MakeMap()
	// elem := value.Index(0)
	// fmt.Println(value.MapKeys())
	// for _, key := range value.MapKeys() {
	// 	keyVal := value.MapIndex(key)
	// 	fmt.Println(keyVal)
	// }

}

func castIntKeyMap(field, value *reflect.Value) {
	if _, ok := allowedInts[field.Type().Key().Kind()]; !ok {
		panic("can no cast non int type map to int")
	}

	if _, ok := allowedInts[value.Type().Key().Kind()]; !ok {
		panic("can no cast non int type map to int")
	}

	fmt.Println("came here")

	return
}

func castFloatKeyMap(field, value *reflect.Value) {

}

func castStringKeyMap(field, value *reflect.Value) {

}
