package dynamics

// 最长公共子序列(leetcode-1143)
// 给定两个字符串text1和text2，返回这两个字符串的最长公共子序列的长度。如果不存在公共子序列，返回0。
// 一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
// 例如，"ace"是"abcde"的子序列，但"aec"不是"abcde"的子序列。
// 两个字符串的公共子序列是这两个字符串所共同拥有的子序列。

// 示例1：输入：text1="abcde",text2="ace",输出：3(即最长公共子序列是"ace"，长度为3)。
// 示例2：输入：text1="abc",text2="abc",输出：3(最长公共子序列是"abc"，长度为3)。
// 示例3：输入：text1="abc",text2="def",输出：0(两个字符串没有公共子序列，返回0)。

// 动态规划
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j]表示text1的前i个字符与text2的前j个字符的最长公共子串的长度
	// 此处用text1表示纵坐标，text2表示横坐标
	var dp = make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	// 第一行为空字符与text2的公共子串，全为0,第一列同理
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if text1[i-1] == text2[j-1] { // 匹配
				dp[i][j] = dp[i-1][j-1] + 1 // 取上一个匹配数量加1
			} else { // 不匹配
				// 则取max(text1当前字符与text2上一个字符匹配数量,text1上一个字符与text2当前字符匹配数量)
				dp[i][j] = maxInTwo(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
