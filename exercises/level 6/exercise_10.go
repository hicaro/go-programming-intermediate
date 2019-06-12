package main

import "fmt"

func exercise_10() {
	fn := next_fibonnaci()

	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}

func next_fibonnaci() func() int {
	first := 1
	next := 1

	return func() int {
		value := first

		first = next
		next = first + value

		return value
	}
}
