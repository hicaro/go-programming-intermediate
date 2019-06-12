package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) speak() {
	fmt.Printf("Hi, I'm %s %s and am %d years old\n", p.firstName, p.lastName, p.age)
}

func exercise_4() {
	p := person{
		firstName: "James",
		lastName:  "Morrison",
		age:       53,
	}

	p.speak()
}
