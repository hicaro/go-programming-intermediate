package main

import "fmt"

func exercise_4() {
	year := 1987

	for {
		fmt.Println(year)
		year++
		if year > 2019 {
			break
		}
	}
}
