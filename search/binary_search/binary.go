package binary

import (
	"fmt"
)

func BinarySearch(sli []int, num int) {
	var low, high = 0, len(sli)
	for low < high {
		mid := (low + high) / 2
		if sli[mid] == num {
			fmt.Println("find it........................")
			return
		} else if sli[mid] > num {
			high = mid
			BinarySearch(sli[0:mid], num)
			break
		} else {
			low = mid + 1
			BinarySearch(sli[mid+1:], num)
			break
		}
	}
}
