package array

// 46. 全排列
// 返回不重复数字的数组nums的所有可能的全排列(可以按任意顺序返回答案)。
// 如输入：nums = [1,2,3], 则输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 如输入：nums = [0,1], 则输出：[[0,1],[1,0]]
// 如输入：nums = [1], 则输出：[[1]];

// permute11 使用回溯法，因元素都不会重复，因此可不记录是否遍历过
func permute11(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	backTrace(nums, []int{}, &res)
	return res
}

func backTrace(nums []int, path []int, res *[][]int) {
	if len(nums) == 0 { // 结束条件，即path中元素的数量等于给定元素的个数
		var temp = make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}

	for i := 0; i < len(nums); i++ { // 遍历数组，让不同的数字作第一项
		var cur = nums[i]
		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...)                      // 截取元素
		backTrace(nums, path, res)                                  // 将剩余元素继续遍历
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...) // 回溯，将截取掉的元素恢复,注意此处后面的元素的索引
		path = path[:len(path)-1]                                   // 回溯，将path中的元素剔除，恢复到上一个状态
	}
}

// permute12 使用匿名函数，理论上同permute1
func permute12(nums []int) [][]int {
	var (
		res       [][]int
		backTrace func(nums []int, path []int)
	)

	backTrace = func(nums []int, path []int) {
		if len(nums) == 0 {
			//var temp = make([]int, len(path))
			//copy(temp, path)
			//res = append(res, temp)
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			var cur = nums[i]
			path = append(path, cur)
			nums = append(nums[:i], nums[i+1:]...)
			backTrace(nums, path)
			nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
			path = path[:len(path)-1]
		}
		return
	}

	backTrace(nums, []int{})
	return res
}

// permute21 深度优先遍历（回溯）
func permute21(nums []int) [][]int {
	var (
		res     [][]int
		visited = make(map[int]struct{}, len(nums))
		dfs     func(path []int)
	)
	if len(nums) == 0 {
		return res
	}
	dfs = func(path []int) {
		if len(path) == len(nums) {
			var temp = make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}

		for _, v := range nums {
			if _, ok := visited[v]; ok {
				continue
			}
			path = append(path, v)
			visited[v] = struct{}{}   // 记录已遍历的字典
			dfs(path)                 // 深度遍历
			path = path[:len(path)-1] // 回溯，将上一个记录移除
			delete(visited, v)        // 回溯，将记录遍历的字典的记录移除
		}
	}

	dfs([]int{})
	return res
}

// permute22 深度遍历（同上）
func permute22(nums []int) [][]int {
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

// 参考permutation3
func permute3(nums []int) [][]int {
	var (
		res [][]int
		dfs func(idx int)
	)
	dfs = func(idx int) {
		if idx == len(nums)-1 {
			res = append(res, nums)
			return
		}
		var dict = make(map[int]bool)
		for i := idx; i < len(nums); i++ {
			if dict[nums[i]] {
				continue
			}
			nums[idx], nums[i] = nums[i], nums[idx]
			dict[nums[idx]] = true
			dfs(idx + 1)
			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}
	dfs(0)
	return res
}
