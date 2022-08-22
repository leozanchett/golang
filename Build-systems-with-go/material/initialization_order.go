package main

import "fmt"

var x = xSetter()

func xSetter() int {
	fmt.Println("xSetter")
	return 42
}

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("This is the main")
}
