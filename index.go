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

type Embed struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type Foo struct {
	Number int
	Text   string
	fumber int
}

func main() {
	// foo := Foo{123, "Hello", 111}
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

func overrideStruct(v interface{}, t interface{}) interface{} {
	old := reflect.ValueOf(t)
	fmt.Println(old.FieldByName("fumber"))
	abc := reflect.New(reflect.Indirect(reflect.ValueOf(v)).Type()).Elem()
	fmt.Println(abc.Kind())
	test := old.FieldByName("fumber").Int()
	abc.FieldByName("Number").SetInt(test)
	return abc.Interface()
}

func overrideStructFinal(input interface{}, output interface{}) interface{} {
	// if both of them are not of struct type then you need to panic
	inputValue := reflect.ValueOf(input)
	dummyOutput := reflect.New(reflect.Indirect(reflect.ValueOf(output)).Type()).Elem()
	outputType := reflect.ValueOf(output).Type()

	if (inputValue.Type().Kind() != reflect.Struct) || dummyOutput.Kind() != reflect.Struct {
		panic("Input and output both must be of struct types")
	}

	// test := dummyOutput.Field(1)
	// test.Set(4)

	for i := 0; i < dummyOutput.NumField(); i++ {
		field := dummyOutput.Field(i)
		fmt.Println(outputType.Field(i))
		value := inputValue.FieldByName(outputType.Field(i).Name)

		field.Set(value)
		// fmt.Println()
	}

	return dummyOutput.Interface()
}

func parseStruct(input interface{}, output interface{}) {
	typeInput := reflect.ValueOf(input)
	typeOutput := reflect.ValueOf(output).Type()
	fmt.Println(typeInput.Kind())
	fmt.Println(typeOutput.Kind())
	// fmt.Println(typeInput.)
	value := reflect.ValueOf(input)

	if value.Kind() == reflect.Struct {
		f := value.FieldByName("Name")
		fmt.Println("printing the fields")
		if f.IsValid() {
			fmt.Println("fucking yes")
			if f.CanSet() {
				fmt.Println("yes can be changed")
			}
		} else {
			fmt.Println("damn no")
		}
	}
	fmt.Println(value.Kind())

	// fmt.Println(reflect.ValueOf(input).FieldByName("Name").CanSet())
	// // fmt.Println(reflect.ValueOf(output).Set(value))
	// fmt.Println(reflect.ValueOf(output).FieldByName("Name").CanSet())
	// fmt.Println(typeInput.Field(0))
	// fmt.Println(typeInput.Field(1))
	// fmt.Println(typeInput.Field(3))

}
