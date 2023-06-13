package array

// N皇后(leetcode-51/面试题08.12)
// 按照国际象棋规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子,
// 但n皇后问题研究的是如何将n个皇后放置在n×n的棋盘上，并且使皇后彼此之间不能相互攻击,
// 即：每个皇后都不同行、不同列，也不在对角线上,
// 注：“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。
// 每一种解法包含一个不同的n皇后问题的棋子放置方案，该方案中'Q'和'.'分别代表了皇后和空位。

// 输入：n = 1
// 输出：[["Q"]]
// 输入：4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]

// solveNQueens1 基于集合的回溯
func solveNQueens1(n int) [][]string {
	var (
		ans            = make([][]string, 0)
		queens         = make([]int, n)       // 用于标明x行y列有皇后, like:[1,3,0,2]:第一行的皇后在1列
		yRowMap        = make(map[int]int, 0) // 记录y列皇后所在行
		mDiagonalLeft  = make(map[int]struct{})
		mDiagonalRight = make(map[int]struct{})
		dfs            func(x int)
		genBoad        func()
	)

	for k := range queens { // 初始化都为-1,即每行都没有皇后
		queens[k] = -1
	}

	genBoad = func() {
		var board []string
		for _, v := range queens {
			var path = make([]byte, 0, n)
			for i := 0; i < n; i++ {
				if i == v {
					path = append(path, 'Q')
				} else {
					path = append(path, '.')
				}
			}
			board = append(board, string(path))
		}
		ans = append(ans, board)
	}

	dfs = func(x int) {
		if x == n {
			genBoad()
			return
		}

		/*
		   0 1 0 0
		   0 0 0 1
		   1 0 0 0
		   0 0 1 0
		*/

		// 遍历每一列，看哪一列可设置皇后
		for y := 0; y < n; y++ {

			if _, ok := yRowMap[y]; ok {
				continue
			}

			diagonal1 := y - x // 右对角线的y-x值都相等
			if _, ok := mDiagonalRight[diagonal1]; ok {
				continue
			}

			diagonal2 := x + y // 左对角线的x+y值都相等
			if _, ok := mDiagonalLeft[diagonal2]; ok {
				continue
			}

			queens[x] = y
			yRowMap[y] = x
			mDiagonalRight[diagonal1] = struct{}{}
			mDiagonalLeft[diagonal2] = struct{}{}
			dfs(x + 1)
			queens[x] = -1
			delete(yRowMap, y)
			delete(mDiagonalRight, diagonal1)
			delete(mDiagonalLeft, diagonal2)

		}

	}

	dfs(0)
	return ans
}

// solveNQueens2 (TODO 基于位运算的回溯)
func solveNQueens2(n int) [][]string {
	return nil
}
