package main

import "fmt"

func anonymous() {
	(func() {
		fmt.Println("Hello from anonymous function")
	})()

	(func(x int) {
		fmt.Println("Hello from anonymous function with", x)
	})(50)
}
