package heap

import (
	"goalgorithms/common"
	"testing"
)

func TestHeapSort(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{0, 6, 2, 3, 1, 4, 5, 7, 8},
			Expected: []int{0, 1, 2, 3, 4, 5},
		},
		{
			Nums:     []int{33, 24, 8, 3, 10, 15, 16, 15, 30, 17, 19},
			Expected: []int{},
		},
		{
			Nums:     []int{2, 7, 4, 1, 8, 1},
			Expected: []int{},
		},
		{
			Nums:     []int{90, 70, 8, 60, 10, 4, 5, 30, 20},
			Expected: []int{},
		},
	}
	var res = make([]int, 0)
	for _, item := range list {
		res = HeapSort(item.Nums)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)
		t.Logf("res: %v", res)
	}
}
