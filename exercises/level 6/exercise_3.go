package main

import "fmt"

func exercise_3() {
	defer df()

	fmt.Println("I am second")
}

func df() {
	fmt.Println("I am first")
}
