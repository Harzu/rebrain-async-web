package main

import (
	"context"
	"fmt"
)

func worker(ctx context.Context) {
	value, ok := ctx.Value("key").(string)
	if ok {
		fmt.Println(value)
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "key", "value")
	ctx = context.WithValue(ctx, "key2", "value2")

	worker(ctx)
}
