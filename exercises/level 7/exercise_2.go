package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func exercise_2() {
	p := person{
		first: "John",
		last:  "Doe",
		age:   24,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	(*p).first = "Something else"
}
