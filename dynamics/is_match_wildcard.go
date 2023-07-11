package dynamics

// 通配符匹配(leetcode-44)
// 给定输入字符串s和一个字符模式p，请实现一个支持'?'和'*'匹配规则的通配符匹配：
// '?'可以匹配任何单个字符。'*'可以匹配任意字符序列（包括空字符序列）。
// 判定匹配成功的充要条件是：字符模式必须能够完全匹配输入字符串（而不是部分匹配）。
// 示例1：输入：s="aa",p="a"输出：false解释："a"无法匹配"aa"整个字符串。
// 示例2：输入：s="aa",p="*"输出：true解释：'*'可以匹配任意字符串。
// 示例3：输入：s="cb",p="?a"输出：false解释：'?'可以匹配'c',但第二个'a'无法匹配'b'。

func isMatchWildCard(s string, p string) bool {
	// dp[i][j]表示s中前i个字符与模式串中前j个字符匹配的结果
	var dp = make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	// 第一行为空字符与模式串匹配结果
	// 第一列为字符串与空模式串匹配结果
	// 处理第一行匹配
	for i := 1; i < len(dp[0]); i++ {
		if p[i-1] == '*' {
			// 如果模式串为*,则取上一个字符匹配结果
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			// 相同或模式串字符为?时，取上一个字符匹配结果
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			// 不匹配
			if p[j-1] == '*' {
				if j == 1 { // *匹配作意0个或多个字符
					dp[i][j] = true
					continue
				}
				// 表示
				// dp[i][j-1]表示前一个字符与模式串匹配,如a*与a,引处*表示空字符
				// dp[i-1][j]表示当前字符之前的字符与模式串匹配，如：a*与abc，此处*可以匹配任意一个或多个字符
				if dp[i][j-1] || dp[i-1][j] {
					dp[i][j] = true
				}
			}

		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
