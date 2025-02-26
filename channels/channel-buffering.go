package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	messages2 := make(chan string, 2)

	messages2 <- "buffered"
	fmt.Println(<-messages2)
	messages2 <- "channel"
	fmt.Println(<-messages2)

	/*
		messages3 := make(chan string, 2)

		Trying to use the value from an empty channel will give error
		fmt.Println(<-messages3)
		messages3 <- "buffered"
		messages3 <- "channel"
		fmt.Println(<-messages3)
	*/
}
