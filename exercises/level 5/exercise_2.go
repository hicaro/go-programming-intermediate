package main

import "fmt"

func exercise_2() {
	p1 := person{
		firstName: "Josh",
		lastName:  "Morrad",
		flavors:   []string{"Vanilla", "Chocolatte"},
	}

	people := map[string]person{}

	people[p1.lastName] = p1

	fmt.Println(people)

	for key, value := range people {
		fmt.Println(key, value)
	}
}
