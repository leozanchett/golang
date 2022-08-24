package main

import (
	"fmt"
	"strconv"
	"time"
)

func sendNumbers(out chan<- int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		out <- i
	}
	fmt.Println("no more numbers")
}

func sendMsgs(out chan<- string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 250)
		out <- strconv.Itoa(i)
	}
	fmt.Println("no more messages")
}

func main() {
	numbers := make(chan int)
	msgs := make(chan string)

	go sendNumbers(numbers)
	go sendMsgs(msgs)
	for i := 0; i < 10; i++ {
		//for {
		select {
		case num := <-numbers:
			fmt.Println("number:", num)
		case msg := <-msgs:
			fmt.Println("message:", msg)
		}
	}
}
