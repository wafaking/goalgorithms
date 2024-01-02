package dynamics

import "goalgorithms/common"

//最小路径和(leetcode-64)
//给定一个包含非负整数的mxn网格grid，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//说明：每次只能向下或者向右移动一步。
//示例1：输入：grid=[[1,3,1],[1,5,1],[4,2,1]]输出：7
//解释：因为路径1→3→1→1→1的总和最小。
//示例2：输入：grid=[[1,2,3],[4,5,6]]输出：12
//提示：
//	m==grid.length
//	n==grid[i].length
//	1<=m,n<=200
//	0<=grid[i][j]<=200

// 动态规划
func minPathSum1(grid [][]int) int {
	var (
		// 注：题中已说明m,n>0
		m = len(grid)
		n = len(grid[0])
	)
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = common.MinInTwo(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[m-1][n-1]
}
