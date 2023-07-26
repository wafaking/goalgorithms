package dynamics

//判断子
//给定字符串s和t，判断s是否为t的子串,例如，"bcd"是"abcde"的一个子序列，而"abd"不是;

// 动态规划+二维数组
func isSubstring1(s string, t string) bool {
	// dp[i][j]表示前s串的前i个字符是否是t串的前j个字符的子串
	var dp = make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(t)+1)
	}
	for i := range dp[0] {
		dp[0][i] = true
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if s[i-1] == t[j-1] && dp[i-1][j-1] {
				dp[i][j] = true
				// 注：最后一行只要有一个为真，即为真
				if i == len(dp)-1 {
					return true
				}
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

// TODO 动态规划+滚动数组
func isSubstring2(s string, t string) bool {
	return false
}
