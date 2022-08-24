package main

import (
	"fmt"
	"reflect"
	"strings"
)

type T struct {
	A string
	B int32
	C string
}

func main() {
	t := T{"hello", 42, "bye"}
	fmt.Println(t)

	valueOfT := reflect.ValueOf(&t).Elem()
	for i := 0; i < valueOfT.NumField(); i++ {
		field := valueOfT.Field(i)
		if field.Kind() == reflect.String {
			if field.CanSet() {
				field.SetString(strings.ToUpper(field.String()))
			}
		}
	}
	fmt.Println(t)
}
