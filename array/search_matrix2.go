package array

//搜索二维矩阵II(leetcode-240/sword-04)
//在一个n*m的二维数组中，行按照从左到右递增排序，列按照从上到下递增排序。
//请判断整数n是否在数组中。
//示例1：输入：matrix=[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],target=5输出：true
//示例2：输入：matrix=[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],target=20输出：false

func searchMatrix21(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}

	for row, col := 0, len(matrix[0])-1; row < len(matrix) && col >= 0; {
		num := matrix[row][col]
		if target == num { // 目标值相等
			return true
		} else if target > num { // 目标值大，下移
			row++
			continue
		} else { // 目标值小，左移
			col--
		}
	}

	return false
}
