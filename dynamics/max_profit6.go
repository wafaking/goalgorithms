package dynamics

//买卖股票的最佳时机含手续费(leetcode-714)
//给定一个整数数组prices，其中prices[i]表示第i天的股票价格；整数fee代表了交易股票的手续费用。
//你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
//返回获得利润的最大值。注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

//示例1：输入：prices=[1,3,2,8,4,9],fee=2输出：8
//解释：能够达到的最大利润:
//	在此处买入prices[0]=1
//	在此处卖出prices[3]=8
//	在此处买入prices[4]=4
//	在此处卖出prices[5]=9
//总利润:((8-1)-2)+((9-4)-2)=8
//示例2：输入：prices=[1,3,7,5,10,3],fee=3输出：6

// 动态规划
func maxProfit61(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}

	var buy, sell = -prices[0], 0
	for i := 1; i < len(prices); i++ {
		sell = maxInTwo(sell, prices[i]+buy-fee)
		buy = maxInTwo(buy, -prices[i]+sell)
	}

	return sell
}

// 单调函数，取上升段
func maxProfit62(prices []int, fee int) int {
	var n = len(prices)
	if n == 0 {
		return 0
	}

	var (
		buy    = prices[0] + fee // 初始成本价格
		profit int               // 收益
	)
	for i := 1; i < n; i++ {
		if buy > prices[i]+fee {
			buy = prices[i] + fee
			continue
		}

		if prices[i] > buy {
			profit += prices[i] - buy
			buy = prices[i]
		}
	}

	return profit
}
