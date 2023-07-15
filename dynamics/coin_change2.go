package dynamics

//零钱兑换II(leetcode-518)
//给你一个整数数组coins表示不同面额的硬币，另给一个整数amount表示总金额。
//计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回0。
//假设每一种面额的硬币有无限个。

//示例1：输入：amount=5,coins=[1,2,5]输出：4,解释：有四种方式可以凑成总金额：
//	5=5
//	5=2+2+1
//	5=2+1+1+1
//	5=1+1+1+1+1
//示例2：输入：amount=3,coins=[2]输出：0;解释：用面额2的硬币不能凑成总金额3。
//示例3：输入：amount=10,coins=[10]输出：1;
//示例4：输入：amount=0,coins=[7]输出：1;

// 动态规划，二维数组
func change1(amount int, coins []int) int {
	// dp[i][j]表示使用前i种硬币组成总金额为j的组合数
	//sort.Ints(coins) 可以不排序
	var dp = make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}

	//第一行：使用面值为0的硬币兑换金额为j的组合数为0
	//第一列：使用coins中的硬币兑换金额为0的组合数为0
	for i := 1; i < len(dp); i++ {
		for j := 0; j < amount+1; j++ {
			// 注：根据示例4,如果总金额为0，可以有1种方案就是什么都不取
			if j == 0 {
				dp[i][j] = 1
				continue
			}

			if coins[i-1] > j {
				// 取不包含此硬币时兑换j的组合数
				dp[i][j] = dp[i-1][j]
			} else {
				// 1+用一枚硬币coins[i-1]兑换j剩下的金额对应的组合数
				// 选择此硬币+不选择此硬币的方案数之和
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

// 动态规划（一维数组）
func change2(amount int, coins []int) int {
	// 原理：由原来的二维数据退化成滚动数组，再退化成一维数组
	// dp[i]表示当总金额为i时的组合数
	var dp = make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins { // 遍历每种硬币，
		for i := 1; i < amount+1; i++ { // 分别计算对于指定金额选用此硬币的组合数
			if i >= coin {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[amount]
}
