package main

func main() {
	channel := make(chan string, 1)
	channel <- "message one"
	channel <- "message two"
}
