package dynamics

// (0/1背包问题)有n件物品和一个最多能背重量为w的背包。第i件物品的重量是weight[i]，得到的价值是value[i]。
// 每件物品只能用一次，求解将哪些物品装入背包里物品价值总和最大。

// 法一：使用动态规划, 一一将每一步最大价值记录下来
func weightBag11(weight, value []int, bagWeight int) int {
	var dp = make([][]int, len(weight)+1)
	// dp[i][j]表示当背包我还是为j时前i价物品的最大价值
	// 初始化：共有len(value)+1行,有bagWeight+1列
	// 第一行全为0，表示当有0种物品可选时，背包可容纳的最大价值为0
	// 第一列全为0，肤浅当背包可承重为0时，背包可容纳的最大价值为0
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}

	// i表示第i行，第i行物品的我重量为weight[i-1],价值为value[i-1]
	for i := 1; i < len(value)+1; i++ {
		for j := 1; j < bagWeight+1; j++ { // j取代表背包的承重量
			// 每列的背包重量为i,每行的重量为weight[j],价值为value[j-1]
			if j < weight[i-1] {
				// 当前背包小于对应的物品重量，则取前i-1件物品的最大价值dp[i-1][j]
				dp[i][j] = dp[i-1][j]
			} else {
				// 当背包承重量大于第i件物品的重量时，则最大价值取下面两个较大值：
				// 		1.前i-1件物品的最大价值dp[i-1][j]；
				// 		2.第i件物品的价值与选用第i件物品后剩余承重对应的最大价值:value[i-1]+ dp[i-1][j-weight[i-1]]
				//			其中：j-weight[i-1]表示背包剩余承重量，dp[i-1][j-weight[i-1]]表示剩余承重量对应的最大价值.
				//				即dp上一行，第j-weight[i-1]列
				// value[i-weight]表示
				// weight[i-1]表示第j件商品的重量
				pre := dp[i-1][j]                            // 前i-1件物品的最大价值dp[i-1][j]；
				curValue := value[i-1]                       // 第i件物品的价值
				leftWeight := j - weight[i-1]                // 剩余重量
				leftValue := dp[i-1][leftWeight]             // 剩余重量对应的最大价值
				dp[i][j] = maxInTwo(pre, curValue+leftValue) // 两者取最大价值
			}
		}
	}
	return dp[len(weight)][bagWeight]
}

// TODO 使用动态规划+滚动数组
func weightBag12(weight, value []int, bagWeight int) int {
	return 0
}
