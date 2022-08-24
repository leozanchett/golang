package main

import (
	"fmt"
	"time"
)

func MrA(ch chan string) {
	time.Sleep(time.Millisecond * 1200)
	ch <- "Mr. A"
}

func MrB(ch chan string) {
	time.Sleep(time.Millisecond * 500)
	ch <- "Mr. B"
}

func main() {
	//ch := make(chan string)
	ch := make(chan string, 1)
	ch <- "This is main"
	go MrA(ch)
	go MrB(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	time.Sleep(time.Second * 3)
}
