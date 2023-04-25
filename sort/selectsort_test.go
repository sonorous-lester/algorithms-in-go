package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testcase struct {
	in     []int
	expect []int
}

func TestSelectionSort(t *testing.T) {
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
		actual := selectionSort(ts.in)
		assert.ElementsMatch(t, ts.expect, actual)
	}

}

func selectionSort(nums []int) []int {
	// For i from 0 to n-1
	//  Find the smallest number between numbers[i] and numbers[n-1]
	//	Swap the smallest number with numbers[i]

	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				minIndex = j
			}
		}
		if minIndex != i {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
	return nums
}
