package main

import (
	"fmt"
	"reflect"
)

func ValuePrint(i interface{}) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int32:
		fmt.Println("int32 with value:", v.Int())
	case reflect.String:
		fmt.Println("string with value:", v.String())
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	var a int32 = 42
	var b string = "forty-two"
	ValuePrint(a)
	ValuePrint(b)
}
