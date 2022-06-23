package goentities

import (
	"testing"
)

type input struct {
	IntField int
}

type output struct {
	IntField int32 `entity:"IntField"`

	// Fields to be calculated by functions
	FuncField1 int `method:"FuncFieldOne"`
}

// let's take a bussiness logic that FuncField1 is 10X of IntField
func (i output) FuncFieldOne() int {
	i.FuncField1 = int(i.IntField) * 10
	return i.FuncField1
}

func Test_FunctionFields(t *testing.T) {
	testInput := input{
		IntField: -5,
	}

	testOutput := output{}

	outputValue := Present(testInput, testOutput)
	castedOutput := outputValue.(output)

	if castedOutput.IntField != int32(testInput.IntField) {
		t.Errorf("Failed in int casting")
	}

	if castedOutput.FuncField1 != 10*testInput.IntField {
		t.Errorf("Failed in method casting")
	}
}
