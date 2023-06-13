package dynamics

// 使用最小花费爬楼梯(leetcode-746)
// 给定整数数组cost，其中cost[i]是从楼梯第i个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
// 你可以选择从下标为0或下标为1的台阶开始爬楼梯。请你计算并返回达到楼梯顶部的最低花费。

// 示例1：输入：cost=[10,15,20],输出：15; 解释：
// 你将从下标为1的台阶开始, 支付15，向上爬两个台阶，到达楼梯顶部。
// 总花费为15。

// 示例2：输入：cost=[1,100,1,1,1,100,1,1,100,1],输出：6
// 解释：你将从下标为0的台阶开始。
// -支付1，向上爬两个台阶，到达下标为2的台阶。
// -支付1，向上爬两个台阶，到达下标为4的台阶。
// -支付1，向上爬两个台阶，到达下标为6的台阶。
// -支付1，向上爬一个台阶，到达下标为7的台阶。
// -支付1，向上爬两个台阶，到达下标为9的台阶。
// -支付1，向上爬一个台阶，到达楼梯顶部。
// 总花费为6。

// minCostClimbingStairs1 动态规划
func minCostClimbingStairs1(cost []int) int {
	var (
		// 设楼顶对应的下标为n，则创建dp长度为n+1
		// 由于初始台阶可以从0或1,因此设dp[0]=dp[1]=1，后台加上此台阶的花费即可
		n   = len(cost)
		dp  = make([]int, n+1)
		min func(a, b int) int
	)

	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 2; i < n+1; i++ {
		// dp[i]表示到达第i阶的最小花费
		// min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])表示第i-1花费cost[i-1]到第i阶、第i-2花费cost[2-1]到第i阶的花费的最小值
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2]) // 状态转移方程
	}

	return dp[len(cost)]
}

// minCostClimbingStairs2 动态规划，使用滚动数组，将空间复杂度优化到O(1)
func minCostClimbingStairs2(cost []int) int {
	var (
		n        = len(cost)
		min      func(a, b int) int
		pre, cur int
	)

	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 2; i <= n; i++ {
		pre, cur = cur, min(pre+cost[i-2], cur+cost[i-1])
	}

	return cur
}
