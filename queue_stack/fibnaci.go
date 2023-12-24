package queue_stack

import (
	"math"
	"strconv"
)

//斐波那契数列(leetcode-509)
//斐波那契数形成的序列称为斐波那契数列。该数列由0和1开始，后面的每一项数字都是前面两项数字的和。也就是：
//F(0)=0，F(1)=1,F(n)=F(n-1)+F(n-2)，其中n>1,给定n，请计算F(n)。
//示例1：输入：n=2,输出:1,解释：F(2)=F(1)+F(0)=1+0=1
//示例2：输入：n=3,输出:2,解释：F(3)=F(2)+F(1)=1+1=2
//示例3：输入：n=4,输出:3,解释：F(4)=F(3)+F(2)=2+1=3

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

// 使用字符串相加(可以解决n过大时溢出的问题)
func bigFibonacci1(n int) string {
	if n == 0 {
		return "0"
	} else if n == 1 {
		return "1"
	}

	var (
		str1, str2    = "0", "1"
		ans           string
		addTwoStrings func(num1 string, num2 string) string
	)

	addTwoStrings = func(num1 string, num2 string) string {
		var (
			ans    string
			sign   uint8
			n1, n2 = len(num1) - 1, len(num2) - 1
		)
		for i := 0; i <= n1 || i <= n2 || sign != 0; i++ {
			var a, b uint8
			if n1 >= i {
				a = num1[n1-i] - '0'
			}
			if n2 >= i {
				b = num2[n2-i] - '0'
			}
			sum := sign + a + b
			if sum > 9 {
				sign = 1
				sum %= 10
			} else {
				sign = 0
			}
			ans = strconv.Itoa(int(sum)) + ans
		}
		return ans
	}

	for i := 2; i <= n; i++ {
		ans = addTwoStrings(str1, str2)
		str1, str2 = str2, ans
	}

	return str2
}

// 使用字符串相加,开始时用数值相加，溢出时用字符串相加
func bigFibonacci2(n int) string {
	if n == 0 {
		return "0"
	} else if n == 1 {
		return "1"
	}

	var (
		i                      = 2
		num1, num2, numN int64 = 0, 1, 0
		str1, str2, ans  string
		addTwoStrings    func(num1 string, num2 string) string
	)

	addTwoStrings = func(num1 string, num2 string) string {
		var (
			ans    string
			sign   uint8
			n1, n2 = len(num1) - 1, len(num2) - 1
		)
		for i := 0; i <= n1 || i <= n2 || sign != 0; i++ {
			var a, b uint8
			if n1 >= i {
				a = num1[n1-i] - '0'
			}
			if n2 >= i {
				b = num2[n2-i] - '0'
			}
			sum := sign + a + b
			if sum > 9 {
				sign = 1
				sum %= 10
			} else {
				sign = 0
			}
			ans = strconv.Itoa(int(sum)) + ans
		}
		return ans
	}

	for ; i <= n; i++ {
		if math.MaxInt64-num2 <= num1 {
			break
		}
		numN = num1 + num2
		num1, num2 = num2, numN
	}

	if i == n {
		return strconv.FormatInt(num2, 10)
	}

	str1 = strconv.FormatInt(num1, 10)
	str2 = strconv.FormatInt(num2, 10)
	for ; i <= n; i++ {
		ans = addTwoStrings(str1, str2)
		str1, str2 = str2, ans
	}

	return str2
}
