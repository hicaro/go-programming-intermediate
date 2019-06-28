package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	v := Years(5)

	if v != 35 {
		t.Error(
			"For", 5,
			"got", v,
			"expected", 35,
		)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(5)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func TestYearsTwo(t *testing.T) {
	v := YearsTwo(5)

	if v != 35 {
		t.Error(
			"For", 5,
			"got", v,
			"expected", 35,
		)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(5)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(5))
	// Output:
	// 35
}
