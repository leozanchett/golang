package main

import (
	"os"
)

func main() {
	msg := []string{"Rule", " the", " world", " with", " Go!"}
	filePath := "C:/Users/lzandrade/Documents/GitHub/golang/Build-systems-with-go/material/input_output/tmp/msg/msg.txt"
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, s := range msg {
		file.WriteString(s + "\n")
	}
}
