package main

import (
	"fmt"
	"go-programming/section-28-testing-and-benchmarking/mymath"
)

/*
	Tests must:
		* be in a file that ends with `_test.go`
		* put the file in the same package as the one being tested
		* be in a func with an signature `func TestXxx(*testing.T)`

	Run a test
		* go test

	Deal with test failure
		* use t.Error to signal failure

	Nice idiom
		* expected
		* got

	Coverage
		* go test -cover ./...
		* go test -coverprofile cover.out ./...
		* go tool cover -html=cover.out

*/

func main() {
	values := []float64{3, 4, 5}

	fmt.Println(mymath.Average(values...))
	fmt.Println(mymath.Sum(values...))
}
