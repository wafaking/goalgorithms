package str

import (
	"goalgorithms/common"
	"strings"
)

//反转字符串II(leetcode-541)
//给定一个字符串s和一个整数k，从字符串开头算起，每计数至2k个字符，就反转这2k字符中的前k个字符。
//如果剩余字符少于k个，则将剩余字符全部反转。
//如果剩余字符小于2k但大于或等于k个，则反转前k个字符，其余字符保持原样。
//示例1：输入：s="abcdefg",k=2,输出："bacdfeg"
//示例2：输入：s="abcd",k=2,输出："bacd"

// 法一：将字符串切分放入数组，然后分别翻转再组合
func reverseStr1(s string, k int) string {
	if k < 0 {
		return s
	}

	//"abcdefg"
	if len(s) <= k {
		return common.Reverse(s)
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
		return common.Reverse(str[:k]) + str[k:]
	}
	return common.Reverse(str)
}

// 法二：将字符串切分放入数组，然后分别翻转再组合
func reverseStr2(s string, k int) string {
	if k < 0 {
		return s
	}
	if len(s) <= k {
		return common.Reverse(s)
	}

	//"abcdefg"
	var res string
	for s != "" {
		if len(s) <= k {
			res += common.Reverse(s)
			s = ""
		} else if len(s) > k && len(s) <= 2*k {
			res += common.Reverse(s[:k]) + s[k:]
			s = ""
		} else { // len(s)>2*k
			res += common.Reverse(s[:k]) + s[k:2*k]
			s = s[2*k:]
		}
	}

	return res
}

func reverseStr3(s string, k int) string {
	var (
		str   = []byte(s)
		n     = len(s)
		cnt   int
		start int

		reverse = func(start, end int) { // 左闭右闭
			for i := start; i <= (end-start)>>1+start; i++ {
				str[i], str[end-i+start] = str[end-i+start], str[i]
			}
		}
	)

	for i := 0; i < n; i++ {
		cnt++
		if cnt == k { // 每k个翻转一次
			reverse(start, start+k-1)
		} else if cnt == 2*k {
			cnt = 0
			start += 2 * k
		}
	}

	//abcdefg
	if (n - start) >= k {
		return string(str)
	} else {
		reverse(start, n-1)
	}

	return string(str)
}

// 模拟：反转每个下标从2k的倍数开始的，长度为k的子串。若该子串长度不足k，则反转整个子串;
func reverseStr4(s string, k int) string {
	t := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		sub := t[i:common.MinInTwo(i+k, len(s))]
		for j, n := 0, len(sub); j < n/2; j++ {
			sub[j], sub[n-1-j] = sub[n-1-j], sub[j]
		}
	}
	return string(t)
}
