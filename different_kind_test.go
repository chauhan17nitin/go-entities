package goentities

import (
	"reflect"
	"testing"
)

func Test_intCasting(t *testing.T) {
	type input struct {
		Test int32
	}

	type output struct {
		Test int64
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

	finalOutput := dummyOutput.Interface()

	if finalOutput.(output).Test != 10 {
		t.Errorf("Failed Int casting without panic")
	}
}

func Test_floatCasting(t *testing.T) {
	type input struct {
		Test float32
	}

	type output struct {
		Test float64
	}

	testInput := input{
		Test: 29.5,
	}

	testOutput := output{}

	inputReflect := reflect.ValueOf(testInput)
	dummyOutput := reflect.New(reflect.Indirect(reflect.ValueOf(testOutput)).Type()).Elem()

	field := dummyOutput.FieldByName("Test")
	value := inputReflect.FieldByName("Test")

	floatCasting(&field, &value)

	finalOutput := dummyOutput.Interface()

	if finalOutput.(output).Test != 29.5 {
		t.Errorf("Failed Int casting without panic")
	}
}
