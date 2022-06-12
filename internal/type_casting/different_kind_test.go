package typecasting

import (
	"reflect"
	"testing"
)

func Test_intCasting(t *testing.T) {
	type input struct {
		Test int
	}

	type output struct {
		Test int
	}

	testInput := input{
		Test: 10,
	}

	testOutput := output{}

	inputReflect := reflect.ValueOf(testInput)
	dummyOutput := reflect.New(reflect.Indirect(reflect.ValueOf(testOutput)).Type()).Elem()

	field := dummyOutput.FieldByName("Test")
	value := inputReflect.FieldByName("Test")

	intCasting(&field, &value)
}
