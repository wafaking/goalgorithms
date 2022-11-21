package str

// 数值的整数次方(leetcode-50/sword-16)
//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，x^n）。不得使用库函数，同时不需要考虑大数问题。
//示例 1： 输入：x = 2.00000, n = 10 输出：1024.00000
//示例 2： 输入：x = 2.10000, n = 3 输出：9.26100
//示例 3： 输入：x = 2.00000, n = -2 输出：0.25000,解释：2-2 = 1/22 = 1/4 = 0.25

// myPow11 循环(当n较大时，时长较长)
func myPow11(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		x = 1 / x
		n = -n
	}

	var res = x
	for i := 2; i <= n; i++ {
		res *= x
	}
	return res
}

// myPow12 循环(将n拆解成2->4->8)
func myPow12(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	var (
		res  float64 = 1
		temp         = x
	)
	for n > 0 {
		if n%2 == 1 {
			res *= temp
		}
		temp *= temp
		n /= 2
	}
	return res
}

// myPow13 递归，将n拆解
func myPow13(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	var (
		res  float64 = 1
		temp         = x
		dfs  func()
	)
	dfs = func() {
		if n == 0 {
			return
		}
		if n%2 == 1 {
			res *= temp
		}
		temp *= temp
		n /= 2
		dfs()
	}
	dfs()
	return res
}
