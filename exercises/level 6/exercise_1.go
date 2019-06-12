package main

import "fmt"

func exercise_1() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 321
}

func bar() (int, string) {
	return 532, "Hello, World!"
}
