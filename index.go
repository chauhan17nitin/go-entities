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
	Embed Embed  `json:"Embed"`
}

type Embed struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type Test2 struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
	Test  string `json:"Test"`
}

type Foo struct {
	Number int
	Text   string
	fumber int
}

func main() {
	foo := Foo{123, "Hello", 111}
	output := overrideStruct(foo, foo)

	b, err := json.Marshal(output)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))

	return
	t := Test{
		Name:  "Nitin",
		Value: 5,
		Test:  "testing baby",
	}

	f := Test2{}

	// b, err := json.Marshal(t)
	// if err != nil {
	// 	fmt.Printf("Error: %s", err)
	// 	return
	// }
	// fmt.Println(string(b))

	parseStruct(t, f)
	// fmt.Println(reflect.ValueOf(t))
}

func overrideStruct(v interface{}, t interface{}) interface{} {
	// value := reflect.ValueOf(v)
	// val := reflect.Struct(value)
	// numFields := value.Elem().NumField()
	// fmt.Println(numFields)
	old := reflect.ValueOf(t)
	fmt.Println(old.FieldByName("fumber"))
	abc := reflect.New(reflect.Indirect(reflect.ValueOf(v)).Type()).Elem()
	test := old.FieldByName("fumber").Int()
	abc.FieldByName("Number").SetInt(test)
	return abc.Interface()
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
