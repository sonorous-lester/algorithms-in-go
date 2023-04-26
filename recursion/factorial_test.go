package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testcase struct {
	input  int
	expect int
}

func TestFactorial(t *testing.T) {
	tss := []testcase{
		{
			input:  5,
			expect: 120,
		},
		{
			input:  8,
			expect: 40320,
		},
	}

	for _, ts := range tss {
		actual := calculateFactorial(ts.input)
		assert.ElementsMatch(t, ts.expect, actual)
	}

}

func calculateFactorial(num int) int {
	// base case

	// recursive case

	return 0
}
