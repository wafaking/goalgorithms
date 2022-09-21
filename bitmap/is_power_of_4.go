package bitmap

//4的幂(leetcode-342)
//给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。
//整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4x
//示例 1： 输入：n = 16 输出：true
//示例 2： 输入：n = 5 输出：false
//示例 3： 输入：n = 1 输出：true

// isPowerOfFour1 判断对4取模是否为0且整除后结果为1
func isPowerOfFour1(n int) bool {
	for n > 0 && n%4 == 0 {
		n /= 4
	}
	return n == 1
}

// isPowerOfFour2 法二：判断是不是2的倍数，再判断是否能被3取模得1
func isPowerOfFour2(n int) bool {
	if !(n > 0 && n&(n-1) == 0) {
		return false
	}
	return n%3 == 1
}

// isPowerOfFour3 二进制的4的幂只有一个1且都在奇数位
func isPowerOfFour3(n int) bool {
	// 0000 0001 1  2^0
	// 0000 0100 4  2^2
	// 0001 0000 16 2^4
	// 0100 0000 64 2^6
	// 32位数字奇数位全为1，偶数位全为0，与运算，结果大于1则是4的倍数
	// 0101 0101 0101 0101 0101 0101 0101 0101 32位数字，即0x55555555
	// 或构造数字即0xAAAAAAAA与运算结果为0
	// 1010	1010 1010 1010 1010 1010 1010 1010
	//if !(n > 0 && n&(n-1) == 0) {
	//	return false
	//}
	//return (n & 0xAAAAAAAA) == 0
	return n > 0 && (n&(n-1) == 0) && (n&0x55555555) > 0
}
