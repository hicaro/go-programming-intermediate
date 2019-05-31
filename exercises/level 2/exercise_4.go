package main

import "fmt"

func exercise_4() {
	n := 41
	fmt.Printf("%d, %b, %#x\n", n, n, n)

	v := n << 1
	fmt.Printf("%d, %b, %#x\n", v, v, v)
}
