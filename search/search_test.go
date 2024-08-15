package search

import (
	"goalgorithms/common"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var list = []item1{
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
		{
			nums:     []int{1, 2},
			target:   1,
			expected: 0,
		},
		{
			nums:     []int{1, 2},
			target:   2,
			expected: 1,
		},
		{
			nums:     []int{1, 2},
			target:   0,
			expected: -1,
		},
		{
			nums:     []int{1, 2},
			target:   3,
			expected: -1,
		},
	}

	var res int
	for _, item := range list {
		res = binarySearch1(item.nums, item.target)
		common.PrintDiffTwoInt(res, item.expected, item, t)

		res = binarySearch2(item.nums, item.target)
		common.PrintDiffTwoInt(res, item.expected, item, t)

		res = binarySearch3(item.nums, item.target)
		common.PrintDiffTwoInt(res, item.expected, item, t)

		common.PrintSplit(t)

	}
}
