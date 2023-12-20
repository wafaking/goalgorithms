package array

//组合总和Ⅳ(leetcode-377)
//给定由不同整数组成的数组nums，和目标整数target。请从nums中找出并返回总和为target的元素组合的个数。
//示例1：输入:nums=[1,2,3],target=4，输出：7
//	所有可能的组合为：(1,1,1,1),(1,1,2),(1,2,1),(1,3),(2,1,1),(2,2),(3,1)
//示例2：输入：nums=[9],target=3,输出：0

// 全排列，当nums=[4,2,1],target=32时超时（不推荐）
func combinationSum41(nums []int, target int) int {
	var (
		ans  [][]int
		path []int
		dfs  func(idx int)
	)
	//sort.Ints(nums)
	dfs = func(idx int) {
		if target == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		} else if target < 0 {
			return
		}
		for i := idx; i < len(nums); i++ { // 无重复元素
			//for ; i+1 < len(nums) && nums[i] == nums[i+1]; i++ {
			//	continue
			//}
			target -= nums[i]
			path = append(path, nums[i])
			dfs(idx)
			target += nums[i]
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return len(ans)
}

// combinationSum42 动态规划法，类似背包问题
func combinationSum42(nums []int, target int) int {
	var dp = make([]int, target+1) // 构建目标数组
	dp[0] = 1                      //即当目标值为0时，什么也不取，是有一个值的
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				// 大的目标数由小的的构成
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

//动态规划(二维数组)
//背包问题：这个是是求排列数，而518零钱兑换2题是求组合数
//求组合数就是外层for循环遍历物品，内层for遍历背包，即物品作Y轴，背包作X轴；
//求排列数就是外层for遍历背包，内层for循环遍历物品，即物品作X轴，背包作Y轴；
func combinationSum43(nums []int, target int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	for j := 0; j <= target; j++ {
		for i := 0; i < n; i++ {
			if j < nums[i] {
				dp[i+1][j] = dp[i][j]
				continue
			}
			for k := i; k >= 0; k-- {
				if j >= nums[k] {
					dp[i+1][j] += dp[i+1][j-nums[k]]
				}
			}
		}
	}
	return dp[n][target]
}
