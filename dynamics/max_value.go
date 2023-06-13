package dynamics

// 礼物的最大价值(sword-47)
// 在一个m*n的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于0）。
// 你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
// 给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
// 示例1: 输入:
//	[
//		[1,3,1],
//		[1,5,1],
//		[4,2,1]
//	]
// 输出:12,解释:路径1→3→5→2→1可以拿到最多价值的礼物;

// 动态规划，替换原有数组
func maxValue11(grid [][]int) int {
	// 状态：f(i,j)表示第i,j位置能获取到的最大价值
	// 状态转移方程：f(i,j)=max(f(i-1), f(j-1))+nums[i][j]
	if len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}
	var maxInTwo = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = maxInTwo(grid[i][j-1], grid[i-1][j]) + grid[i][j]
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

// 动态规划，滚动数组
func maxValue12(grid [][]int) int {
	// 状态：f(i,j)表示第i,j位置能获取到的最大价值
	// 状态转移方程：f(i,j)=max(f(i-1), f(j-1))+nums[i][j]
	if len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}
	var (
		maxInTwo = func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		pre = make([]int, len(grid[0]))
		cur = make([]int, len(grid[0]))
	)

	// 初始化pre
	for i := 0; i < len(grid[0]); i++ {
		if i == 0 {
			pre[0] = grid[0][0]
			continue
		}
		pre[i] = pre[i-1] + grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j == 0 {
				cur[0] = pre[0] + grid[i][0]
				continue
			}
			cur[j] = maxInTwo(cur[j-1], pre[j]) + grid[i][j]
		}

		copy(pre, cur)
	}

	return pre[len(cur)-1]
}
