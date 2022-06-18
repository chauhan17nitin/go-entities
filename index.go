package main

import (
	"encoding/json"
	"fmt"
	typecasting "go-entities/internal/type_casting"
	"reflect"
)

type Test struct {
	Name  string  `json:"Name"`
	Value int     `json:"Value"`
	Test  string  `json:"Test"`
	Slice [][]int `json:"SliceInt"`
	Extra Extra
}

type Test2 struct {
	Name  string        `json:"Name"`
	Value interface{}   `json:"Value"`
	Test  string        `json:"Test"`
	Slice []interface{} `json:"SliceInt"`
	Extra Extra2
}

type Extra struct {
	Test int
}

type Extra2 struct {
	Test int
}

func main() {
	test := Test{
		Name:  "Nitin",
		Value: 5,
		Test:  "fvfvfv",
		Slice: [][]int{{2, 2, 2}, {3, 3, 3}, {4, 4, 4}},
		Extra: Extra{
			Test: 4343,
		},
	}
	output := Present(test, Test2{})
	fmt.Println(output)
	// output := overrideStruct(foo, foo)

	b, err := json.Marshal(output)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))
}

func Present(input interface{}, output interface{}) interface{} {
	// if both of them are not of struct type then you need to panic
	inputValue := reflect.ValueOf(input)
	dummyOutput := reflect.New(reflect.Indirect(reflect.ValueOf(output)).Type()).Elem()

	if (inputValue.Type().Kind() != reflect.Struct) || dummyOutput.Kind() != reflect.Struct {
		panic("Input and output both must be of struct types")
	}

	typecasting.CastStructs(&dummyOutput, &inputValue)

	return dummyOutput.Interface()
}
