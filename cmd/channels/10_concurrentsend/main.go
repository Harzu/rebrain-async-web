package main

import "time"

func main() {
	channel := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(10 * time.Millisecond)
			channel <- i
		}(i)
	}

	channel <- 1
	time.Sleep(time.Second)
}
