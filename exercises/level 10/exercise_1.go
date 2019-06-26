package main

import (
	"fmt"
)

func exercise_1() {
	c := make(chan int, 1)

	// anonymouns func
	// go func() {
	// 	c <- 42
	// }()

	// fmt.Println(<-c)

	// buffered channel
	c <- 42

	fmt.Println(<-c)
}
