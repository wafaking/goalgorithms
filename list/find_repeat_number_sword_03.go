package list

import (
	"sort"
)

//数组中重复的数字(sword-03)
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
//数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
//示例 1： 输入：[2, 3, 1, 0, 2, 5, 3], 输出：2 或 3

// findRepeatNumber1 使用map
func findRepeatNumber1(nums []int) int {
	var mapNum = map[int]int{}
	for _, v := range nums {
		if _, ok := mapNum[v]; ok {
			return v
		}
		mapNum[v]++
	}
	return -1
}

// findRepeatNumber2 排序，对比相邻元素
func findRepeatNumber2(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i-1]
		}
	}
	return -1
}

// findRepeatNumber3 使用位数组(效率最高且内存消耗最小)
func findRepeatNumber3(nums []int) int {
	n := len(nums)
	sli := make([]int, n, n)

	for _, v := range nums {
		if sli[v] > 0 {
			return v
		}
		sli[v] = 1
	}
	return -1
}

//[3,4,2,1,1,0]
//[1,4,2,3,1,0]
// findRepeatNumber4 使用原地置换
func findRepeatNumber4(nums []int) int {
	for i := 0; i < len(nums); {
		if i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else {
			i++
		}
	}

	return -1
}
