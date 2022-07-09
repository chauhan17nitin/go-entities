package goentities

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	tests := []struct {
		name       string
		input      input
		outputType output
		want       output
	}{
		{
			name: "test 1",
			input: input{
				IntField: -5,
			},
			outputType: output{},
			want: output{
				IntField:   -5,
				FuncField1: -50,
			},
		},
		{
			name: "test 2",
			input: input{
				IntField: 0,
			},
			outputType: output{},
			want: output{
				IntField:   0,
				FuncField1: 0,
			},
		},
		{
			name: "test 3",
			input: input{
				IntField: 11,
			},
			outputType: output{},
			want: output{
				IntField:   11,
				FuncField1: 110,
			},
		},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Present(tt.input, tt.outputType).(output)

			assert.Equal(t, tt.want, got, "ouput does not matches")
		})
	}
}
