package main

import "fmt"

func range_channels() {
	c := make(chan int)

	go func(ch chan<- int) {
		for i := 0; i < 100; i++ {
			ch <- i * 2
		}
		close(ch)
	}(c)

	for value := range c {
		fmt.Println(value)
	}
}
