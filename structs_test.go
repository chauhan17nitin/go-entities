package goentities

import (
	"reflect"
	"testing"
)

func Test_CastStructs(t *testing.T) {

	type input struct {
		Field1 int
		Field2 float32
		Field3 string
	}

	type output struct {
		Field1 int32   `entity:"Field1"`
		Field2 float64 `entity:"Field2"`
		Field3 string  `entity:"Field3"`
	}

	testInput := input{
		Field1: 1,
		Field2: 2.5,
		Field3: "test cases",
	}

	testOutput := output{}

	inputReflect := reflect.ValueOf(testInput)
	dummyOutput := reflect.New(reflect.Indirect(reflect.ValueOf(testOutput)).Type()).Elem()

	castStructs(&dummyOutput, &inputReflect)

	finalOutput := dummyOutput.Interface().(output)

	if finalOutput.Field1 != int32(testInput.Field1) {
		t.Errorf("Failed in casting int field")
	}

	if finalOutput.Field2 != float64(testInput.Field2) {
		t.Errorf("Failed in casting float field")
	}

	if finalOutput.Field3 != testInput.Field3 {
		t.Errorf("Failed in casting string field")
	}
}
