package main

import "fmt"

func exercise_6() {
	message := "Some random message"
	(func(m string) {
		fmt.Println(m)
	})(message)
}
