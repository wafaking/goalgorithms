package array

import (
	"sort"
)

// 重复数字的全排列（leetcode-47）
// 给定一个可包含重复数字的序列nums，按任意顺序，返回所有不重复的全排列。
// 示例1： 输入：nums = [1,1,2], 输出： [[1,1,2], [1,2,1], [2,1,1]]
// 示例2： 输入：nums = [1,2,3]，输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// permuteUnique11 深度遍历，有重复遍历
func permuteUnique11(nums []int) [][]int {
	var (
		res     [][]int
		dfs     func(nums []int, path []int)
		visited = make(map[int]bool, len(nums)) // 记录全局i位置是否遍历过
	)
	sort.Ints(nums) // 先排序

	dfs = func(nums []int, path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int(nil), path...))
			return
		}

		var m = make(map[int]struct{}) // 记录此次循环，相同元素是否出现过
		for i := 0; i < len(nums); i++ {
			// 因每次循环的都是nums，因此用visited记录i是否已遍历过
			if v, ok := visited[i]; ok && v {
				continue
			}
			// nums[i]是否在此次遍历时已出现过，如：1 1 2，1出现过，当i=1时就不应该再出现1
			if _, ok := m[nums[i]]; ok {
				continue
			}

			m[nums[i]] = struct{}{}
			visited[i] = true
			path = append(path, nums[i])
			dfs(nums, path)
			visited[i] = false
			path = path[:len(path)-1]
		}
		return
	}
	dfs(nums, []int{})
	return res
}

// permuteUnique12 深度遍历，有重复遍历(同上）
func permuteUnique12(nums []int) [][]int {
	var (
		res     [][]int
		visited = make(map[int]bool, len(nums)) // 标记i位置是否标记过，也可以使用数组
		dfs     = func(path []int) {}
	)
	sort.Ints(nums) // 先排序
	dfs = func(path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}

		for i, v := range nums {
			// 1：visited[i]条件为真表示i位置的数已使用过
			// 2：如1,1,2,取2,1(首位1)和2,1(第二位1)的visited值为[1,0,1]和[0,1,1]
			if visited[i] || (i > 0 && nums[i] == nums[i-1] && !visited[i-1]) {
				continue
			}
			path = append(path, v)
			visited[i] = true
			dfs(path)
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs([]int{})
	return res
}

// permuteUnique2 深度遍历，每次把遍历过的元素去除，效率更高
func permuteUnique2(nums []int) [][]int {
	var (
		res [][]int
		dfs func(nums []int, path []int)
	)

	sort.Ints(nums) // 先排序
	dfs = func(nums []int, path []int) {
		if len(nums) == 0 {
			var temp = make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		// 包含-1等数，因此最多为19个数字，但索引最长为19，因此富余一位
		// 也可以用map[int]struct来代替更好理解些
		var used = [20]int{} // 每一层记录，同一层不使用重复的数
		for i := 0; i < len(nums); i++ {
			var cur = nums[i]
			if used[cur+10] == 1 {
				continue
			}

			path = append(path, cur)
			nums = append(nums[:i], nums[i+1:]...)
			used[cur+10] = 1
			dfs(nums, path)
			nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
			path = path[:len(path)-1]
		}

	}
	dfs(nums, []int{})
	return res
}
