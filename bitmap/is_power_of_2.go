package bitmap

import "math/bits"

//判断一个数是否是2的幂(leetcode231)
//如果存在一个整数 x使得n==2^x，则认为n是2的幂次方。
//输入：n = 1, 输出：true
//输入：n = 16,输出：true
//输入：n = 3, 输出：false

// 利用n&(n-1)=0的性质
func isPowerOfTwo1(n int) bool {
	return (n > 0) && (n&(n-1) == 0)
}

// 利用n&(-n)=n的特性
func isPowerOfTwo2(n int) bool {
	return n > 0 && n&(-n) == n
}

// 递归算法(能否被2整除=1)
func isPowerOfTwo3(n int) bool {
	if n < 1 {
		return false
	} else if n == 1 || n == 2 {
		return true
	} else if n%2 != 0 { // 排队奇数和n/2为奇数的情况
		return false
	}
	return isPowerOfTwo3(n / 2)
}

// 利用go语言特性
func isPowerOfTwo4(n int) bool {
	return 1 == bits.OnesCount(uint(n))
}
