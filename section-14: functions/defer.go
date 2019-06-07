package main

import "fmt"

// A "defer" statement invokes a function whose execution is deferred to th emoment the surrounding function returns,
// either because the surrounding function executed a `return statement`, reached the end of its `function body`,
// or because the corresponding goroutine is `panicking`
func defer_function() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
