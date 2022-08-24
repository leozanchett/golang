package main

import (
	"fmt"
	"time"
)

func receiver(input <-chan int) {
	for i := range input {
		fmt.Println(i)
	}
}

func sender(output chan<- int, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("Sending", i)
		output <- i
	}
	close(output)
}

func main() {
	ch := make(chan int)
	go sender(ch, 5)
	receiver(ch)
	fmt.Println("Done")

}
