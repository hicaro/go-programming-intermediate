// Package mymath implement custom math functions
package mymath

import "sort"

// Average function returns the average of the provided numbers
func Average(numbers ...float64) float64 {
	var avg float64
	var total float64

	l := len(numbers)
	if l > 0 {
		for _, number := range numbers {
			total += number
		}

		avg = total / float64(l)
	}

	return avg
}

// Sum function returns the sum of the provided numbers
func Sum(numbers ...float64) float64 {
	var sum float64

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// CenteredAverage computes the average of a list of numbers
// after removing the largest and smallest values.
func CenteredAverage(xi ...float64) float64 {
	sort.Float64s(xi)
	xi = xi[1:(len(xi) - 1)]

	var n float64
	for _, v := range xi {
		n += v
	}

	f := float64(n) / float64(len(xi))
	return f
}
