package array

//组合总和(leetcode-39)
//给定无重复元素的整数数组candidates和目标整数target，找出candidates中可以使数字和为目标数target的所有不同组合。
//candidates中的同一个数字可以无限制重复被选取。如果至少一个数字的被选数量不同，则两种组合是不同的。
//示例1：输入：candidates=[2,3,6,7],target=7,输出：[[2,2,3],[7]]
//示例2：输入:candidates=[2,3,5],target=8,输出:[[2,2,2,2],[2,3,3],[3,5]]
//示例3：输入:candidates=[2],target=1,输出:[]

// 深度遍历,改变candidates(使用数组,考虑重复数字)
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

// 深度遍历（使用下标）
func combinationSum12(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	var (
		n   = len(candidates)
		ans = make([][]int, 0)
		dfs func(i int, path []int)
	)
	dfs = func(idx int, path []int) {
		//if target < 0 {
		//	return
		//} else if 0 == target {
		if 0 == target {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		for i := idx; i < n; i++ {
			path = append(path, candidates[i])
			target -= candidates[i]
			dfs(i, path)
			target += candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{})
	return ans
}
