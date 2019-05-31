package main

import "fmt"

var a int = 42
var b string = "James Bond"
var c bool = true

func exercise_3() {
	s := fmt.Sprintf("%d, %s, %t", a, b, c)
	fmt.Println(s)
}
