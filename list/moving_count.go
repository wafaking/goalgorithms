package list

// 机器人的运动范围(sword2-13)
// 地上有一个m行n列的方格，从坐标[0,0]到坐标[m-1,n-1]。一个机器人从坐标[0, 0]的格子开始移动，
// 它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
// 例如，当k为18时，机器人能够进入方格[35, 37] 因为3+5+3+7=18。但它不能进入方格[35,38]，因为3+5+3+8=19。
// 请问该机器人能够到达多少个格子？

// 如输入：m = 2, n = 3, k = 1,输出：3;
// 如输入：m = 3, n = 1, k = 0,输出：1

/*
	1, 2, 3, 4
	5, 6, 7, 8
	9, 3, 2, 1
*/

func movingCount(m int, n int, k int) int {
	// m行n列
	if m == 0 && n == 0 {
		return 0
	}

	var visited = make([][]bool, m) // 已遍历过的元素
	var dfs func(i, j int) int
	for k, _ := range visited {
		visited[k] = make([]bool, n)
	}
	dfs = func(i, j int) int {
		// 当前不满足条件即退出
		if i < 0 || j < 0 || i >= m || j >= n || visited[i][j] || (getWordSum(i)+getWordSum(j)) > k {
			return 0
		}
		// log.Printf("i,j: %d, %d", i, j)
		visited[i][j] = true
		return 1 + dfs(i+1, j) + dfs(i, j+1) // 只遍历右侧和下侧相邻坐标即可
	}
	return dfs(0, 0)
}

func getWordSum(i int) (sum int) {
	for ; i > 0; i = i / 10 {
		sum += i % 10
	}
	return sum
}
