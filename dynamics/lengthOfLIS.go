package dynamics

import "goalgorithms/common"

// 最长递增子序列(leetcode-300)
// 给你一个整数数组nums，找到其中最长严格递增子序列的长度。
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
// 示例1：输入nums=[10,9,2,5,3,7,101,18],输出：4,解释：最长递增子序列是[2,3,7,101]，因此长度为4 。
// 示例2：输入nums=[0,1,0,3,2,3],输出：4
// 示例3：输入nums=[7,7,7,7,7,7,7],输出：1

// 动态规划
func lengthOfLIS11(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		maxNum = 1
		// dp[i]表示前i个元素中的最长递增子数列数，一定包含nums[i]
		dp = make([]int, len(nums))
	)
	for i := range dp {
		// 初始化都为1,即最短递增子序列只有一个元素
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = common.MaxInTwo(dp[j]+1, dp[i])
			}
		}
		maxNum = common.MaxInTwo(maxNum, dp[i])
	}
	//for _, v := range dp {
	//	maxNum = common.MaxInTwo(v, maxNum)
	//}

	return maxNum
}

// 动态规划+深度遍历
func lengthOfLIS12(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		maxNum = 0 // 初始化为0,最后都加1即可
		dp     = make([]int, len(nums))
		dfs    = func(i int) {
			for j := 0; j < i; j++ {
				if nums[j] < nums[i] {
					dp[i] = common.MaxInTwo(dp[j]+1, dp[i])
				}
			}
		}
	)

	for i := 1; i < len(nums); i++ {
		dfs(i)
	}
	for _, v := range dp {
		maxNum = common.MaxInTwo(v, maxNum)
	}

	return maxNum + 1
}

// 动态规划，法一的简单版
func lengthOfLIS13(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		maxLength int
		dp        = make([]int, len(nums))
	)

	for i, x := range nums {
		for j, y := range nums[:i] {
			if y < x {
				dp[i] = common.MaxInTwo(dp[i], dp[j])
			}
		}
		dp[i]++
		maxLength = common.MaxInTwo(maxLength, dp[i])
	}
	return maxLength
}

// lengthOfLIS14 贪心算法
func lengthOfLIS14(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// list用来存储最长递增子序列
	var list = make([]int, 0)
	for _, x := range nums {
		// 从左到右遍历list数组,找到合适位置
		var i = 0
		for i < len(list) && x > list[i] {
			i++
		}
		if i == len(list) {
			list = append(list, x)
		} else {
			list[i] = x
		}
	}
	return len(list)
}

// lengthOfLIS14 贪心算法+二分查找
func lengthOfLIS15(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// list用来存储最长递增子序列
	var list = make([]int, 0)
	for _, x := range nums {
		// 从左到右遍历list数组,找到合适位置
		var left, right = 0, len(list)
		// 使用二分查找，找出合适位置
		for left < right {
			mid := (right-left)/2 + left
			if x > list[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == len(list) {
			list = append(list, x)
		} else {
			list[left] = x
		}
	}
	return len(list)
}
