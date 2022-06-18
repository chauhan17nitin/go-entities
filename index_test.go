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
		IntField    int32
		FloatField  float64
		StringField string
		UintField   uint64
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

func Test_PresentStructNesting(t *testing.T) {
	type inputNesting struct {
		IntField    int
		StringField string
	}

	type input struct {
		NestingField inputNesting
	}

	type outputNesting struct {
		IntField    int64
		StringField string
	}

	type output struct {
		NestingField outputNesting
	}

	testInput := input{
		inputNesting{
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

func Test_PresentSliceofStructs(t *testing.T) {
	type input struct {
		IntField    int
		StringField string
	}

	type output struct {
		IntField    int
		StringField string
	}

	testInput := []input{
		{
			IntField:    1,
			StringField: "test string",
		},
		{
			IntField:    2,
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
		if testInput[i].IntField != castedOutput[i].IntField {
			t.Errorf("Failed Casting Array to Struct Int Field")
		}

		if testInput[i].StringField != castedOutput[i].StringField {
			t.Errorf("Failed Casting Array to Struct String Field")
		}
	}
}
