package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tss := []testcase{
		{
			in:     []int{9, 8, 5, 7, 3, 2, 1, 4, 6, 0},
			expect: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			in:     []int{2, 1},
			expect: []int{1, 2},
		},
		{
			in:     []int{0},
			expect: []int{0},
		},
	}

	for _, ts := range tss {
		actual := mergeSort(ts.in)
		assert.ElementsMatch(t, ts.expect, actual)
	}

}

func mergeSort(nums []int) []int {
	return nil
}