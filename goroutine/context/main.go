package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "run sub goroutine")

	time.Sleep(10 * time.Second)
	fmt.Println("stop sub goroutine")
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			fmt.Println(name, "sub goroutine stop")
			return
		default:
			fmt.Println(name, "sub goroutine running")
			time.Sleep(3 * time.Second)
		}
	}
}
