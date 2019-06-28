package main

import (
	"fmt"
	"go-programming/exercises/level_12/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(7),
	}

	fmt.Println(fido)
}
