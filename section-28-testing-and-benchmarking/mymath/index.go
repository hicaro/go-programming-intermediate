// Package mymath implement custom math functions
package mymath

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
