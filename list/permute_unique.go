package list

import (
	"sort"
)

// 重复数字的全排列（leetcode-47）
// 给定一个可包含重复数字的序列nums，按任意顺序，返回所有不重复的全排列。
// 示例1： 输入：nums = [1,1,2], 输出： [[1,1,2], [1,2,1], [2,1,1]]
// 示例2： 输入：nums = [1,2,3]，输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// permuteUnique1 深度遍历，有重复遍历
func permuteUnique1(nums []int) [][]int {
	var (
		res     [][]int
		visited = make(map[int]bool, len(nums)) // 标记i位置是否标记过，也可以使用数组
		// visited =
		dfs = func(path []int) {}
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
