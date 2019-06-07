package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type agent struct {
	person
	ltk bool
}

// func (r receiver) idetifier(parameters) (return(s)) { code }

func (a agent) speak() {
	fmt.Println("Hello, I'm", a.person.firstName, a.person.lastName)
}

func methods() {
	a := agent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
		},
		ltk: true,
	}

	a.speak()
}
