package main

import "fmt"

func exercise_5() {
	c := make(chan int)

	go func() {
		c <- 43
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
