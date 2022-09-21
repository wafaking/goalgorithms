package bitmap

// 3的幂(leetcode-326)
//给定一个整数，判断它是否是3的幂次方。
//整数n是3的幂次方需满足：存在整数x使得 n == 3x
//示例 2： 输入：n = 0 输出：false
//示例 3： 输入：n = 9 输出：true
//示例 4： 输入：n = 45 输出：false

func isPowerOfThree1(n int) bool {
	for n > 0 { // 12
		if n == 1 {
			return true
		}
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return false
}

func isPowerOfThree11(n int) bool {
	for n > 0 && n%3 == 0 { // 12
		n /= 3
	}
	return n == 1
}

// isPowerOfThree2 每次以2的倍数递增
func isPowerOfThree2(n int) bool {
	for temp := 1; temp <= n; temp *= 3 {
		if temp == n {
			return true
		}
	}
	return false
}

// isPowerOfThree3 判断是否为最大3的幂的约数
func isPowerOfThree3(n int) bool {
	return n > 0 && 1162261467%n == 0
}
