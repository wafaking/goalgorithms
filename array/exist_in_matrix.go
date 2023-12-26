package array

//矩阵中的路径(sword12)
//给定一个mxn二维字符网格board和一个字符串单词word。如果word存在于网格中，返回true否则返回false。
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
//同一个单元格内的字母不允许被重复使用。
//例如，在下面的3×4的矩阵中包含单词"ABCCED"（单词中的字母已标出）。
//如：输入board=[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]],word="ABCCED",输出：true;
//如输入：board=[["a","b"],["c","d"]],word="abcd",输出：false;

// 深度优先遍历
func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0]) // n行m列
	if len(word) > m*n {
		return false
	}

	// i为行数，j为列数，k为word字符索引
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		// 终止条件
		if k == len(word) {
			return true
		}

		// 未找到,退出
		if i < 0 || j < 0 || i >= n || j >= m || board[i][j] != word[k] {
			return false
		}

		// 找到了,将此值标为'0',防止下次回找时再找到此位置字符
		board[i][j] = '0'
		// 递归回溯，只要有一个相同就继续找下去
		return dfs(i, j+1, k+1) || dfs(i, j-1, k+1) || dfs(i+1, j, k+1) || dfs(i-1, j, k+1)
	}

	// 深度优先遍历
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
