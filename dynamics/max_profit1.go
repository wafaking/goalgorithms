package dynamics

import "goalgorithms/common"

//买卖股票的最佳时机(leetcode-121)
//给定数组prices，第i个元素prices[i]表示一支给定股票第i天的价格。
//你只能选择某一天买入这只股票，并选择在未来的某一个不同的日子卖出该股票。设计一个算法来计算你所能获取的最大利润。
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回0。

//示例1：输入：[7,1,5,3,6,4],输出：5
//解释：在第2天（股票价格=1）的时候买入，在第5天（股票价格=6）的时候卖出，最大利润=6-1=5。
//注意利润不能是7-1=6,因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//示例2：输入：prices=[7,6,4,3,1],输出：0,解释：在这种情况下,没有交易完成,所以最大利润为0。

//双层遍历(超出时间限制)
func maxProfit11(prices []int) int {
	var max int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			max = common.MaxInTwo(max, prices[j]-prices[i])
		}
	}
	return max
}

// 一次遍历
func maxProfit12(prices []int) int {
	var min, max = prices[0], 0
	for i := 0; i < len(prices); i++ {
		// 每次取最大值
		max = common.MaxInTwo(prices[i]-min, max)
		// 不断取最小数值
		min = common.MinInTwo(min, prices[i])
	}
	return max
}

// 动态规划
func maxProfit13(prices []int) int {
	// dp[i][0]表示第i天未持有股票时的收益
	// dp[i][1]表示第i天持有股票的收益
	if len(prices) == 0 {
		return 0
	}
	var (
		n  = len(prices)
		dp = make([][2]int, n)
	)
	// dp[0][0]表示第1天未持有股票收益为0
	// dp[1][0]表示第1天持有股票收益为-price[0]
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		// dp[i-1][0]表示第i-1天也未持有股票
		// prices[i]+dp[i-1][0]表示第i天卖出股票后未持有,dp[i-1][1]表示前一天持有股票的收益
		dp[i][0] = common.MaxInTwo(dp[i-1][0], prices[i]+dp[i-1][1])

		// dp[i-1][1]表示第i-1天持有股票
		//  -prices[i]表示第i天买入股票的收益
		dp[i][1] = common.MaxInTwo(dp[i-1][1], -prices[i])
	}
	return common.MaxInTwo(dp[n-1][0], dp[n-1][1])
}
