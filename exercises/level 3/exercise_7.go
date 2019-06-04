package main

import "fmt"

func exercise_7() {
	for i := 10; i < 100; i++ {
		if (i % 4) == 0 {
			fmt.Println(i, " is divisible by 4")
		} else if (i % 71) == 0 {
			fmt.Println(i, " is divisible by 71")
		} else {
			fmt.Println(i)
		}
	}
}
