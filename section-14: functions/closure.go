package main

import "fmt"

func closure() {
	fn := incrementor()
	fna := incrementor()

	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	fmt.Println(fna())
	fmt.Println(fna())
	fmt.Println(fna())
}

func incrementor() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}
