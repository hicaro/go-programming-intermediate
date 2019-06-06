package main

import "fmt"

func exercise_2() {
	arr := []int{2, 3, 22, 43, 54, 64, 73, 101, 121, 200}

	for _, value := range arr {
		fmt.Println(value)
	}

	fmt.Printf("%T\n", arr)
}
