package str

//各位相加(leetcode-258)
//给定一个非负整数num，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。
//示例1:输入:num=38输出:2解释:各位相加的过程为：
//	38-->3+8-->11
//	11-->1+1-->2
//	由于2是一位数，所以返回2。
//示例2:输入:num=0输出:0

// 递归
func addDigits1(num int) int {
	var f func(num int) int
	f = func(num int) int {
		if num < 10 {
			return num
		}
		var ans int
		for num != 0 {
			ans += num % 10
			num /= 10
		}
		return f(ans)
	}
	return f(num)
}

// 循环
func addDigits2(num int) int {
	if num < 10 {
		return num
	}
	var ans int
	for {
		ans += num % 10
		num /= 10
		if num == 0 {
			if ans < 10 {
				break
			}
			num = ans
			ans = 0
		}
	}
	return ans
}

// TODO
// 数学
func addDigits3(num int) int {
	return (num-1)%9 + 1
}
