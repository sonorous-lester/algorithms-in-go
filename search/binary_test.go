package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testcase struct {
	input            []int
	numberWantToFind int
	result           bool
}

func TestBinarySearch(t *testing.T) {

	cases := []testcase{
		{
			input:            []int{1, 2, 3, 4, 6, 9, 12, 34, 68, 89},
			numberWantToFind: 12,
			result:           true,
		},
		{
			input:            []int{1, 2, 3, 4, 6, 9, 12, 34, 68, 89},
			numberWantToFind: 20,
			result:           false,
		},
	}

	for _, ts := range cases {
		find := binarySearch(ts.input)
		assert.Equal(t, ts.result, find)
	}

}

func binarySearch(numbers []int) bool {
	return false
}
