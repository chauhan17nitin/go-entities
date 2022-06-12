package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Test struct {
	Name  string `json:"Name"`
	Value int    `json:"Value"`
	Test  string `json:"Test"`
}

type Test2 struct {
	Name  string `json:"Name"`
	Value int    `json:"Value"`
	Test  string `json:"Test"`
}

func main() {
	test := Test{
		Name:  "Nitin",
		Value: 5,
		Test:  "fvfvfv",
	}
	output := overrideStructFinal(test, Test2{})
	fmt.Println(output)
	// output := overrideStruct(foo, foo)

	b, err := json.Marshal(output)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))

	return
}

func overrideStructFinal(input interface{}, output interface{}) interface{} {
	// if both of them are not of struct type then you need to panic
	inputValue := reflect.ValueOf(input)
	dummyOutput := reflect.New(reflect.Indirect(reflect.ValueOf(output)).Type()).Elem()
	outputType := reflect.ValueOf(output).Type()

	if (inputValue.Type().Kind() != reflect.Struct) || dummyOutput.Kind() != reflect.Struct {
		panic("Input and output both must be of struct types")
	}

	for i := 0; i < dummyOutput.NumField(); i++ {
		field := dummyOutput.Field(i)
		value := inputValue.FieldByName(outputType.Field(i).Name)

		field.Set(value)
	}

	return dummyOutput.Interface()
}
