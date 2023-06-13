package dynamics

import (
	"math"
	"strconv"
)

// 剪绳子(sword2-14,同leetcode343-整数拆分)
// 一根长度为n的绳子，请把长度为n的绳子剪成整数长度的m段（m、n都是整数，n>1并且m>1），
// 每段绳子的长度记为 k[0],k[1]...k[m-1]。求k[0]*k[1]*...*k[m-1]可能的最大乘积是多少？
// 例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

// 输入: 2 输出: 1, 解释: 2 = 1 + 1, 1 × 1 = 1
// 输入: 10 输出: 36 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36

// 递归算法，自上而下，会有重复计算
func cuttingRope1(n int) int {
	// 先判断特殊情况
	if n <= 1 {
		return 0
	} else if n == 2 { // 此处只能拆分成1+1
		return 1
	} else if n == 3 {
		return 2 // 1+2
	}

	var dfs func(n int) int
	dfs = func(n int) int {
		// 拆分后的最小值
		if n < 2 {
			return 1 // 不能取0，取0后积为0了
		} else if n == 2 { // 此处是2
			return 2
		} else if n == 3 {
			return 3
		}

		res1 := dfs(2) * dfs(n-2)
		res2 := dfs(3) * dfs(n-3)

		if res1 > res2 {
			return res1
		} else {
			return res2
		}
	}
	return dfs(n)
}

// 循环算法，自下而上，不会有重复计算
func cuttingRope2(n int) int {
	if n < 2 {
		return 0
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	// 拆分后的最小值
	var productMap = map[int]int{
		// 0: 0, 不会取0
		1: 1,
		2: 2,
		3: 3,
	}
	for i := 4; i <= n; i++ {
		var max = 0
		for j := 1; j <= i/2; j++ { // 5=2+3、5=3+2,因此取到中间值即可
			res := productMap[j] * productMap[i-j]
			if res > max {
				max = res
			}
		}
		// 每次内部循环得出i的最大值
		productMap[i] = max
	}
	return productMap[n]
}

// 简单算法
func cuttingRope3(n int) int {
	if n < 2 {
		return 0
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2 // 1+2
	}

	// case n>=4
	// n = 4, 2*2
	// n = 5, 2*3
	// n = 6, 3*3
	// n = 7, 3*4
	// n = 8, 3*3*2
	// n = 9, 3*3*3
	// 即n>4时尽可能的拆分成长度为3的段子
	// 最后一节拆分成4的段子

	// 拆分后的最小值
	var (
		productMap = map[int]int{
			0: 1,
			1: 1,
			2: 2,
			3: 3,
			// 4: 4,
		}
		times = n / 3
		mod   = n % 3
		rate  = 1
	)

	if mod == 1 { // 最后一段可以截取成4的情况
		times -= 1
		rate = 4 // 取值4
		// (n - 4)*
	}

	res, _ := strconv.Atoi(strconv.FormatFloat(math.Pow(3, float64(times)), 'f', 0, 64))
	return res * rate * productMap[mod]
}
