package main

import "fmt"

// Important name_channel chan<- type_value, chan<- means that the channel is
// for sending values only
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Important name_channel <-chan type_value, <-chan means that the channel is
// for reveiving values
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// For sending		chan<-
// For receiving	<-chan

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
