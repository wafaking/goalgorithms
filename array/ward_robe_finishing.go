package array

//衣橱整理(leetcode-LCR130)
//家居整理师将待整理衣橱划分为mxn的二维矩阵grid，其中grid[i][j]代表一个需要整理的格子。整理师自grid[0][0]开始逐行逐列地整理每个格子。
//整理规则为：在整理过程中，可以选择向右移动一格或向下移动一格，但不能移动到衣柜之外。
//	同时，不需要整理digit(i)+digit(j)>cnt的格子，其中digit(x)表示数字x的各数位之和。
//请返回整理师总共需要整理多少个格子。
//示例1：输入：m=4,n=7,cnt=5输出：18

// 动态规划
func wardrobeFinishing1(m int, n int, cnt int) int {
	var (
		ans int
		// 记录已经处理过的位置
		visited  = make(map[int]struct{})
		dfs      func(m, n int)
		digitSum = func(num int) int {
			sum := 0
			for ; num > 0; num /= 10 {
				sum += num % 10
			}
			return sum
		}
	)
	dfs = func(row, col int) {
		if row >= m || col >= n {
			return
		} else if digitSum(row)+digitSum(col) > cnt {
			return
		} else if _, ok := visited[row*n+col]; ok {
			return
		}
		ans++
		visited[row*n+col] = struct{}{}
		dfs(row+1, col)
		dfs(row, col+1)
	}
	dfs(0, 0)
	return ans
}

// 广度优先搜索(迭代)
func wardrobeFinishing2(m int, n int, cnt int) int {
	var (
		ans      int
		digitSum = func(num int) int {
			sum := 0
			for ; num > 0; num /= 10 {
				sum += num % 10
			}
			return sum
		}
		visited = make(map[int]bool, 0)
	)

	for i := 0; i < n; i++ { // 第一行
		if digitSum(i) > cnt {
			break
		}
		ans++
		visited[i] = true
	}

	for i := 1; i < m; i++ { // 第一列(从第二行开始)
		if digitSum(i) > cnt {
			break
		}
		ans++
		visited[i*n] = true
	}

	for i := 1; i < m; i++ {
		sum1 := digitSum(i)
		// 如当m=16,n=16,cnt=2时，i=0,j=10满足但无法穿过
		// 但cnt=9时,i=1,j=9不满足，但i=1,j=10时满足，因此
		// cur满足条件，同时需要其左侧或上侧满足条件才能到达
		for j := 1; j < n; j++ {
			if sum1+digitSum(j) > cnt {
				continue
			}

			// 左侧和上侧
			if visited[i*n+j-1] || visited[(i-1)*n+j] {
				visited[i*n+j] = true
				ans++
			}
		}
	}
	return ans
}
