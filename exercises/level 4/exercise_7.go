package main

import "fmt"

func exercise_7() {
	records := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for _, data := range records {
		for _, value := range data {
			fmt.Println(value)
		}
	}
}
