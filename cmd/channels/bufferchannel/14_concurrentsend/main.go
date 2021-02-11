package main

import "time"

func main() {
	bufSize := 10
	channel := make(chan int, bufSize)
	for i := 0; i < bufSize; i++ {
		go func(i int) {
			time.Sleep(10 * time.Millisecond)
			channel <- i
		}(i)
	}

	channel <- 1
	time.Sleep(1 * time.Second)
}
