package goentities

import "reflect"

func pointerValueCasting(field, value *reflect.Value) {
	if field.Type().Kind() == reflect.Ptr {
		pointerBothCasting(field, value)
	}

	derefValue := value.Elem()
	castField(field, &derefValue)
}

func pointerBothCasting(field, value *reflect.Value) {

}

func pointerFieldCasting(field, value *reflect.Value) {
	// field.SetPointer(unsafe.Pointer(value.UnsafeAddr()))

	// newField := reflect.ValueOf(field.Elem()).Elem()

	// newField := field.Elem()
	// field.Set(reflect.ValueOf((value.Interface())))
	// ptr := reflect.ValueOf(field.Interface())
	// fmt.Println(ptr)
	// fmt.Println(ptr.Type().Name())
	// fmt.Println(ptr.Type().Kind())
	// // fmt.Println(newField.Type())
	// fmt.Println(field.Type())
	// // fmt.Println(newField.Kind())
	// fmt.Println(value.Int())
	// // castField(&newField, value)
	// // field.Set(newField)
	// fmt.Println("came in field pointer casting")
}
