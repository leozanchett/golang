package main

import (
	"fmt"
	"time"
)

func worker(x *int) {
	for {
		time.Sleep(time.Millisecond * 500)
		*x++
	}
}

func main() {
	timer := time.NewTimer(time.Second * 3)
	ticker := time.NewTicker(time.Second)

	x := 0
	go worker(&x)
	for {
		select {
		case <-timer.C:
			fmt.Printf("timer -> %d\n", x)
		case <-ticker.C:
			fmt.Printf("ticker -> %d\n", x)
		}
		if x > 10 {
			break
		}
	}
}
