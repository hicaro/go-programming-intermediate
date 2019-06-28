package word

import (
	"testing"
)

type testPair struct {
	input  string
	output int
}

var values = []testPair{
	testPair{
		input:  "",
		output: 0,
	},
	testPair{
		input:  "This is mine!",
		output: 3,
	},
	testPair{
		input:  "Oh, don't be silly - EVERYONE wants this. Everyone wants to be *us*.",
		output: 13,
	},
	testPair{
		input:  "Once upon a time, there was a princess.",
		output: 8,
	},
}

func TestCount(t *testing.T) {
	for _, pair := range values {
		v := Count(pair.input)
		if v != pair.output {
			t.Error(
				"For", pair.input,
				"got", v,
				"expected", pair.output,
			)
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("Oh, don't be silly - EVERYONE wants this. Everyone wants to be *us*.")
	}
}
