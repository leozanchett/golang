package main

import (
	"fmt"
	"time"
)

func generator(ch chan int) {
	sum := 0
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 2)
		sum = sum + 1
	}
	ch <- sum
}

func teste() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("teste")
	}
}

func main() {
	ch := make(chan int)
	go generator(ch)
	go teste()
	fmt.Println("Main waiting for goroutine to finish")
	result := <-ch
	fmt.Println("Main received result:", result)
}
