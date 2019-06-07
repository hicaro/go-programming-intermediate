package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	flavors   []string
}

func exercise_1() {
	p1 := person{
		firstName: "Josh",
		lastName:  "Morrad",
		flavors:   []string{"Vanilla", "Chocolatte"},
	}

	fmt.Printf("%s %s likes the following flavors:\n", p1.firstName, p1.lastName)
	for _, v := range p1.flavors {
		fmt.Printf("\t%s\n", v)
	}
}
