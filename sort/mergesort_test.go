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
			in:     []int{3, 2, 1},
			expect: []int{1, 2, 3},
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

	if len(nums) == 1 {
		return nums
	}

	middle := len(nums) / 2
	left := mergeSort(nums[:middle])
	right := mergeSort(nums[middle:])
	return merge(left, right)

}

func merge(left, right []int) []int {

	var out []int

	leftIndex := 0
	rightIndex := 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			out = append(out, left[leftIndex])
			leftIndex++
		} else {
			out = append(out, right[rightIndex])
			rightIndex++
		}
	}

	for ; rightIndex < len(right); rightIndex++ {
		out = append(out, right[rightIndex])
	}

	for ; leftIndex < len(left); leftIndex++ {
		out = append(out, left[leftIndex])
	}

	return out

}
