package dynamics

//目标和(leetcode-494)
//给你一个整数数组nums和一个整数target,向数组中的每个整数前添加'+'或'-'，然后串联起所有整数，可以构造一个表达式：
//例如，nums=[2,1]，可以在2之前添加'+'，在1之前添加'-'，然后串联起来得到表达式"+2-1"。
//返回可以通过上述方法构造的、运算结果等于target的不同表达式的数目。

//示例1：输入：nums=[1,1,1,1,1],target=3,输出：5
//解释：一共有5种方法让最终目标和为3。
//	-1+1+1+1+1=3
//	+1-1+1+1+1=3
//	+1+1-1+1+1=3
//	+1+1+1-1+1=3
//	+1+1+1+1-1=3
//示例2：输入：nums=[1],target=1,输出：1;

// 动态规划+拆分成两个数组
func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum < target {
		return 0
	}

	//设添加+的数的和为positive，添加-号的数和为negative，则有：
	//	positive + negative  = sum
	//	positive - negative = target
	//即：positive=(sum-target)/2
	//因此问题又转换成了分割等和子集问题：求和为positive的数组的组合方案数
	//参考：leetcode-416(canPartition)

	// dp[i][j]表示前i个元素和为j的方案数
	var dp = make([][]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]int, (sum-target)>>1+1)
	}
	// 初始化数组：dp[0][0]=1表示0个元素和为0有1种方案
	dp[0][0] = 1
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if j == 0 {
				dp[i][0] = 1
				continue
			}
			if nums[i-1] < j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-nums[i-1]]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
