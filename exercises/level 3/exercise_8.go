package main

import "fmt"

func exercise_8() {
	i := 31
	switch {
	case (i % 2) == 0:
		{
			fmt.Println("It is an even number")
		}
	default:
		{
			fmt.Println("It is not an even number")
		}
	}
}
