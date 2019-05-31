package main

import "fmt"

/**
examples: https://github.com/golang/go/wiki/Iota
*/
const (
	_  = iota // ignore first value by assigning to blank identifier
	kb = 1 << (iota * 10)
	mb
	gb
)

/**
can also be done as
const (
	_  = iota // ignore first value by assigning to blank identifier
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)
*/

func main() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
