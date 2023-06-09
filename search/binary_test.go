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
			input:            []int{1, 2, 3, 4, 6, 9, 12, 34, 68},
			numberWantToFind: 20,
			result:           false,
		},
		{
			input:            []int{-1, 2},
			numberWantToFind: 0,
			result:           false,
		},
		{
			input:            []int{-1, 2},
			numberWantToFind: 2,
			result:           true,
		},
		{
			input:            []int{0},
			numberWantToFind: 2,
			result:           false,
		},
	}

	for _, ts := range cases {
		find := binarySearch(ts.numberWantToFind, ts.input)
		assert.Equal(t, ts.result, find)
	}

	for _, ts := range cases {
		find := binarySearch2(ts.numberWantToFind, ts.input)
		assert.Equal(t, ts.result, find)
	}

}

func binarySearch(want int, numbers []int) bool {
	//	find the middle number in the array(actual)
	//	if the actual is equal to want return true
	//	if the actual is latest and is not equal to want return false
	//	if the actual is bigger then want
	//		search right side of the array
	//	if the actual is smaller than want
	//		search left side of the array

	if len(numbers) == 1 && numbers[0] != want {
		return false
	}

	middle := (len(numbers) - 1) / 2
	actual := numbers[middle]
	if actual == want {
		return true
	}

	if actual > want {
		return binarySearch(want, numbers[:middle])
	} else {
		return binarySearch(want, numbers[middle:])
	}
}

func binarySearch2(want int, nums []int) bool {

	start := 0
	end := len(nums) - 1

	for start <= end {
		middle := (start + end) / 2
		if nums[middle] == want {
			return true
		}

		if nums[middle] > want {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return false
}
