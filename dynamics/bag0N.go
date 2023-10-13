package dynamics

import "goalgorithms/common"

// 完全背包问题
// 有n件物品和一个最多能背重量为w的背包。第i件物品的重量是weight[i]，得到的价值是value[i]。
// 每件物品可以使用无数次，求解将哪些物品装入背包里物品价值总和最大。

// 动态规划
func weightBag31(weight, value []int, bagWeight int) int {
	var dp = make([][]int, len(weight)+1)
	for i := 0; i <= len(dp); i++ {
		dp[i] = make([]int, bagWeight+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < bagWeight+1; j++ {
			if bagWeight < weight[i-1] {
				// 此处与0/1背包、多重背包不同，因为同一物品可以多次选择
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = common.MaxInTwo(dp[i-1][j], value[i-1]+dp[i][j-weight[i-1]])
			}
		}
	}

	return dp[len(dp)-1][bagWeight]
}
