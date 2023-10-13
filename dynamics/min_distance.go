package dynamics

import "goalgorithms/common"

//编辑距离(leetcode-72)
//给定两个单词word1和word2，请返回将word1转换成word2所使用的最少操作数。
//你可以对一个单词进行如下三种操作：插入一个字符、删除一个字符、替换一个字符

//示例1：输入：word1="horse",word2="ros"输出：3
//解释：
//	horse->rorse(将'h'替换为'r')
//	rorse->rose(删除'r')
//	rose->ros(删除'e')
//示例2：输入：word1="intention",word2="execution"输出：5
//解释：
//	intention->inention(删除't')
//	inention->enention(将'i'替换为'e')
//	enention->exention(将'n'替换为'x')
//	exention->exection(将'n'替换为'c')
//	exection->execution(插入'u')

func minDistance(word1 string, word2 string) int {
	// dp[i][j]表示word1的前i个字符转换成word2的前j个字符的最小操作数
	var dp = make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	// 第一列，即空字符转换成word1需要的最少操作数
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	// 第一行，即空字符转换成word2需要的最少操作数
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if word1[i-1] == word2[j-1] { // 匹配
				// 取word1上一个字符操作数、word1前一个字操作数+1、word2前一个字符操作数+1中的最小值
				dp[i][j] = common.MinInThree(dp[i-1][j-1], dp[i][j-1]+1, dp[i-1][j]+1)
			} else {
				dp[i][j] = common.MinInThree(dp[i-1][j-1]+1, dp[i][j-1]+1, dp[i-1][j]+1)
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
