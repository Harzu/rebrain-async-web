package main

import (
	"fmt"
	"time"
)

func sender(channel chan<- string) {
	channel <- "Hello"
}

func receiver(channel <-chan string) {
	msg := <- channel
	fmt.Printf("%s world\n", msg)
}

func main() {
	channel := make(chan string)

	go sender(channel)
	go receiver(channel)

	time.Sleep(time.Second)
}
