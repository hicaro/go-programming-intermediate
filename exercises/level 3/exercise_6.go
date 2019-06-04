package main

import "fmt"

func exercise_6() {
	for i := 10; i < 100; i++ {
		if (i % 4) == 0 {
			fmt.Println(i, " is divisible by 4")
		}
	}
}
