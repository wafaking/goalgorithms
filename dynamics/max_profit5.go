package dynamics

import "goalgorithms/common"

//买卖股票的最佳时机含冷冻期(leetcode-309)
//给定一个整数数组prices，其中第prices[i]表示第i天的股票价格。
//设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//卖出股票后，你无法在第二天买入股票(即冷冻期为1天)。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

//示例1:输入:prices=[1,2,3,0,2]输出:3
//解释:对应的交易状态为:[买入,卖出,冷冻期,买入,卖出]
//示例2:输入:prices=[1]输出:0

func maxProfit51(prices []int) int {
	//dp[i][j]表示第i天持有股票的最大收益
	var (
		n  = len(prices)
		dp = make([][3]int, n)
	)
	dp[0][0] = -prices[0] // 持有股票
	dp[0][1] = 0          // 未持有股票(今天卖出股票或处于冷冻期)
	dp[0][2] = 0          // 未持有股票(冷冻期之后或一直未持有)
	for i := 1; i < n; i++ {
		// 持有股票:今天持有 vs 之前就持有
		dp[i][0] = common.MaxInTwo(-prices[i]+dp[i-1][2], dp[i-1][0])
		// 未持有股票(今天卖出股票或处于冷冻期,两者相同)
		dp[i][1] = prices[i] + dp[i-1][0]
		// 未持有股票(冷冻期之后或一直未持有)
		dp[i][2] = common.MaxInTwo(dp[i-1][1], dp[i-1][2])
	}
	return common.MaxInTwo(common.MaxInTwo(dp[n-1][0], dp[n-1][1]), dp[n-1][2])
}

// 动态规划(比法一更好理解)
func maxProfit52(prices []int) int {
	//dp[i][j]表示第i天持有股票的最大收益
	var (
		n  = len(prices)
		dp = make([][4]int, n)
	)
	dp[0][0] = -prices[0] // 持有股票
	dp[0][1] = 0          // 未持有股票(今天卖出股票)
	dp[0][2] = 0          // 未持有股票(冷冻期)
	dp[0][3] = 0          // 未持有股票(冷冻期后第一天或n天)
	for i := 1; i < n; i++ {
		// 持有股票:今天买入 vs 之前就持有
		dp[i][0] = common.MaxInTwo(common.MaxInTwo(-prices[i]+dp[i-1][2], -prices[i]+dp[i-1][3]), dp[i-1][0])
		// 未持有股票(今天卖出股票)
		dp[i][1] = prices[i] + dp[i-1][0]
		// 未持有股票(冷冻期)
		dp[i][2] = dp[i-1][1]
		// 未持有股票(冷冻期后第一天或n天)
		dp[i][3] = common.MaxInTwo(dp[i-1][2], dp[i-1][3])
	}
	return common.MaxInNums(dp[n-1][0], dp[n-1][1], dp[n-1][2], dp[n-1][3])
}
