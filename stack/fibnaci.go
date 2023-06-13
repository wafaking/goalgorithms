package stack

import (
	str "goalgorithms/string"
	"math"
	"strconv"
)

// 斐波那契数(leetcode-509)
//斐波那契数形成的序列称为斐波那契数列。该数列由0和1开始，后面的每一项数字都是前面两项数字的和。也就是：
//F(0)=0，F(1)=1, F(n)=F(n-1)+F(n-2)，其中n>1,给定n，请计算F(n)。

// 示例1：输入：n=2,输出:1,解释：F(2)=F(1)+F(0)=1+0=1
// 示例2：输入：n=3, 输出:2,解释：F(3)=F(2)+F(1)=1+1=2
// 示例3：输入：n=4, 输出:3,解释：F(4)=F(3)+F(2)=2+1=3

// 1. 使用降序，依次将大的数值分解成小的(大量重复计算)
func fibonacci1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci1(n-1) + fibonacci1(n-2)
}

// 2. 使用升序，避免重复计算
func fibonacci2(n int) int {
	if n <= 1 {
		return n
	}

	var (
		fib0 = 0
		fib1 = 1
		fibN int
	)
	for i := 2; i <= n; i++ {
		fibN = fib0 + fib1
		fib0, fib1 = fib1, fibN
	}
	return fibN
}

//Fibonacci3 3. 使用升序，避免重复计算(可以解决n过大时溢出的问题)
func Fibonacci3(n int) string {
	if n == 0 {
		return "0"
	} else if n == 1 {
		return "1"
	}

	var (
		fib0     int64 = 0
		fib1     int64 = 1
		fibN     int64
		fs0, fs1 string
		res      string
		i        = 2
	)

	for i <= n {
		if math.MaxInt64-fib1 <= fib0 {
			fs0 = strconv.FormatInt(fib0, 10)
			fs1 = strconv.FormatInt(fib1, 10)
			break
		}
		fibN = fib0 + fib1
		fib0, fib1 = fib1, fibN
		res = strconv.FormatInt(fibN, 10)
		i++
	}

	for i <= n {
		res = str.AddStrings2(fs0, fs1)
		fs0, fs1 = fs1, res
		i++
	}

	return res
}

// 青蛙跳台阶问题(sword 10)
// 一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个n级的台阶总共有多少种跳法。
// 当n=0时，取1；
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
func numWays1(n int) int {
	if n <= 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	var (
		fib0, fib1 = 1, 1
		fibN       int
	)
	for i := 2; i <= n; i++ {
		fibN = (fib0 + fib1) % (1e9 + 7)
		fib0, fib1 = fib1, fibN
	}
	return fibN
}

func numWays2(n int) int {
	var (
		fibA, fibB = 1, 1
	)
	for i := 2; i <= n; i++ {
		fibA, fibB = fibB, (fibA+fibB)%(1e9+7)
	}
	return fibB
}
