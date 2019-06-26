package main

import "fmt"

func directional_channel() {
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("--------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)

	// doesn't assign - specific to general
	// c = cr
	// c = cs

	// does assign - general to specific
	// cr = c
	// cs = c
}
