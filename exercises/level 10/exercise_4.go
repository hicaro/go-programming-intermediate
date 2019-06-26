package main

import "fmt"

func exercise_4() {
	q := make(chan int)
	c := func(q chan<- int) <-chan int {
		c := make(chan int)

		go func() {
			for i := 0; i < 100; i++ {
				c <- i
				q <- i * 2
			}
		}()

		return c
	}(q)

	func(a, b <-chan int) {
		for {
			select {
			case value := <-a:
				fmt.Println("Coming from a", value)
			case value := <-b:
				fmt.Println("Coming from b", value)
				if value >= 198 {
					return
				}
			}
		}
	}(c, q)

	fmt.Println("about to exit")
}
