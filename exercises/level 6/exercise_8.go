package main

import "fmt"

func exercise_8() {
	fn := factory()

	fmt.Println(fn(3, 2))
}

func factory() func(int, int) int {
	return func(a int, b int) int {
		return a * b
	}
}
