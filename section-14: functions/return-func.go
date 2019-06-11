package main

import "fmt"

func return_func() {
	f := zoo()

	fmt.Println(f())
}

func zoo() func() int {
	return func() int {
		return 541
	}
}
