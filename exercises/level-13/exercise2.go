package main

import (
	"fmt"
	"go-programming/exercises/level-13/quote"
	"go-programming/exercises/level-13/word"
)

func exercise2() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
