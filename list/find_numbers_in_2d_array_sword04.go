package list

//剑指 Offer 04. 二维数组中的查找

//在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
//请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

//示例:
//现有矩阵 matrix 如下：
//[
//	[1,   4,  7, 11, 15],
//	[2,   5,  8, 12, 19],
//	[3,   6,  9, 16, 22],
//	[10, 13, 14, 17, 24],
//	[18, 21, 23, 26, 30]
//]
//给定 target = 5，返回 true。
//给定 target = 20，返回 false。
//
//限制：
//0 <= n <= 1000
//0 <= m <= 1000

func findNumberIn2DArray(matrix [][]int, target int) bool {
	var find bool

	//m := len(matrix) // m：行数
	if len(matrix) == 0 {
		return find
	}

	n := len(matrix[0]) - 1 // n：列数
	if n < 0 {
		return find
	}

	for m := 0; m < len(matrix) && n >= 0; {

		idx := matrix[m][n]
		if target == idx { // 目标值相等
			find = true
			return find
		} else if target > idx { // 目标值大，下移
			m++
			continue
		} else { // 目标值小，左移
			n--
		}
	}

	return find
}
