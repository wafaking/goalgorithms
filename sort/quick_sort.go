package sort

import (
	"math/rand"
	"time"
)

func QuickSort1(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	var quickSort func(nums []int) []int
	rand.Seed(time.Now().UnixNano())
	quickSort = func(nums []int) []int {
		if len(nums) <= 1 {
			return nums
		}
		var (
			pivot       = nums[rand.Intn(len(nums)-1)]
			left, right = make([]int, 0), make([]int, 0)
		)
		for i := 1; i < len(nums); i++ {
			if nums[i] > pivot {
				right = append(right, nums[i])
			} else {
				left = append(left, nums[i])
			}
		}
		return append(append(quickSort(left), pivot), quickSort(right)...)
	}
	return quickSort(nums)
}
