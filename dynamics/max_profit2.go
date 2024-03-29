package dynamics

import "goalgorithms/common"

//买卖股票的最佳时机II(leetcode-122)
//给你一个整数数组prices，其中prices[i]表示某支股票第i天的价格。
//在每一天，你可以决定是否购买和/或出售股票。你在任何时候最多只能持有一股股票。你也可以先购买，然后在同一天出售。
//返回你能获得的最大利润。

//示例1：输入：prices=[7,1,5,3,6,4],输出：7;
//解释：在第2天（股票价格=1）的时候买入，在第3天（股票价格=5）的时候卖出,这笔交易所能获得利润=5-1=4。
//随后，在第4天（股票价格=3）的时候买入，在第5天（股票价格=6）的时候卖出,这笔交易所能获得利润=6-3=3。总利润为4+3=7。
//示例2：输入：prices=[1,2,3,4,5]输出：4;
// 解释：在第1天（股票价格=1）的时候买入，在第5天（股票价格=5）的时候卖出,这笔交易所能获得利润=5-1=4。总利润为4。
//示例3：输入：prices=[7,6,4,3,1], 输出：0;
// 解释：在这种情况下,交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为0。

// 一次遍历(根据拆线图，只取上升段)
func maxProfit21(prices []int) int {
	var max int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}

// 动态规划
func maxProfit22(prices []int) int {

	var (
		n  = len(prices)
		dp = make([][2]int, n)
	)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		// 1. 第i天没有股票的情况：今天本来就未持有股票和今天卖出股票后未持有股票
		// dp[i-1][0]:今天未持有股票，因此取前一天未持有股票的收益
		//prices[i]+dp[i-1][1]: 今天卖出股票的收益：今天的价格+昨天持有股票的收益(即-price[i-1])
		dp[i][0] = common.MaxInTwo(dp[i-1][0], prices[i]+dp[i-1][1])

		// 1. 第i天持有股票的情况：今天本来就持有股票和今天刚买入股票两种情况
		// dp[i-1][0]:今天未持有股票，因此取前一天未持有股票的收益
		//prices[i]+dp[i-1][1]: 今天卖出股票的收益：今天的价格+昨天持有股票的收益(即-price[i-1])
		dp[i][1] = common.MaxInTwo(-prices[i]+dp[i-1][0], dp[i-1][1])

	}
	return common.MaxInTwo(dp[n-1][0], dp[n-1][1])
}
