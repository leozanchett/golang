package main

import (
	"context"
	"fmt"
	"time"
)

func calc(ctx context.Context) {
	switch ctx.Value("action") {
	case "quick":
		fmt.Println("Quick calculation")
	case "slow":
		time.Sleep(time.Second * 2)
		fmt.Println("Slow calculation")
	case <-ctx.Done():
		fmt.Println("Done!!!")
	default:
		panic("Unknown action")
	}
}

func main() {
	t := time.Millisecond * 1900
	ctx, cancel := context.WithTimeout(context.Background(), t)
	qCtx := context.WithValue(ctx, "action", "quick")
	defer cancel()
	fmt.Println("go quick")
	go calc(qCtx)
	<-qCtx.Done()

	ctx2, cancel2 := context.WithTimeout(context.Background(), t)
	sCtx := context.WithValue(ctx2, "action", "slow")
	defer cancel2()
	fmt.Println("go slow")
	go calc(sCtx)
	<-sCtx.Done()
	fmt.Println("Finished")
}
