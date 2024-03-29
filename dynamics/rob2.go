package dynamics

import "goalgorithms/common"

// 打家劫舍II(leetcode-213)
// 一个小偷计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。
// 同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，今晚能够偷窃到的最高金额。
// 示例1：输入：nums=[2,3,2], 输出：3,解释：你不能先偷窃1号房屋（金额=2），然后偷窃3号房屋（金额=2）,因为他们是相邻的。
// 示例2：输入：nums=[1,2,3,1], 输出：4,解释：你可以先偷窃1号房屋（金额=1），然后偷窃3号房屋(金额=3)偷窃到的最高金额=1+3=4。
// 示例3：输入：nums=[1,2,3],输出：3;

// 动态规划（将情况分成:不偷窃第一间和偷窃第一间两种情况）
func rob21(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return common.MaxInTwo(nums[0], nums[1])
	}

	var rob = func(nums []int) int {
		var pre, cur = nums[0], common.MaxInTwo(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			pre, cur = cur, common.MaxInTwo(pre+nums[i], cur)
		}
		return common.MaxInTwo(pre, cur)
	}
	//nums[1:]表示不偷第一间，num[:len(nums)-1]表示偷第一间不偷最后一间
	return common.MaxInTwo(rob(nums[1:]), rob(nums[0:len(nums)-1]))
}
