package array

// 组合总和(leetcode-39)
// 给定无重复元素的整数数组candidates和目标整数target，找出candidates中可以使数字和为目标数target的所有不同组合。
// candidates中的同一个数字可以无限制重复被选取。如果至少一个数字的被选数量不同，则两种组合是不同的。
// 示例 1： 输入：candidates = [2,3,6,7], target = 7,输出：[[2,2,3],[7]]
// 示例 2： 输入: candidates = [2,3,5], target = 8,输出: [[2,2,2,2],[2,3,3],[3,5]]
// 示例 3： 输入: candidates = [2], target = 1,输出: []

// combinationSum11 深度遍历(使用数组,考虑重复数字)
func combinationSum11(candidates []int, target int) [][]int {
	var (
		ans  [][]int
		path []int
		dfs  func(nums []int, left int)
	)

	dfs = func(nums []int, left int) {
		if left == 0 {
			path = append(path)
			ans = append(ans, append([]int(nil), path...))
			return
		} else if left < 0 {
			return
		}

		for i := 0; i < len(nums); i++ {
			//用于兼容有重复数字，可以不需要
			for i+1 < len(nums) && nums[i] == nums[i+1] {
				i++
			}
			var cur = nums[i]
			left -= cur
			path = append(path, cur)
			dfs(nums[i:], left)
			left += cur
			path = path[:len(path)-1]
		}
	}

	dfs(candidates, target)
	return ans
}

// combinationSum12 深度遍历，不使用数组，使用下标
func combinationSum12(candidates []int, target int) [][]int {
	var (
		ans  [][]int
		path []int
		dfs  func(idx, left int)
	)

	dfs = func(idx, left int) {
		if idx == len(candidates) {
			return
		}
		if left == 0 {
			var temp = make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		} else if left < 0 {
			return
		}

		for i := idx; i < len(candidates); i++ {
			left -= candidates[i]
			path = append(path, candidates[i])
			dfs(i, left)
			path = path[:len(path)-1]
			left += candidates[i]
		}
	}

	dfs(0, target)
	return ans
}

func combinationSum13(candidates []int, target int) (ans [][]int) {
	var (
		dfs  func(target, idx int)
		path []int
	)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			path = append(path, candidates[idx])
			dfs(target-candidates[idx], idx)
			path = path[:len(path)-1]
		}
	}
	dfs(target, 0)
	return
}
