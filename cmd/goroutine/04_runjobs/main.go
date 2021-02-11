package main

import (
	"fmt"
	"time"
)

func worker(i int) {
	fmt.Printf("job %d\n", i)
}

func main() {
	for i := 0; i < 100; i++ {
		go worker(i)
	}

	time.Sleep(time.Second)
}
