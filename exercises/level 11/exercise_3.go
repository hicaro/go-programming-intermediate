package main

import (
	"fmt"
	"log"
)

type customErr struct {
	message string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Custom Error: %s", ce.message)
}

func exercise_3() {

	err := customErr{
		message: "This is an error",
	}

	func(e error) {
		log.Println(e)
	}(err)
}
