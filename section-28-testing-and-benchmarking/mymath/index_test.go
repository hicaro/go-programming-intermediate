package mymath

import (
	"fmt"
	"testing"
)

type testpair struct {
	values   []float64
	expected float64
}

var avgTests = []testpair{
	{values: []float64{1, 2}, expected: 1.5},
	{values: []float64{1, 2, 3}, expected: 2},
	{values: []float64{1, 4, 13}, expected: 6},
}

var sumTests = []testpair{
	{values: []float64{1, 2}, expected: 3},
	{values: []float64{1, 2, 3}, expected: 6},
	{values: []float64{1, 4, 13}, expected: 18},
}

func ExampleAverage() {
	fmt.Println(Average(1, 2))
	// Output:
	// 1.5
}

func TestAverage(t *testing.T) {
	for _, pair := range avgTests {
		v := Average(pair.values...)

		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}

func ExampleSum() {
	fmt.Println(Sum(1, 2))
	// Output:
	// 3
}

func TestSum(t *testing.T) {
	for _, pair := range sumTests {
		v := Sum(pair.values...)

		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}
