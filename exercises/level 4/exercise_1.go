package main

import "fmt"

func exercise_1() {
	arr := [5]int{1, 2, 3, 4, 5}

	for _, value := range arr {
		fmt.Println(value)
	}

	fmt.Printf("%T\n", arr)
}
