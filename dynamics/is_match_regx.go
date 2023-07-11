package dynamics

// 正则表达式匹配(leetcode-10)
// 给你一个字符串s和一个字符规律p，请你来实现一个支持'.'和'*'的正则表达式匹配。
// '.'匹配任意单个字符,'*'匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖整个字符串s的，而不是部分字符串。

// 示例1：输入：s="aa",p="a"输出：false解释："a"无法匹配"aa"整个字符串。
// 示例2:输入：s="aa",p="a*"输出：true
// 解释：因为'*'代表可以匹配零个或多个前面的那一个元素,在这里前面的元素就是'a'。因此，字符串"aa"可被视为'a'重复了一次。
// 示例3：输入：s="ab",p=".*"输出：true解释：".*"表示可匹配零个或多个（'*'）任意字符（'.'）。

// 动态规划
func isMatchRegx1(s string, p string) bool {
	//dp[i][j]表示前i个字符与前j个模式串的匹配结果
	var dp = make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	// 空字符与空模式匹配
	dp[0][0] = true

	// 第一行为空字符串与模式匹配(除空字符串外，其余全不匹配, 不考虑字符串中有空格情况)
	// 第一列为空模式与字符串匹配
	for i := 1; i < len(dp[0]); i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1] || dp[0][i-2]
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' { // 匹配
				dp[i][j] = dp[i-1][j-1] // 取上一个字符串匹配结果
				continue
			}

			// 不匹配，判断模式串中是否有特殊字符
			if p[j-1] == '*' { // 当前模式串是否是*
				if j == 1 { // 即j=1，如：*与a, 一定不匹配
					continue
				}

				// j-2>=0的情况
				if dp[i][j-2] {
					// 如果当前字符串已与*前两个字符串匹配，则当前字符串一定匹配，如：abaC*与aba
					dp[i][j] = true
				} else if dp[i-1][j] == true && (p[j-2] == s[i-1] || p[j-2] == '.') {
					// 如果上一个字符匹配，那么当前字符与上一个字符相同或上一个字符为.时，匹配；
					// 如：1、ab*与ab、a.*与ab；其中b与*不相同，但字符串上一个字符a已匹配，且模式串中*前面是.或b与当前b匹配，因此匹配
					dp[i][j] = true
				}
			}

		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
