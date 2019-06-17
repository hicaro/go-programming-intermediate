package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hi, I'm", p.First, p.Last)
}

func exercise_2() {
	p := person{
		First: "John",
		Last:  "Lennon",
		Age:   48,
	}

	saySomething(&p)
}

func saySomething(h human) {
	h.speak()
}
