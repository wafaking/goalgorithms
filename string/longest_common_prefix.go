package str

import (
	"goalgorithms/common"
	"math"
)

//最长公共前缀(leetcode-14)
//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。
//示例 1： 输入：strs = ["flower","flow","flight"] 输出："fl"
//示例 2： 输入：strs = ["dog","racecar","car"] 输出：""
//解释：输入不存在公共前缀。
//注： strs[i] 仅由小写英文字母组成

// 纵向扫描(扫描每个字符串的相同位置字符)
func longestCommonPrefix1(strs []string) string {
	var (
		ans string
		pos int
		n   = len(strs)
	)

	if n == 0 {
		goto endLabel
	}

	for i := pos; i < math.MaxInt16; i++ {
		for j := 0; j < n; j++ {
			if len(strs[j]) <= i {
				goto endLabel
			}
		}

		for j := 0; j < n; j++ {
			if strs[j][i] != strs[0][i] {
				goto endLabel
			}
		}
		ans += string(strs[0][i])
	}

endLabel:
	return ans
}

// 横向扫描(LCS=LCS(LCS(str1,str2),str3))
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var (
		initStr = strs[0]
		lcp     = func(str1, str2 string) string {
			var ans string
			for i := 0; i < len(str1) && i < len(str2); i++ {
				if str1[i] == str2[i] {
					ans += string(str1[i])
				} else {
					break
				}
			}
			return ans
		}
	)

	for i := 1; i < len(strs); i++ {
		initStr = lcp(initStr, strs[i])
	}
	return initStr
}

// 分治(TODO)
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)
		minLength := common.MinInTwo(len(lcpLeft), len(lcpRight))
		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[:minLength]
	}
	return lcp(0, len(strs)-1)
}

// 二分查找(TODO)
func longestCommonPrefix4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	low, high := 0, minLength
	for low < high {
		mid := (high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}
