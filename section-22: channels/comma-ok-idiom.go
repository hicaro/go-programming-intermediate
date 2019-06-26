package main

import "fmt"

func comma_ok_idiom() {
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
		close(q)
	}(even, odd, quit)

	func(e, o, q <-chan int) {
		for {
			select {
			case v := <-e:
				fmt.Println("From the even channel:", v)
			case v := <-o:
				fmt.Println("From the odd channel:", v)
			case v, ok := <-q:
				fmt.Println("From the quit channel:", v, ok)
				return
			}
		}
	}(even, odd, quit)
}
