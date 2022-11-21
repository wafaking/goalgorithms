package array

// 子集(leetcode-78)
// 给定无重复元素的整数数组nums。返回该数组所有可能的子集（幂集）。
// 示例1：输入：nums=[1,2,3], 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 示例2：输入：nums = [0], 输出：[[],[0]]

// subset1 回朔法
func subsets1(nums []int) [][]int {
	var (
		ans  [][]int
		dfs  func(idx int)
		path []int
	)
	dfs = func(idx int) {
		ans = append(ans, append([]int(nil), path...))
		if len(path) == len(nums) {
			return
		}
		for i := idx; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}

	}
	dfs(0)

	return ans
}

// subsets2 迭代法实现子集枚举
func subsets2(nums []int) [][]int {
	var ans [][]int
	// 设置长度为n的位数组，每位都用0/1分别表示在子集中和不在子集中
	// 因此mask可以有2^n-1中情况，则当n=3时，即111表示都在数组中
	for mask := 0; mask < 1<<len(nums); mask++ {
		var path []int
		for i := range nums { // 注：此处i的取值为[0,len(nums))
			// 或者 ((mask>>i)&1) >0
			if 1<<i&mask > 0 {
				path = append(path, nums[i])
			}
		}
		ans = append(ans, append([]int{}, path...))
	}
	return ans
}
