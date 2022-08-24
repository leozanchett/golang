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
	close(out)
}

func sendMsgs(out chan<- string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 250)
		out <- strconv.Itoa(i)
	}
	fmt.Println("no more messages")
	close(out)
}

func main() {
	numbers := make(chan int)
	msgs := make(chan string)

	go sendNumbers(numbers)
	go sendMsgs(msgs)

	notclosedNums, notclosedMsgs := true, true

	for notclosedNums || notclosedMsgs {
		select {
		case num, ok := <-numbers:
			if notclosedNums = ok; notclosedNums {
				fmt.Println("number:", num)
			}
		case msg, ok := <-msgs:
			if notclosedMsgs = ok; notclosedMsgs {
				fmt.Println("message:", msg)
			}
		}
	}
}
