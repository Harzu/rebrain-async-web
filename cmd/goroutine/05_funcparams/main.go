package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("job %d\n", i)
		}()
	}

	time.Sleep(time.Second)
}
