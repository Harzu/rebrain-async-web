package main

func main() {
	channel := make(chan string)
	<-channel
}
