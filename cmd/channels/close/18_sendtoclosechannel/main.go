package main

func main() {
	channel := make(chan string)
	close(channel)
	channel <- ""
}
