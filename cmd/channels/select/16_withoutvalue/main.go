package main

import (
	"fmt"
)

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)

	select {
	case msg := <-firstChan:
		fmt.Println(msg)
	case msg := <-secondChan:
		fmt.Println(msg)
	}
}
