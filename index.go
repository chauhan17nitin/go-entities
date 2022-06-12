package main

import (
	"encoding/json"
	"fmt"
	typecasting "go-entities/internal/type_casting"
	"reflect"
)

type Test struct {
	Name  string `json:"Name"`
	Value int    `json:"Value"`
	Test  string `json:"Test"`
	Extra Extra
}

type Test2 struct {
	Name  string `json:"Name"`
	Value int64  `json:"Value"`
	Test  string `json:"Test"`
	Extra Extra
}

type Extra struct {
	test int
}

type Extra2 struct {
}

func main() {
	test := Test{
		Name:  "Nitin",
		Value: 5,
		Test:  "fvfvfv",
		Extra: Extra{
			test: 4343,
		},
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

		// If the Field is not present in input then continue
		if value.Kind() == reflect.Invalid {
			continue
		}

		if field.Type() == value.Type() {
			field.Set(value)
			continue
		}

		if field.Type().Kind() == value.Type().Kind() {
			// most probably they are structs
			typecasting.CastSameKind(&field, &value)
			continue
		} else {
			// check how to cast of different types if there is any possibility
		}
	}

	return dummyOutput.Interface()
}
