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
	}

	for _, ts := range cases {
		find := binarySearch(ts.numberWantToFind, ts.input)
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

	middle := len(numbers) / 2
	actual := numbers[middle]
	if actual == want {
		return true
	}

	if middle == 0 && actual != want {
		return false
	}

	if actual > want {
		return binarySearch(want, numbers[:middle])
	} else {
		return binarySearch(want, numbers[middle:])
	}
}
