package goentities

import (
	"testing"
)

func Test_PresentSimple(t *testing.T) {
	type input struct {
		IntField    int
		FloatField  float32
		StringField string
		UintField   uint
	}

	type output struct {
		IntField    int32   `entity:"IntField"`
		FloatField  float64 `entity:"FloatField"`
		StringField string  `entity:"StringField"`
		UintField   uint64  `entity:"UintField"`
	}

	testInput := input{
		IntField:    -5,
		FloatField:  5.5,
		StringField: "this is a test string",
		UintField:   5,
	}

	testOutput := output{}

	outputValue := Present(testInput, testOutput)
	castedOutput := outputValue.(output)

	if castedOutput.IntField != int32(testInput.IntField) {
		t.Errorf("Failed in casting int field")
	}

	if castedOutput.FloatField != float64(testInput.FloatField) {
		t.Errorf("Failed in Float casting")
	}

	if castedOutput.StringField != testInput.StringField {
		t.Errorf("Failed in String casting")
	}

	if castedOutput.UintField != uint64(testInput.UintField) {
		t.Errorf("Failed in UInt casting")
	}
}

func Test_PresentPointer(t *testing.T) {
	type input struct {
		IntField    int
		FloatField  *float32
		StringField *string
		UintField   *uint
		IntSlices   *[]int
	}

	type output struct {
		IntField    int32   `entity:"IntField"`
		FloatField  float64 `entity:"FloatField"`
		StringField string  `entity:"StringField"`
		UintField   uint64  `entity:"UintField"`
		IntSlices   []int64
	}

	floatValue := float32(5.5)
	stringValue := "this is a test string"
	uintField := uint(5)
	inputSlices := []int{2, 3, 4, 5}

	testInput := input{
		IntField:    -5,
		FloatField:  &floatValue,
		StringField: &stringValue,
		UintField:   &uintField,
		IntSlices:   &inputSlices,
	}

	testOutput := output{}

	outputValue := Present(testInput, testOutput)
	castedOutput := outputValue.(output)

	if castedOutput.IntField != int32(testInput.IntField) {
		t.Errorf("Failed in casting int field")
	}

	if castedOutput.FloatField != float64(*testInput.FloatField) {
		t.Errorf("Failed in Float casting")
	}

	if castedOutput.StringField != *testInput.StringField {
		t.Errorf("Failed in String casting")
	}

	if castedOutput.UintField != uint64(*testInput.UintField) {
		t.Errorf("Failed in UInt casting")
	}

	for i := 0; i < len(castedOutput.IntSlices); i++ {
		if castedOutput.IntSlices[i] != int64((*testInput.IntSlices)[i]) {
			t.Errorf("failed in int slice casting")
		}
	}
}

func Test_PresentPointerStruct(t *testing.T) {
	type input struct {
		IntField    int
		FloatField  *float32
		StringField *string
		UintField   *uint
		IntSlices   *[]int
	}

	type output struct {
		IntField    int32   `entity:"IntField"`
		FloatField  float64 `entity:"FloatField"`
		StringField string  `entity:"StringField"`
		UintField   uint64  `entity:"UintField"`
		IntSlices   []int64
	}

	floatValue := float32(5.5)
	stringValue := "this is a test string"
	uintField := uint(5)
	inputSlices := []int{2, 3, 4, 5}

	testInput := input{
		IntField:    -5,
		FloatField:  &floatValue,
		StringField: &stringValue,
		UintField:   &uintField,
		IntSlices:   &inputSlices,
	}

	testOutput := output{}

	outputValue := Present(&testInput, testOutput)
	castedOutput := outputValue.(output)

	if castedOutput.IntField != int32(testInput.IntField) {
		t.Errorf("Failed in casting int field")
	}

	if castedOutput.FloatField != float64(*testInput.FloatField) {
		t.Errorf("Failed in Float casting")
	}

	if castedOutput.StringField != *testInput.StringField {
		t.Errorf("Failed in String casting")
	}

	if castedOutput.UintField != uint64(*testInput.UintField) {
		t.Errorf("Failed in UInt casting")
	}

	for i := 0; i < len(castedOutput.IntSlices); i++ {
		if castedOutput.IntSlices[i] != int64((*testInput.IntSlices)[i]) {
			t.Errorf("failed in int slice casting")
		}
	}
}

func Test_PresentStructNesting(t *testing.T) {
	type inputNesting struct {
		IntField    int
		StringField string
	}

	type input struct {
		NestingField *inputNesting
	}

	type outputNesting struct {
		IntField    int64  `entity:"IntField"`
		StringField string `entity:"StringField"`
	}

	type output struct {
		NestingField outputNesting `entity:"NestingField"`
	}

	testInput := input{
		&inputNesting{
			IntField:    10,
			StringField: "test string",
		},
	}

	testOutput := output{}

	outputValue := Present(testInput, testOutput)
	castedOutput := outputValue.(output)

	if castedOutput.NestingField.IntField != int64(testInput.NestingField.IntField) {
		t.Errorf("Failed Casting Nested Struct Int Field")
	}

	if castedOutput.NestingField.StringField != testInput.NestingField.StringField {
		t.Errorf("Failed Casting Nested Struct String Field")
	}
}

func Test_PresentMapsInStruct(t *testing.T) {
	type input struct {
		MapField map[int32]int
	}

	type output struct {
		MapField map[int]int `entity:"MapField"`
	}

	testInput := input{
		MapField: map[int32]int{
			2: 2,
			3: 3,
		},
	}

	testOutput := output{}
	outputValue := Present(testInput, testOutput)
	castedOutput := outputValue.(output)

	for key, value := range testInput.MapField {
		if value != castedOutput.MapField[int(key)] {
			t.Errorf("error while casting map fields")
		}
	}
}

func Test_PresentSliceofStructs(t *testing.T) {
	type input struct {
		IntField    *int
		StringField string
	}

	type output struct {
		IntField1    int    `entity:"IntField"`
		StringField1 string `entity:"StringField"`
	}

	intValue1 := 1
	intValue2 := 2

	testInput := []input{
		{
			IntField:    &intValue1,
			StringField: "test string",
		},
		{
			IntField:    &intValue2,
			StringField: "test string 2",
		},
	}

	testOutput := output{}

	outputValue := Present(testInput, testOutput)
	castedOutput := outputValue.([]output)

	if len(testInput) != len(castedOutput) {
		t.Errorf("Unmacthed output and input length")
	}

	for i := 0; i < len(castedOutput); i++ {
		if *testInput[i].IntField != castedOutput[i].IntField1 {
			t.Errorf("Failed Casting Array to Struct Int Field")
		}

		if testInput[i].StringField != castedOutput[i].StringField1 {
			t.Errorf("Failed Casting Array to Struct String Field")
		}
	}
}
