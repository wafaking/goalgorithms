package heap

import "testing"

func TestHeapSort(t *testing.T) {
	var list = [][]int{
		{33, 24, 8, 3, 10, 15, 16, 15, 30, 17, 19},
	}
	for _, item := range list {
		res := HeapSort(item)
		t.Logf("res: %v", res)
	}
}
