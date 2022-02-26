package binary

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var list = []int{2995, 4615, 6336, 7637, 8540, 9083, 9900, 14108, 16122,
		25260, 25675, 29390, 29649, 40660, 45738, 46120, 49689, 53777, 54044,
		57415, 58943, 60660, 79294, 80717, 85363, 90436, 91323, 92904, 95588}

	for i := 2; i < 10; i++ {
		list = list[:i]

		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+1] {
				fmt.Println(list)
				t.Error()
			}
		}
		BinarySearch(list, 2995)
	}
}
