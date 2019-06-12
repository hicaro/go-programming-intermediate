package main

import (
	"fmt"
	"math"
)

/**
Method Sets

a NON-POINTER RECEIVER works with values that are POINTERS or NON-POINTERS.
a POINTER RECEIVER only works with values that are POINTERS.

Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T

*/

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{5}
	info(&c)
	// fmt.Println(c.area())
}
