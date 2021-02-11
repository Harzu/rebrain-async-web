package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancel")
			return
		default:
		}
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	go worker(ctx)

	time.Sleep(2 * time.Second)
}
