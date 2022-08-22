package main

import (
	"fmt"
	_ "material/package/a"
)

func init() {
	fmt.Println("Init from my program")
}

func main() {
	fmt.Println("My program")
}
