package main

import (
	"context"
	"fmt"
	"time"
)

func work(i int, info chan<- int) {
	t := time.Duration(i*100) * time.Millisecond
	time.Sleep(t)
	fmt.Println("Work done for", i)
	info <- i
}

func main() {
	d := time.Millisecond * 500
	ch := make(chan int)
	i := 0
	for {
		fmt.Println("Value of i:", i)
		ctx, cancel := context.WithTimeout(context.Background(), d)
		go work(i, ch)
		select {
		case x := <-ch:
			fmt.Println("Received", x)
		case <-ctx.Done():
			fmt.Println("Done !!!")
		}
		if ctx.Err() != nil {
			fmt.Println(ctx.Err())
			return
		}
		cancel() // cancel the context
		i++
	}
}
