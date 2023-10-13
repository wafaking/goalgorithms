package dynamics

import (
	"goalgorithms/common"
	"math"
)

//零钱兑换(leetcode-322)
//给你一个整数数组coins，表示不同面额的硬币；以及一个整数amount，表示总金额。
//计算并返回可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回-1。
//你可以认为每种硬币的数量是无限的。
//示例1：输入：coins=[1,2,5],amount=11输出：3,解释：11=5+5+1
//示例2：输入：coins=[2],amount=3输出：-1
//示例3：输入：coins=[1],amount=0输出：0

// 动态规划
func coinChange(coins []int, amount int) int {
	// dp[i][j]表示使用前i种硬币组成总金额为j所需要的最小硬币数量
	var dp = make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}

	//第一行：使用面值为0的硬币兑换金额为j所需要的最小硬币数(无法兑换，使用-1表示)
	//第一列：使用coins中的硬币兑换金额为0所需要的最小硬币数(无需硬币,因此全为0)
	for i := 1; i < amount+1; i++ {
		dp[0][i] = math.MaxInt32
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < amount+1; j++ {
			if coins[i-1] > j {
				// 取不包含此硬币时兑换j所需的最少的硬币数量
				dp[i][j] = dp[i-1][j]
			} else {
				//dp[i][j-coins[i-1]]表示选用此硬币后剩余的钱对应的最少需要的硬币数量
				dp[i][j] = common.MinInTwo(dp[i-1][j], 1+dp[i][j-coins[i-1]])
			}
		}
	}
	if dp[len(dp)-1][len(dp[0])-1] >= math.MaxInt32 {
		return -1
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
