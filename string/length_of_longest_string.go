package str

// 无重复字符的最长子串(leetcode-3)
// 给定一个字符串s，请你找出其中不含有重复字符的最长子串的长度。
// 注："pwke"是"pwwkew"子序列，不是子串。

// 示例1: 输入:s="abcabcbb" 输出: 3(abc)
// 示例2: 输入:s="bbbbb" 输出: 1(b)
// 示例3: 输入:s="pwwkew" 输出: 3(wke)

// lengthOfLongestSubstring1 使用map，双层遍历
func lengthOfLongestSubstring1(s string) int {
	var maxLength int
	for i := 0; i < len(s); i++ {
		var m = map[uint8]struct{}{s[i]: {}}
		for j := i + 1; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				break
			}
			m[s[j]] = struct{}{}
		}
		if len(m) > maxLength {
			maxLength = len(m)
		}
	}
	return maxLength
}

// lengthOfLongestSubstring2 使用滑动窗口
func lengthOfLongestSubstring2(s string) int {
	var (
		maxLength int
		m         = make(map[uint8]struct{})
	)
	// i, j 分别表示滑动窗口的左、右边界
	for i, j := 0, 0; i < len(s) && j < len(s); {
		// 右边界没有重复，添加
		if _, ok := m[s[j]]; !ok {
			m[s[j]] = struct{}{}
			j++
			if maxLength < len(m) {
				maxLength = len(m)
			}
			continue
		}
		// 有重复,则从左边界i判断重复的字符
		for {
			if s[i] == s[j] { // 此位置重复，则左、右边界都右移
				i++
				j++
				break
			} else { // 不重复，则左边界右移, 并移除已记录的元素
				delete(m, s[i])
				i++
			}
		}
	}

	return maxLength
}

// lengthOfLongestSubstring21 滑动窗口
func lengthOfLongestSubstring21(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为-1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
