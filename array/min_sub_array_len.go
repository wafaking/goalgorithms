package array

import (
	"goalgorithms/common"
	"math"
	"sort"
)

// 长度最小的子数组(leetcode-209)
// 给定一个含有n个正整数的数组和一个正整数target。
// 找出该数组中满足其总和[大于等于]target的长度最小的连续子数组
// [numsl,numsl+1,...,numsr-1,numsr]，并返回其长度。如果不存在符合条件的子数组，返回0。
// 示例1：输入：target=7,nums=[2,3,1,2,4,3]输出：2
// 解释：子数组[4,3]是该条件下的长度最小的子数组。
// 示例2：输入：target=4,nums=[1,4,4]输出：1
// 示例3：输入：target=11,nums=[1,1,1,1,1,1,1,1]输出：0
// 进阶：
//	如果你已经实现O(n)时间复杂度的解法,请尝试设计一个O(nlog(n))时间复杂度的解法。

// 暴力破解(易超时)
func minSubArrayLen1(target int, nums []int) int {
	var (
		ans = math.MaxInt64
		sum int
	)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				ans = common.MinInTwo(ans, j-i+1)
				break
			}
		}
		sum = 0
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

// 滑动窗口
func minSubArrayLen2(target int, nums []int) int {
	var ans = math.MaxInt64
	for i, j := 0, 0; j < len(nums); {
		// 每次减去右窗口
		target -= nums[j]
		// 和大于等于target，则获取长度，并将左窗口右移
		for target <= 0 {
			ans = common.MinInTwo(ans, j-i+1)
			target += nums[i]
			i++
		}
		// 右窗口右移
		j++
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

// 前缀和+二分查找
func minSubArrayLen3(target int, nums []int) int {
	var (
		ans  = math.MaxInt64
		n    = len(nums)
		sums = make([]int, n+1)
	)
	// 为了方便计算，令 size = n + 1
	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
	// sums[1] = A[0] 前 1 个元素的前缀和

	// 构造前缀和数组sums
	for i := 1; i < n+1; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	// 		2,3,1,2,4,3
	// sums 0,2,5,6,8,12,15
	// 即sums[j]-sums[i]表示nums中[i]到nums[j-1]的和,即sums[3]-sums[1]表示num[1]到nums[3]的和即4，
	// 因此要求满足sum[j]-sums[i]>=target的最小长度(j-i)
	// 即：target+sums[i]<=sum[j],因此转化成了寻找target+sum[i]在sums数组中的位置、
	for i := 0; i < n+1; i++ {
		// target+sums[i]：把前i个元素的和加到target上相当于移除了前i个元素,即从num[i-1]个元素又算了一遍和
		sum := target + sums[i]
		pivot := sort.SearchInts(sums, sum)
		if pivot <= n {
			ans = common.MinInTwo(ans, pivot-i)
		}
	}

	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

// 双指针
func minSubArrayLen4(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		minLength = math.MaxInt
		sum       int
	)

	for i, j := 0, 0; j < len(nums); {
		sum += nums[j]
		for sum >= target {
			minLength = common.MinInTwo(minLength, j-i+1)
			sum -= nums[i]
			i++
		}
		j++
	}

	if minLength == math.MaxInt {
		return 0
	}
	return minLength
}
