package str

import (
	"strconv"
	"strings"
)

//字符串相加(leetcode-415)
//求给定两个字符串形式的非负整数num1和num2的和;

//示例 1：
//输入：num1 = "11", num2 = "123"
//输出："134"

//示例 2：
//输入：num1 = "456", num2 = "77"
//输出："533"

//示例 3：
//输入：num1 = "0", num2 = "0"
//输出："0"

func addStrings(num1 string, num2 string) string {
	// 1. 特殊情况
	for _, v := range num1 {
		if v-'0' < 0 || v-'9' > 9 {
			return ""
		}
	}
	for _, v := range num2 {
		if v-'0' < 0 || v-'9' > 9 {
			return ""
		}
	}

	// 2. 拼凑长度使得m=n
	if len(num1) > len(num2) {
		num2 = strings.Repeat("0", len(num1)-len(num2)) + num2
	} else {
		num1 = strings.Repeat("0", len(num2)-len(num1)) + num1
	}

	// 3.求和并拼接
	var (
		level uint8
		res   string
	)
	for i := len(num1) - 1; i >= 0; i-- {
		temp := num1[i] - '0' + num2[i] - '0' + level
		if temp > 9 {
			level = 1
			temp %= 10
		} else {
			level = 0
		}
		res = strconv.FormatUint(uint64(temp), 10) + res
	}
	if level == 1 {
		res = "1" + res
	}

	return res
}

func addStrings2(num1 string, num2 string) string {
	var sign uint8
	var res string
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || sign > 0; i, j = i-1, j-1 {
		var x, y uint8
		if i >= 0 {
			x = num1[i] - '0'
			if x > 9 || x < 0 {
				return ""
			}
		}
		if j >= 0 {
			y = num2[j] - '0'
			if y > 9 || y < 0 {
				return ""
			}
		}
		temp := x + y + sign
		if temp > 9 {
			sign = 1
			temp %= 10
		} else {
			sign = 0
		}
		res = strconv.Itoa(int(temp)) + res
	}
	return res
}
