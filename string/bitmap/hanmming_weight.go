package bitmap

// 二进制中1的个数(leetcode-191/sword-15)
// 编写一个函数，输入是一个无符号整数（以二进制串的形式），
// 返回其二进制表达式中数字位数为'1'的个数（也被称为汉明重量).）。

// 如输入：n=11 (控制台输入00000000000000000000000000001011), 输出：3
// 如输入：n=128 (控制台输入 00000000000000000000000010000000), 输出：1
// 如输入：n=4294967293 (控制台输入 11111111111111111111111111111101，部分语言中 n = -3),输出：31

// hammingWeight1 左移比较(注：不能右移，右移对负数不适合)
func hammingWeight1(nums uint32) int {
	var res int
	for i := 0; i < 32; i++ {
		val := nums & uint32(1<<i)
		if val != 0 {
			res++
		}
	}
	return res
}

// 左移比较，同上
func hammingWeight11(num uint32) int {
	var pivot uint32 = 1
	var count int

	for pivot > 0 { // 移动31位后pivot小于0，但pivot是uint32,所以pivot=0
		if num&pivot != 0 {
			count++
		}
		pivot = pivot << 1
	}
	return count
}

// hammingWeight2 右移与1比较，直至为0（对于负数也不适合）
// 负数为正数取反加1，最高位为1，如果右移，则数值位的高端全补1
// 负数-1左移1位
// 1111 1111 -1
// 1111 1110 <<1
// 0000 0010 -2 取反加1

// 负数-1右称1位
// 1111 1111 -1
// 1111 1111 >> 1 最高位仍补1,一直右移，一直补1
// 1000 0001 -1：取反加1
func hammingWeight2(nums uint32) int {
	var res int
	for nums != 0 {
		if nums&1 != 0 {
			res++
		}
		nums = nums >> 1
	}
	return res
}

// 利用位数运算 n&n-1可以把最低位的1变为0的特性
// 如：6&(6-1)=4, 其中6(110),4(100)把6的最低位1变成0
func hammingWeight3(num uint32) int {
	var count int
	for ; num > 0; num &= num - 1 {
		count++
	}

	return count
}
