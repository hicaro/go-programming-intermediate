package main

import "fmt"

func using_channels() {
	c := make(chan int)

	// send
	go send(c)

	// receive
	receive(c)
}

func send(c chan<- int) {
	c <- 401
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
