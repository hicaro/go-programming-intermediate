package main

import (
	"fmt"
	"math"
)

type square struct {
	side float32
}

type circle struct {
	radius float32
}

type shape interface {
	area() float32
}

func (s square) area() float32 {
	return s.side * s.side
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func exercise_5() {
	s := square{
		side: 5,
	}
	c := circle{
		radius: 5,
	}

	printArea(s)
	printArea(c)
}

func printArea(s shape) {
	fmt.Println(s.area())
}
