package main

import "fmt"

func exercise_2() {
	for i := 65; i < 91; i++ {
		fmt.Println(i)
		for y := 0; y < 3; y++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
