package array

//搜索二维矩阵(leetcode-74)
//给你一个满足下述两条属性的mxn整数矩阵：
//	每行中的整数从左到右按非严格递增顺序排列。
//	每行的第一个整数大于前一行的最后一个整数。
//给你一个整数target，如果target在矩阵中，返回true；否则，返回false。
//示例1：输入：matrix=[[1,3,5,7],[10,11,16,20],[23,30,34,60]],target=3输出：true
//示例2：输入：matrix=[[1,3,5,7],[10,11,16,20],[23,30,34,60]],target=13输出：false

func searchMatrix11(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for row, col := 0, len(matrix[0])-1; row < len(matrix) && col >= 0; {
		num := matrix[row][col]
		if num == target {
			return true
		} else if target > num {
			row++
		} else {
			col--
		}
	}
	return false
}
