package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("sent all jobs")
	jobs <- 4
	close(jobs)
	time.Sleep(2 * time.Second)

	<-done
}
