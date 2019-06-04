package main

import "fmt"

func exercise_9() {
	var favSport string = "Soccer"

	switch favSport {
	case "Soccer", "Volleyball", "American Football":
		{
			fmt.Println("You like something")
		}
	default:
		{
			fmt.Println("You do not like anything")
		}
	}
}
