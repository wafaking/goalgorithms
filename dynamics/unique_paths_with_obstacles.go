package dynamics

// 不同路径II(leetcode-63)
// 机器人位于mxn网格的左上角, 机器人每次只能向下或向右移动一步。
// 网格中有障碍物,那么从左上角到右下角将会有多少条不同的路径？
// 网格中的障碍物和空位置分别用1和0来表示。
// 示例1：输入：obstacleGrid=[[0,0,0],[0,1,0],[0,0,0]],输出：2;
// 解释：3x3 网格的正中间有一个障碍物。 从左上角到右下角一共有2条不同的路径：
//	1. 向右 -> 向右 -> 向下 -> 向下
//	2. 向下 -> 向下 -> 向右 -> 向右
//示例 2： 输入：obstacleGrid=[[0,1],[0,0]],输出：1

// 动态规划
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {

	var (
		m   = len(obstacleGrid)
		n   = len(obstacleGrid[0])
		ans = make([][]int, m)
	)

	// 1. 初始化第一行及第一列
	for i := 0; i < m; i++ { // 初始化列
		temp := make([]int, n)
		ans[i] = temp
	}
	for i := 0; i < n; i++ { // 初始化第一行
		if obstacleGrid[0][i] == 1 { // 第一行只要有一个障碍物，后面就无法到达
			break
		}
		ans[0][i] = 1
	}
	for i := 0; i < m; i++ { // 初始化第一列
		if obstacleGrid[i][0] == 1 { // 第一列只要有一个障碍物，后面就无法到达
			break
		}
		ans[i][0] = 1
	}

	for i := 1; i < m; i++ { // row
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			ans[i][j] = ans[i-1][j] + ans[i][j-1]
		}
	}
	return ans[m-1][n-1]
}

// 动态规划，动态数组
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {

	var (
		m   = len(obstacleGrid)
		n   = len(obstacleGrid[0])
		cur = make([]int, n)
		//注：next在循环外初始化时，循环利用时需要重新赋值
		//next = make([]int, n)
	)

	for i := 0; i < n; i++ { // 初始化第一行
		if obstacleGrid[0][i] == 1 { // 第一行只要有一个障碍物，后面就无法到达
			break
		}
		cur[i] = 1
	}

	var hasObstacle = cur[0] == 0
	for i := 1; i < m; i++ { // row
		next := make([]int, n)
		if obstacleGrid[i][0] == 1 {
			// next[0] = 0
			hasObstacle = true
		} else {
			if !hasObstacle {
				next[0] = 1
			}
		}

		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				// next[j] = 0
				continue
			}
			next[j] = next[j-1] + cur[j]
		}
		cur = next[:]
	}
	return cur[n-1]
}

// 动态规划，使用一个数组
func uniquePathsWithObstacles3(obstacleGrid [][]int) int {
	var (
		m   = len(obstacleGrid)
		n   = len(obstacleGrid[0])
		ans = make([]int, n)
	)

	if obstacleGrid[0][0] == 0 {
		ans[0] = 1 // 第一行的第一个元素
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				ans[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 { // 左侧不能是障碍物
				ans[j] += ans[j-1]
			}
		}
	}

	return ans[n-1]
}
