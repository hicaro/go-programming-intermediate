package main

import "fmt"

func exercise_2() {
	x := []int{23, 23, 53, 11, 2, 76, 3}

	fmt.Println(sum(x...))
	fmt.Println(total(x))
}

func sum(x ...int) int {
	var t int

	for _, v := range x {
		t += v
	}

	return t
}

func total(x []int) int {
	var t int

	for _, v := range x {
		t += v
	}

	return t
}
