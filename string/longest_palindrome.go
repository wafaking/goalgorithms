package str

// 最长回文子串(leetcode-5)
//给你一个字符串 s，找到 s 中最长的回文子串。

//示例 1：
//输入：s = "babad", 输出："bab"
//解释："aba"同样是符合题意的答案。

//示例2:
//输入：s = "cbbd",输出："bb"

func longestPalindrome1(s string) string {
	if len(s) <= 1 {
		return s
	}

	// "abcdcba"
	var res = s[:1]
	for i := 0; i < len(s)-1; i++ {
		for j := i + 2; j <= len(s); j++ {
			if len(s[i:]) <= len(res) {
				return res
			}

			var str = s[i:j]
			if isSymmetry(str) {
				if len(s[i:j]) > len(res) {
					res = s[i:j]
				}
			}
		}
	}
	return res
}

func isSymmetry(str string) bool {
	if len(str) <= 1 {
		return true
	}
	for start, end := 0, len(str)-1; start < end; start, end = start+1, end-1 {
		if str[start] != str[end] {
			return false
		}
	}
	return true
}

func longestPalindrome2(s string) string {
	if len(s) <= 1 {
		return s
	}

	// "abcdcba"

	//b abc
	//c bcd abcdcbac
	//

	for i := 0; i < len(s); i++ {

	}
	var res = s[:1]
	for i := 0; i < len(s)-1; i++ {
		for j := i + 2; j <= len(s); j++ {
			if len(s[i:]) <= len(res) {
				return res
			}

			var str = s[i:j]
			if isSymmetry(str) {
				if len(s[i:j]) > len(res) {
					res = s[i:j]
				}
			}
		}
	}
	return res
}

func judge(s string, start, end int) {

}
