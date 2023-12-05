package array

//螺旋矩阵II(leetcode-59)
//给你一个正整数n，生成一个包含1到n^2所有元素，且元素按顺时针顺序螺旋排列的nxn正方形矩阵matrix。
//示例1：输入：n=3输出：[[1,2,3],[8,9,4],[7,6,5]]
//示例2：输入：n=1输出：[[1]]

// 按层模拟
func generateMatrix1(n int) [][]int {
	var ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	var (
		top, bottom = 0, n - 1
		left, right = 0, n - 1
		total       = 1
	)

	for left <= right && top <= bottom {

		// 上一
		for i := left; i <= right; i++ {
			ans[top][i] = total
			total++
		}
		top++

		// 右1
		for i := top; i <= bottom; i++ {
			ans[i][right] = total
			total++
		}
		right--

		if left <= right && top <= bottom {
			// 下一
			for i := right; i >= left; i-- {
				ans[bottom][i] = total
				total++
			}
			bottom--

			// 左1
			for i := bottom; i >= top; i-- {
				ans[i][left] = total
				total++
			}
			left++
		}
	}

	return ans
}

// todo 模拟
func generateMatrix2(n int) [][]int {
	type pair struct{ x, y int }
	var (
		matrix = make([][]int, n)
		dirs   = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上
	)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	row, col, dirIdx := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := dirs[dirIdx]
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIdx = (dirIdx + 1) % 4 // 顺时针旋转至下一个方向
			dir = dirs[dirIdx]
		}
		row += dir.x
		col += dir.y
	}
	return matrix
}
