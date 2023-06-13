package dynamics

// 整数拆分(leetcode-343)
// 给定正整数n，将其拆分为k个正整数的和（k>=2），并使这些整数的乘积最大化。返回此最大乘积。
// 示例1:输入:n=2,输出:1,解释:2=1+1,1×1=1。
// 示例2:输入:n=10,输出:36,解释:10=3+3+4,3×3×4=36。

// 总结归纳 循环
func integerBreak1(n int) int {
	if n <= 0 {
		return n
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	var res = 1
	mod := n % 3

	switch mod {
	case 0:
		cnt := n / 3
		for i := 0; i < cnt; i++ {
			res *= 3
		}
	case 1:
		left := n - 4
		for i := 0; i < left/3; i++ {
			res *= 3
		}
		res *= 4

	case 2:
		for i := 0; i < (n-2)/3; i++ {
			res *= 3
		}
		res *= 2
	}
	return res
}
