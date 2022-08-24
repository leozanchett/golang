package main

import (
	"fmt"
	"reflect"
	"strings"
)

type T struct {
	B int
	C string
}

type S struct {
	C string
	D T
	E map[string]int
}

func pointerReflect(offset int, typeOfX reflect.Type) {
	indent := strings.Repeat(" ", offset)
	println(indent, "Type:", typeOfX.Name())
	fmt.Printf("%s %s: %s {\n", indent, typeOfX.Name(), typeOfX.Kind())
	if typeOfX.Kind() == reflect.Struct {
		for i := 0; i < typeOfX.NumField(); i++ {
			field := typeOfX.Field(i)
			pointerReflect(offset+2, field.Type)
		}
	}
}

func main() {
	x := S{"root", T{42, "forty-two"}, make(map[string]int)}
	pointerReflect(0, reflect.TypeOf(x))
}
