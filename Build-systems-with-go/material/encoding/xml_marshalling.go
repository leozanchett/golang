package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	number, err := xml.Marshal(42)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(number))

	float, _ := xml.Marshal(3.14)
	fmt.Println(string(float))

	msg, _ := xml.Marshal("Hello World")
	fmt.Println(string(msg))

	numbers, _ := xml.Marshal([]int{1, 1, 2, 3, 5, 8})
	fmt.Println(string(numbers))

	aMap, err := xml.Marshal(map[string]int{"one": 1, "two": 2})
	fmt.Println(err)
	fmt.Println("-", string(aMap), "-")
}
