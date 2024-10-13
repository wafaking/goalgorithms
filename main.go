package main

import "fmt"

func main() {
	var list = []int{2, 1, 3, 4}
	list = quickSort2(list)
	fmt.Println(list)
}

func quickSort2(list []int) []int {
	if len(list) < 2 {
		return list
	}
	var (
		l, r  = 0, len(list) - 1
		pivot = list[r]
	)

	for i := range list {
		// 将小于Pivot的元素都交换到最左侧
		if list[i] < pivot {
			list[i], list[l] = list[l], list[i]
			l++
		}
	}
	list[l], list[r] = list[r], list[l]

	quickSort2(list[:l])
	quickSort2(list[l+1:])
	return list
}
