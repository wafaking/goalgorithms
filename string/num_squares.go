package str

import (
	"goalgorithms/common"
	"math"
)

//完全平方数(leetcode-279)
//给你一个整数n，返回和为n的完全平方数的最少数量。
//完全平方数是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9和16都是完全平方数，而3和11不是。
//示例1：输入：n=12输出：3解释：12=4+4+4
//示例2：输入：n=13输出：2解释：13=4+9

// 动态规划(可以简化成一维dp)
func numSquares1(n int) int {
	var (
		// dp[i][j]表示和为j时需要最少数量
		dp   = make([][]int, 0)
		nums = make([]int, 0)
		rows int
	)
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	for i := 1; i <= n/2+1; i++ {
		if i*i <= n {
			nums = append(nums, i*i)
			continue
		}
		break
	}
	rows = len(nums)
	dp = make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < rows; i++ {
		for j := 1; j <= n; j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = common.MinInTwo(dp[i-1][j], dp[i][j-nums[i]]+1)
				//dp[i][j] = dp[i][j-nums[i-1]] + 1
			}
		}
	}

	return dp[rows-1][n]
}

// 数学定理(四平方和定理)
//当且仅当n≠4^k×(8m+7)n时，n可以被表示为至多三个正整数的平方和。因此，当n=4^k×(8m+7)时，n只能被表示为四个正整数的平方和。
//当n≠4^k×(8m+7)时，答案只会是1,2,3中的一个：
//	答案为1时，则n为完全平方数；
//	答案为2时，则有n=a^2+b^2，只需要枚举所有的a(1≤a≤n)，判断n−a^2是否为完全平方数即可；
//	答案为3时，我们很难在一个优秀的时间复杂度内解决它，但我们只需要检查答案为111或222的两种情况，即可利用排除法确定答案。
//四平方和定理证明了任意一个正整数都可以被表示为至多四个正整数的平方和。
func numSquares2(n int) int {
	var (
		m    = make(map[int]bool, 0)
		temp int
	)
	if n <= 0 {
		return -1
	}
	// 满足n=4^k×(8m+7)，则为4个
	for temp = n; temp%4 == 0; temp /= 4 {
	}
	if (temp-7)%8 == 0 {
		return 4
	}

	temp = 0
	for i := 1; i <= n; i++ {
		temp = i * i
		m[temp] = true
		if temp > n {
			break
		} else if temp == n {
			return 1
		}
	}
	for k := range m {
		if m[n-k] {
			return 2
		}
	}
	return 3
}

// 数学(同上)
func numSquares3(n int) int {

	var (
		// 判断是否为完全平方数
		isPerfectSquare = func(x int) bool {
			y := int(math.Sqrt(float64(x)))
			return y*y == x
		}

		// 判断是否能表示为 4^k*(8m+7)
		checkAnswer4 = func(x int) bool {
			for x%4 == 0 {
				x /= 4
			}
			return x%8 == 7
		}
	)

	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}
