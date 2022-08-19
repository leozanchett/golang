package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Create("dat2")
	check(err)

	defer f.Close()
	n1, err := f.WriteString("alo tsom teas dtes aioadsokdas 1231\n ")
	fmt.Printf("wrote %d bytes\n", n1)
	check(err)

	n2, err := f.WriteString("xau ")
	fmt.Printf("wrote %d bytes\n", n2)
	check(err)
}
