package main

import (
	"example.com/greetings"
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(greetings.Hello("Leo"))
}
