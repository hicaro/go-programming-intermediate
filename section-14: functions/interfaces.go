package main

import "fmt"

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("Hello, I'm", p.firstName, p.lastName, "and I'm a regular person")
}

func interfaces() {
	a := agent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
		},
		ltk: true,
	}

	some(a)
	ckeckType(a)

	p := person{
		"John",
		"Doe",
	}

	some(p)
	ckeckType(p)
}

func some(h human) {
	h.speak()
}

func ckeckType(h human) {
	switch h.(type) {
	case person:
		fmt.Println("This is a person")
	case agent:
		fmt.Println("This is an agent")
	}
}
