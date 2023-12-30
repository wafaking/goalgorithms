package dynamics

import "math/big"

//不同路径(leetcode-62)
//一个机器人位于一个mxn网格的左上角,机器人每次只能向下或向右移动一步。
//问机器人达到网格的右下角,总共有多少条不同的路径？
//示例1：输入：m=3,n=7,输出：28;
//示例2：输入：m=3,n=2,输出：3;
//示例3：输入：m=7,n=3,输出：28;
//示例4：输入：m=3,n=3,输出：6

// uniquePaths 递归,时间长，易超时
func uniquePaths1(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	} else if m < 0 || n < 0 {
		return 0
	}

	return uniquePaths1(m-1, n) + uniquePaths1(m, n-1)
}

// uniquePaths 动态规划
func uniquePaths2(m int, n int) int {
	// 解题步骤
	// 1. 确定dp数组（dp table）以及下标的含义；
	//	dp[i][j]:表示从(0,0)出发到(i,j)有dp[i][j]条路径；
	// 2. 确定递推公式；
	//  要求dp[i][j]= dp[i-1][j]+dp[i][j-1],即由这两个路径过来的
	// 3. dp数组如何初始化；
	//   因为从(0, 0)的位置到(i, 0)的路径只有一条，因此dp[i][0]=1,同理dp[0][j]=1。
	// 4. 确定遍历顺序；
	//   因为dp[0][0],dp[0][1],dp[1][0]是有数值的，因此可以从左上到右下遍历
	// 5. 举例推导...

	var dp = make([][]int, m)
	for i := range dp {
		temp := make([]int, n)
		temp[0] = 1
		dp[i] = temp
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 动态规划，循环利用空间
func uniquePaths3(m int, n int) int {
	cur, next := make([]int, n), make([]int, n) // 当前行、下一行
	for i := 0; i < n; i++ {
		cur[i] = 1
	}

	for i := 1; i < m; i++ {
		next[0] = 1 // 每行的第一个元素为1
		for j := 1; j < n; j++ {
			next[j] = next[j-1] + cur[j] // 当前=左+上
		}
		cur = next[:] // 将下一行作为当前行
	}
	return cur[n-1]
}

// 使用排列组合
// 从左上角到右下角的过程中，我们需要移动m+n−2次，其中有m−1次向下移动，n−1次向右移动。因此路径的总数，就等于从
// m+n−2次移动中选择m−1次向下移动的方案数，即组合数Cnm.
func uniquePaths4(m int, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
