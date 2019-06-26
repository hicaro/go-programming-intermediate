package main

import "fmt"

func select_channels() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go func(e, o, q chan<- int) {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				e <- i
			} else {
				o <- i
			}
		}
		q <- 0
	}(even, odd, quit)

	func(e, o, q <-chan int) {
		for {
			select {
			case v := <-e:
				fmt.Println("From the even channel:", v)
			case v := <-o:
				fmt.Println("From the odd channel:", v)
			case v := <-q:
				fmt.Println("From the quit channel:", v)
				return
			}
		}
	}(even, odd, quit)
}
