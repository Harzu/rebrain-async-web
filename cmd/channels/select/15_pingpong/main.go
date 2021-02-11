package main

import (
	"fmt"
	"time"
)

func player(
	name string,
	redSideChan chan int,
	blueSideChan chan int,
	gameOverChan <-chan struct{},
) {
	for {
		select {
		case <-gameOverChan:
			fmt.Printf("stop %s goroutine\n", name)
			return
		case count := <-redSideChan:
			fmt.Printf("red player side count: %d\n", count)
			count++
			time.Sleep(100 * time.Millisecond)
			blueSideChan <- count
		case count := <-blueSideChan:
			fmt.Printf("blue player side count: %d\n", count)
			count++
			time.Sleep(100 * time.Millisecond)
			redSideChan <- count
		}
	}
}

func main() {
	redSideChan := make(chan int, 1)
	blueSideChan := make(chan int, 1)
	gameOverChan := make(chan struct{})

	go player("red player", redSideChan, blueSideChan, gameOverChan)
	go player("blue player", redSideChan, blueSideChan, gameOverChan)

	redSideChan <- 0
	time.Sleep(time.Second)
	gameOverChan <- struct{}{}
	gameOverChan <- struct{}{}
}
