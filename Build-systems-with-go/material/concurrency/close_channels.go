package main

import (
	"fmt"
	"time"
)

func sender(out chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		out <- i
	}
	close(out)
	fmt.Println("Sender finished")
}

func main() {
	ch := make(chan int)

	go sender(ch)
	// exemplo utilizando o for

	// for {
	// 	num, found := <-ch
	// 	if found {
	// 		fmt.Println(num)
	// 	} else {
	// 		fmt.Println("Channel closed")
	// 		break
	// 	}
	// }

	// exemplo utilizando o range
	for num := range ch {
		fmt.Println(num)
	}
	fmt.Println("Channel closed")
}
