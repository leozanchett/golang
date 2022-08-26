package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	msg := "Save the world with Go!!!"
	filePath := "C:/Users/lzandrade/Documents/GitHub/golang/Build-systems-with-go/material/input_output/tmp/msg/msg.txt"
	err := ioutil.WriteFile(filePath, []byte(msg), 0644)
	if err != nil {
		panic(err)
	}
	read, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read: %s\n", read)
}
