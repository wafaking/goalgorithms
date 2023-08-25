package heap

import "testing"

func TestHeapSort(t *testing.T) {
	var list = [][]int{
		{33, 24, 8, 3, 10, 15, 16, 15, 30, 17, 19},
		{2, 7, 4, 1, 8, 1},
		{90, 70, 8, 60, 10, 4, 5, 30, 20},
	}
	for _, item := range list {
		res := HeapSort(item)
		t.Logf("res: %v", res)
	}
}
