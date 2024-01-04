package dynamics

// 分割等和子集(leetcode-416)
// 给定只包含正整数的非空数组nums。请判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
// 示例1：输入：nums=[1,5,11,5],输出：true(数组可以分割成[1,5,5]和[11])。
// 示例2：输入：nums=[1,2,3,5],输出：false(数组不能分割成两个元素和相等的子集)。

// 法一：深度遍历（容易超出时间限制）
func canPartition1(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum&1 > 0 { // 奇数不成立
		return false
	}

	// 每个数值可以选或不选
	//var start int
	var (
		half = sum >> 1
		dfs  func(total, idx int) bool
		temp int
	)

	dfs = func(total, idx int) bool {
		// 不取idx位置的数是否满足条件
		if total == half {
			return true
		} else if total > half {
			return false
		}

		// 超界
		if idx == len(nums) {
			return false
		}

		// 需要取idx位置的数
		temp = total + nums[idx]
		if temp == half {
			return true
		} else if temp > half { // 加上此数超出half值
			return dfs(total, idx+1)
		}

		// 取idx位值值与不取idx位置值
		return dfs(total, idx+1) || dfs(total+nums[idx], idx+1)
	}

	return dfs(0, 0)
}

// 法一：深度遍历（容易超出时间限制）
func canPartition2(nums []int) bool {
	var sum int
	for i := range nums {
		sum += nums[i]
	}
	if sum&1 != 0 { // 奇数不成立
		return false
	}

	var dfs func(idx, sum int) bool
	dfs = func(idx, sum int) bool {
		if idx >= len(nums) {
			if sum == 0 {
				return true
			}
			return false
		}

		// 选择加当前数或减去当前数
		return dfs(idx+1, sum+nums[idx]) || dfs(idx+1, sum-nums[idx])
	}
	return dfs(0, 0)
}

// 法二：使用动态规划
func canPartition3(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum&1 > 0 { // 奇数不成立
		return false
	}

	var (
		half = sum >> 1
		// dp[i][j]表示前i个元素是否可以组合成和为j
		dp = make([][]bool, len(nums)+1)
	)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, half+1)
	}
	dp[0][0] = true

	for i := 1; i < len(dp); i++ {
		for j := 0; j < half+1; j++ {
			// 第一列默认都可以成功，即认为1不选时可以组合成和为0
			if j == 0 {
				dp[i][0] = true
				continue
			}
			if nums[i-1] == j {
				dp[i][j] = true
			} else if nums[i-1] > j {
				// 元素大于要组合的值时，不选用此元素，取不包括此元素的结果
				dp[i][j] = dp[i-1][j]
			} else {
				// 取不选用此元素的结果 || 先用此元素后剩余元素组成的剩余值的结果(即剩余的元素是否能组成j-num)
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

// 法三：动态规划（单个数组）
func canPartition4(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum&1 > 0 { // 奇数不成立
		return false
	}
	if len(nums) < 2 {
		return false
	}

	sum = sum >> 1

	// dp[i]表示i个元素是否可以组合成和为j
	var dp = make([]bool, sum+1)
	for _, num := range nums {
		// 注：此处从后向前推进，防止从前向后推进时出现的当num=1时，i=1,2时都成立的情况
		for i := sum; i > 0; i-- {
			if num == i {
				dp[i] = true
			} else if num < i {
				dp[i] = dp[i] || dp[i-num]
			}
		}
	}

	return dp[sum]
}
