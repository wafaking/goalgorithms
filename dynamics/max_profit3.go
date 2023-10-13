package dynamics

import "goalgorithms/common"

//买卖股票的最佳时机III(leetcode-123)
//给定一个数组，它的第i个元素是一支给定的股票在第i天的价格。
//设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

//示例1:输入：prices=[3,3,5,0,0,3,1,4],输出：6;
//解释：在第4天（股票价格=0）的时候买入，在第6天（股票价格=3）的时候卖出，这笔交易所能获得利润=3-0=3。
//随后，在第7天（股票价格=1）的时候买入，在第8天（股票价格=4）的时候卖出，这笔交易所能获得利润=4-1=3。
//示例2：输入：prices=[1,2,3,4,5], 输出：4;
//解释：在第1天（股票价格=1）的时候买入，在第5天（股票价格=5）的时候卖出,这笔交易所能获得利润=5-1=4。
//注意你不能在第1天和第2天接连购买股票，之后再将它们卖出。
//因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//示例3：输入：prices=[7,6,4,3,1],输出：0;解释：在这个情况下,没有交易完成,所以最大利润为0。
//示例4：输入：prices=[1],输出：0

// 动态规划
func maxProfit31(prices []int) int {
	var (
		n  = len(prices)
		dp = make([][4]int, n)
	)

	// 初始化：第一天的各种情况
	dp[0][0] = -prices[0] // 第一次持有股票
	dp[0][1] = 0          // 第一次未持有(同一天买入后卖出)
	dp[0][2] = -prices[0] // 第二次持有股票
	dp[0][3] = 0          // 第二次未持有股票
	for i := 1; i < n; i++ {
		// 第一次持有股票: 今天买入 vs 之前已经买入
		dp[i][0] = common.MaxInTwo(-prices[i], dp[i-1][0])
		// 第一次未持有: 今天卖出 vs 之前已经卖出
		dp[i][1] = common.MaxInTwo(prices[i]+dp[i-1][0], dp[i-1][1])
		// 第二次持有股票: 今天持有 vs 之前已经持有
		dp[i][2] = common.MaxInTwo(-prices[i]+dp[i-1][1], dp[i-1][2])
		// 第二次未持有： 今天卖出 vs 之前已经卖出
		dp[i][3] = common.MaxInTwo(prices[i]+dp[i-1][2], dp[i-1][3])
	}

	return common.MaxInTwo(common.MaxInTwo(dp[n-1][0], dp[n-1][1]), common.MaxInTwo(dp[n-1][2], dp[n-1][3]))
}

// 动态规划(一维数组)
func maxProfit32(prices []int) int {
	var (
		buy1  = -prices[0] // 第一次持有股票
		sell1 = 0          // 第一次未持有(同一天买入后卖出)
		buy2  = -prices[0] // 第二次持有股票
		sell2 = 0          // 第二次未持有股票
	)
	for i := 1; i < len(prices); i++ {
		// 第一次持有股票: 今天买入 vs 之前已经买入
		buy1 = common.MaxInTwo(-prices[i], buy1)
		// 第一次未持有: 今天卖出 vs 之前已经卖出
		sell1 = common.MaxInTwo(prices[i]+buy1, sell1)
		// 第二次持有股票: 今天持有 vs 之前已经持有
		buy2 = common.MaxInTwo(-prices[i]+sell1, buy2)
		// 第二次未持有： 今天卖出 vs 之前已经卖出
		sell2 = common.MaxInTwo(prices[i]+buy2, sell2)
	}
	return common.MaxInTwo(common.MaxInTwo(buy1, sell1), common.MaxInTwo(buy2, sell2))
}
