package main

import "fmt"

func exercise_10() {
	myMap := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	myMap["some_key"] = []string{"Something", "in", "here"}

	delete(myMap, "no_dr")

	for key, array := range myMap {
		fmt.Println(key)

		for index, value := range array {
			fmt.Printf("\t%v: %s\n", index, value)
		}
	}
}
