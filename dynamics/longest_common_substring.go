package dynamics

import (
	"goalgorithms/common"
)

// 最长公共子串
// 给定两个字符串text1和text2，返回这两个字符串的最长公共子串的长度。如果不存在公共子串，返回0。
// 一个字符串的子串是指这样一个新的字符串：它是取原字符串的一段连续字符串或全部;
// 子串（Substring）是串的一个连续的部分,子序列（Subsequence）则是从不改变序列的顺序，而从序列中去掉任意的元素而获得的新序列；
// 简单说：子串的字符必须连续，子序列LCS则不必。比如字符串acdfg同akdfc的最长公共子串为df，而他们的最长公共子序列是adf。

// 示例1：输入：text1="abcde",text2="ace",输出：1(最长公共子串是"a或c或e"，长度为1)。
// 示例2：输入：text1="abc",text2="abc",输出：3(最长公共子串是"abc"，长度为3)。
// 示例3：输入：text1="abc",text2="bcdef",输出：2(最长公共子串是"bc"，返回2)。
// 示例4：输入：text1="abc",text2="def",输出：0(没有公共子串，返回0)。

// 动态规划
func longestCommonSubstring1(text1 string, text2 string) int {
	// dp[i][j]表示text1的前i个字符与text2的前j个字符的最长公共子串的长度
	// 此处用text1表示纵坐标，text2表示横坐标
	var (
		dp        = make([][]int, len(text1)+1)
		maxLength int // 记录最大长度(最大长度子串可能在前面)
	)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	// 第一行为空字符与text2的公共子串，全为0,第一列同理
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if text1[i-1] == text2[j-1] { // 匹配
				dp[i][j] = dp[i-1][j-1] + 1 // 取上一个匹配数量加1
				maxLength = common.MaxInTwo(maxLength, dp[i][j])
			}
		}
	}

	return maxLength
}
