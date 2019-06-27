package main

import (
	"fmt"
	"os"
)

func panic_func() {
	defer func() {
		fmt.Println("The program panicked, but I still run.")
	}()
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
}
