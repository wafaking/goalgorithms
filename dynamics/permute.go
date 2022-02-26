package dynamics

// 46. 全排列
// 返回不重复数字的数组nums的所有可能的全排列(可以按任意顺序返回答案)。

// 如输入：nums = [1,2,3], 则输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 如输入：nums = [0,1], 则输出：[[0,1],[1,0]]
// 如输入：nums = [1], 则输出：[[1]];

// 深度优先遍历（回溯）
func permute1(nums []int) [][]int {
	var (
		res      [][]int
		addedMap = make(map[int]bool) // 用于记录是否遍历过
		dfs      func(path []int)
	)

	if len(nums) == 0 {
		return res
	}

	dfs = func(path []int) {
		if len(path) == len(nums) { // 已完成遍历，将此添加到结果集
			// 注：此处要用副本添加，path底层会变化
			var temp = make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		// [[1 2 4 3] [1 2 4 3] [1 3 4 2] [1 3 4 2]
		for _, v := range nums {
			if addedMap[v] { // 遍历过
				continue
			}

			// 未遍历过
			path = append(path, v) // 将此元素添加到路径中
			addedMap[v] = true
			dfs(path) // 深度遍历

			// 此轮结束，向上回溯
			path = path[:len(path)-1] // 注：向上回溯时是当前的长度减1
			addedMap[v] = false
		}
	}
	dfs([]int{}) // 起始为空
	return res
}

func permute2(nums []int) [][]int {
	var (
		res     [][]int
		visited = make(map[int]bool)
		dfs     func(path []int)
	)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			res = append(res, path)
			return
		}
		for _, v := range nums {
			if visited[v] {
				continue
			}
			//path = append(path, v) 注：此处返回一个新值
			visited[v] = true
			dfs(append(path, v)) // 注：此处的path未添加新值，传递的是新path
			// 回溯，但不清除路径中的元素
			visited[v] = false
		}
	}
	dfs([]int{})
	return res
}
