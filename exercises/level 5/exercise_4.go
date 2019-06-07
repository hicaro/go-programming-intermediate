package main

import "fmt"

func exercise_4() {
	a := struct {
		species string
		age     int
	}{
		species: "Cat",
		age:     9,
	}

	fmt.Println(a)
}
