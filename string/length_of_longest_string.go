package str

import "goalgorithms/common"

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
		maxLength = common.MaxInTwo(maxLength, len(m))
	}
	return maxLength
}

// 滑动窗口，使用map，单次遍历
func lengthOfLongestSubstring2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var (
		m         = make(map[byte]int)
		maxLength int
		l, r      = 0, 0
	)

	for ; r < len(s); r++ {
		pos, ok := m[s[r]]
		if !ok {
			m[s[r]] = r
			continue
		}

		// 更新最大长度
		maxLength = common.MaxInTwo(maxLength, len(s[l:r]))
		m[s[r]] = r // 更新已存在元素的下标
		// 更新滑动窗口左边界：左侧边界取较大值，防止再回退
		l = common.MaxInTwo(pos+1, l)
	}

	maxLength = common.MaxInTwo(maxLength, len(s[l:r]))
	return maxLength
}

func lengthOfLongestSubstring3(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
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
		ans = common.MaxInTwo(ans, rk-i+1)
	}
	return ans
}
