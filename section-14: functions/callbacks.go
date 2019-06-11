package main

import (
	"fmt"
)

func callbacks() {
	numbers := []int{1, 2, 43, 54, 23, 532, 421}

	fmt.Println(sum(numbers...))
	fmt.Println(even(sum, numbers...))
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}

	return total
}

func even(s func(x ...int) int, values ...int) int {
	var even_numbers []int

	// find even numbers
	for _, v := range values {
		if v%2 == 0 {
			even_numbers = append(even_numbers, v)
		}
	}

	return s(even_numbers...)
}
