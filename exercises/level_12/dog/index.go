// Package dog provides an implementation of cool stuff for dogs
package dog

const humanToDogYearsFactor = 7

// Years function returns the dog human-equivalent years
func Years(humanYears int) int {
	return humanToDogYearsFactor * humanYears
}
