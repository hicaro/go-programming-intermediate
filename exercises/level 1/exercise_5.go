package main

import "fmt"

func exercise_5() {
	var x myType = 43
	exercise_4(x)

	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
