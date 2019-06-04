package main

import (
	"fmt"
	"strings"
)

const (
	UPPER = 1 // upper case
	LOWER = 2 // lower case
	CAP   = 4 // capitalizes
	REV   = 8 // reverses
)

func main() {
	s := processString("HELLO WORLD!", LOWER|CAP|REV)
	fmt.Println(s)
}

func processString(str string, conf byte) string {
	reverser := func(s string) string {
		runes := []rune(s)
		n := len(runes)

		for i := 0; i < n/2; i++ {
			runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
		}

		return string(runes)
	}

	// query config bits
	if (conf & UPPER) != 0 {
		str = strings.ToUpper(str)
	}

	if (conf & LOWER) != 0 {
		str = strings.ToLower(str)
	}

	if (conf & CAP) != 0 {
		str = strings.Title(str)
	}

	if (conf & REV) != 0 {
		str = reverser(str)
	}

	return str
}
