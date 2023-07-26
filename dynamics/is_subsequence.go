package dynamics

import "strings"

//判断子序列(leetcode-392)
//给定字符串s和t，判断s是否为t的子序列。
//例如，"ace"是"abcde"的一个子序列，而"aec"不是
//进阶：如果有大量输入的S，称作S1,S2,...,Sk其中k>=10亿，你需要依次检查它们是否为T的子序列。在这种情况下，你会怎样改变代码？

//示例1：输入：s="abc",t="ahbgdc"输出：true
//示例2：输入：s="axc",t="ahbgdc"输出：false

// 动态规划+二维数组
func isSubsequence1(s string, t string) bool {
	// dp[i][j]表示前s串的前i个字符是否是t串的前j个字符的子序列
	var dp = make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(t)+1)
	}
	for i := range dp[0] {
		dp[0][i] = true
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

// TODO 动态规划+滚动数组
func isSubsequence2(s string, t string) bool {
	return false
}

// 双指针
func isSubsequence3(s string, t string) bool {
	if s == "" {
		return true
	}
	if len(s) > len(t) {
		return false
	}

	var i = 0
	for j := 0; i < len(s) && j < len(t); {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	if i == len(s) {
		return true
	}

	return false
}

// 找出s中每个字符是否在t中
func isSubsequence4(s string, t string) bool {
	if s == "" {
		return true
	}
	if len(s) > len(t) {
		return false
	}

	var lastIdx = 0 // 用于记录在t中找到s[i]的位置
	for i := 0; i < len(s); i++ {
		// 从t的lastIdx位置开始查找
		idx := strings.IndexByte(t[lastIdx:], s[i])
		if idx == -1 {
			return false
		}
		lastIdx += idx + 1
		// 剩余长度比较
		if len(s[i+1:]) > len(t[lastIdx:]) {
			return false
		}
	}

	return true
}
