package main

import "fmt"

func exercise_3() {
	c := func() <-chan int {
		c := make(chan int)

		go func() {
			for i := 0; i < 100; i++ {
				c <- i
			}

			close(c)
		}()

		return c
	}()

	func(ch <-chan int) {
		for value := range ch {
			fmt.Println(value)
		}
	}(c)

}
