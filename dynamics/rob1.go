package dynamics

import "goalgorithms/common"

// 打家劫舍(leetcode-198)
// 一个小偷计划偷窃沿街的房屋。每间房内都藏有一定的现金，但相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。给定一个代表每个房屋存放金额的非负整数数组，
// 计算你不触动警报装置的情况下，一夜之内能够偷窃到的最高金额。

// 示例1：
// 输入：[1,2,3,1], 输出：4
// 解释：偷窃1号房屋(金额=1)，然后偷窃3号房屋(金额=3),偷窃到的最高金额为1+3=4。

// 示例2：输入：[2,7,9,3,1],输出：12;
// 解释：偷窃1号房屋(金额=2), 偷窃3号房屋(金额=9)，接着偷窃5号房屋(金额=1)。
// 偷窃到的最高金额=2+9+1=12 。

// 动态规划
func rob11(nums []int) int {
	// 状态:f(n)表示前n个房间可以偷窥到的最大金额
	// 状态转移方程f(n)=max(f(n-1),f(n-2)+nums[n-1])

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	// 使用dp数组记录前n个房间的最大金额
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = common.MaxInTwo(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = common.MaxInTwo(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

// 动态规划，动态变量
func rob12(nums []int) int {
	// 状态:f(n)表示前n个房间可以偷窥到的最大金额
	// 状态转移方程f(n)=max(f(n-1),f(n-2)+nums[n-1])

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	var pre = nums[0]
	var next = common.MaxInTwo(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		pre, next = next, common.MaxInTwo(next, pre+nums[i])
	}

	return next
}

// 动态规划，动态数组(修改原有数组)
func rob13(nums []int) int {
	// 状态:f(n)表示前n个房间可以偷窥到的最大金额
	// 状态转移方程f(n)=max(f(n-1),f(n-2)+nums[n-1])

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	nums[1] = common.MaxInTwo(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i] = common.MaxInTwo(nums[i-2]+nums[i], nums[i-1])
	}

	return nums[len(nums)-1]
}
