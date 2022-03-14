package str

import "strings"

// 反转字符串II(leetcode-541)
//给定一个字符串s和一个整数k，从字符串开头算起，每计数至2k个字符，就反转这2k字符中的前k个字符。

//如果剩余字符少于k个，则将剩余字符全部反转。
//如果剩余字符小于2k但大于或等于k个，则反转前k个字符，其余字符保持原样。

// 示例1： 输入：s = "abcdefg", k = 2, 输出："bacdfeg"
// 示例2： 输入：s = "abcd", k = 2, 输出："bacd"

// 法一：将字符串切分放入数组，然后分别翻转再组合
func reverseStr1(s string, k int) string {
	if k < 0 {
		return s
	}

	//"abcdefg"
	if len(s) <= k {
		return reverse(s)
	}

	var strList []string
	for s != "" {
		if len(s) >= 2*k {
			strList = append(strList, s[:2*k])
			s = s[2*k:]
		} else {
			strList = append(strList, s[:])
			s = s[:0]
		}
	}

	for i, v := range strList {
		strList[i] = reversePrevK(v, k)
	}
	return strings.Join(strList, "")
}

func reversePrevK(str string, k int) string {
	if len(str) >= k {
		return reverse(str[:k]) + str[k:]
	}
	return reverse(str)
}

// 法二：将字符串切分放入数组，然后分别翻转再组合
func reverseStr2(s string, k int) string {
	if k < 0 {
		return s
	}
	if len(s) <= k {
		return reverse(s)
	}

	//"abcdefg"
	var res string
	for s != "" {
		if len(s) <= k {
			res += reverse(s)
			s = ""
		} else if len(s) > k && len(s) <= 2*k {
			res += reverse(s[:k]) + s[k:]
			s = ""
		} else { // len(s)>2*k
			res += reverse(s[:k]) + s[k:2*k]
			s = s[2*k:]
		}
	}

	return res
}
