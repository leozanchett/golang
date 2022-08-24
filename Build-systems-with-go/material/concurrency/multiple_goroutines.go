package main

import (
	"fmt"
	"time"
)

func ShowIt(t time.Duration, num int) {
	fmt.Println("Starting goroutine", num)
	for {
		time.Sleep(t)
		fmt.Println(num)
	}
}

func main() {
	go ShowIt(time.Second, 1)
	go ShowIt(time.Millisecond*500, 500)
	go ShowIt(time.Millisecond*250, 250)
	fmt.Println("Main is done")
	time.Sleep(time.Millisecond * 1200)
}
