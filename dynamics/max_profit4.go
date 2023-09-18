package dynamics

//买卖股票的最佳时机IV(leetcode-188)
//给你一个整数数组prices和一个整数k，其中prices[i]是某支给定的股票在第i天的价格。
//设计一个算法来计算你所能获取的最大利润。你最多可以完成k笔交易。也就是说，你最多可以买k次，卖k次。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

//示例1：输入：k=2,prices=[2,4,1]输出：2
//解释：在第1天(股票价格=2)的时候买入，在第2天(股票价格=4)的时候卖出，这笔交易所能获得利润=4-2=2。
//示例2：输入：k=2,prices=[3,2,6,5,0,3]输出：7
//解释：在第2天(股票价格=2)的时候买入，在第3天(股票价格=6)的时候卖出,这笔交易所能获得利润=6-2=4。
//随后，在第5天(股票价格=0)的时候买入，在第6天(股票价格=3)的时候卖出,这笔交易所能获得利润=3-0=3。

// 动态规划(参考maxProfit31)
func maxProfit41(k int, prices []int) int {
	var (
		n  = len(prices)
		dp = make([][]int, n)
	)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2*k)
	}

	// 初始化：第一天的各种情况
	for i := 0; i < 2*k; i++ {
		if i&1 == 0 {
			dp[0][i] = -prices[0]
		} else {
			dp[0][i] = -0
		}
	}
	for i := 1; i < n; i++ {
		// 第一次持有股票: 今天买入 vs 之前已经买入
		dp[i][0] = maxInTwo(-prices[i], dp[i-1][0])
		for j := 1; j < 2*k; j++ {
			if j&1 == 1 {
				// 第一次未持有: 今天卖出 vs 之前已经卖出
				dp[i][j] = maxInTwo(prices[i]+dp[i-1][j-1], dp[i-1][j])
			} else {
				// 第二次持有股票: 今天持有 vs 之前已经持有
				dp[i][j] = maxInTwo(-prices[i]+dp[i-1][j-1], dp[i-1][j])
			}
		}
	}

	return maxInNums(dp[n-1])
}

// 动态规划+一维数组(参考maxProfit32)
func maxProfit42(k int, prices []int) int {
	var (
		n  = len(prices)
		dp = make([]int, 2*k)
	)

	// 初始化：第一天的各种情况
	for i := 0; i < 2*k; i++ {
		if i&1 == 0 {
			dp[i] = -prices[0]
		} else {
			dp[i] = -0
		}
	}
	for i := 1; i < n; i++ {
		// 第一次持有股票: 今天买入 vs 之前已经买入
		dp[0] = maxInTwo(-prices[i], dp[0])
		for j := 1; j < 2*k; j++ {
			if j&1 == 1 {
				// 第一次未持有: 今天卖出 vs 之前已经卖出
				dp[j] = maxInTwo(prices[i]+dp[j-1], dp[j])
			} else {
				// 第二次持有股票: 今天持有 vs 之前已经持有
				dp[j] = maxInTwo(-prices[i]+dp[j-1], dp[j])
			}
		}
	}

	return maxInNums(dp)
}
