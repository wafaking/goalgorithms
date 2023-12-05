package array

//螺旋矩阵(leetcode-54)
//给你一个m行n列的矩阵matrix，请按照顺时针螺旋顺序，返回矩阵中的所有元素。
//示例1：输入：matrix=[[1,2,3],[4,5,6],[7,8,9]]输出：[1,2,3,6,9,8,7,4,5]
//示例2：输入：matrix=[[1,2,3,4],[5,6,7,8],[9,10,11,12]]输出：[1,2,3,4,8,12,11,10,9,5,6,7]

// 按层模拟
func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		top, bottom = 0, len(matrix)
		left, right = 0, len(matrix[0])
		ans         = make([]int, 0, bottom*right)
	)
	for left < right && top < bottom {
		// 上一
		for i := left; i < right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++
		if top >= bottom {
			break
		}

		// 右1
		for i := top; i < bottom; i++ {
			ans = append(ans, matrix[i][right-1])
		}
		right--
		if left >= right {
			break
		}

		// 下一
		for i := right - 1; i >= left; i-- {
			ans = append(ans, matrix[bottom-1][i])
		}
		bottom--
		if top >= bottom {
			break
		}

		// 左1
		for i := bottom - 1; i >= top; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if left >= right {
			break
		}
	}
	return ans
}

// 按层模拟(同法1)
func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}

// 模拟
func spiralOrder3(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total          = rows * columns
		order          = make([]int, total)
		row, column    = 0, 0
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directionIndex = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}
