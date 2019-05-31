package main

import "fmt"

type myType int

func exercise_4(x myType) {

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)
}
