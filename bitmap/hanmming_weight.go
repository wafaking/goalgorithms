package bitmap

// 二进制中1的个数(sword2-15)
// 编写一个函数，输入是一个无符号整数（以二进制串的形式），
// 返回其二进制表达式中数字位数为'1'的个数（也被称为汉明重量).）。

// 如输入：n=11 (控制台输入00000000000000000000000000001011), 输出：3
// 如输入：n=128 (控制台输入 00000000000000000000000010000000), 输出：1
// 如输入：n=4294967293 (控制台输入 11111111111111111111111111111101，部分语言中 n = -3),输出：31

// 左移比较(注：不能右移，右移对负数不适合)
func hammingWeight1(num uint32) int {
	var pivot uint32 = 1
	var count int

	for pivot > 0 { // 移动31位后pivot=0
		if num&pivot != 0 {
			count++
		}
		pivot = pivot << 1
	}
	return count
}

// 利用位数运算 n&n-1可以把最低位的1变为0的特性
// 如：6&(6-1)=4, 其中6(110),4(100)把6的最低位1变成0
func hammingWeight2(num uint32) int {
	var count int
	for ; num > 0; num &= num - 1 {
		count++
	}

	return count
}
