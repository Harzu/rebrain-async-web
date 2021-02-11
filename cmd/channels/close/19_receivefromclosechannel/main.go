package main

import "fmt"

func main() {
	channel := make(chan string, 1)
	channel <- "hello"
	// close(channel)

	msg := <-channel
	fmt.Println(msg)

	msg = <-channel
	fmt.Println(msg)
}
