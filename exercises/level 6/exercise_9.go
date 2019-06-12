package main

import "fmt"

func exercise_9() {
	sum_callback(34, 22, func(s int) {
		fmt.Println(s)
	})
}

func sum_callback(a int, b int, cb func(int)) {
	s := a + b
	cb(s)
}
