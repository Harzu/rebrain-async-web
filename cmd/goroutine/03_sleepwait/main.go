package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("job %d\n", i)
		}
	}()

	time.Sleep(time.Second)
}
