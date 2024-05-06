package str

// 字符串的排列(leetcode-567)
// 给定两个字符串s1和s2，写一个函数来判断s2是否包含s1的排列。
// 换句话说，第一个字符串的排列之一是第二个字符串的子串。
// 示例1：输入: s1="ab" s2="eidbaooo", 输出:True, 即s2包含s1的排列之一("ba").
// 示例2：输入: s1="ab", s2="eidboaoo", 输出: False
// 注：s1和s2仅包含小写字母

// 滑动窗口+bitmap
func checkInclusion1(s1 string, s2 string) bool {
	var (
		n1      = len(s1)
		n2      = len(s2)
		bitMap1 = [26]byte{}
		bitMap2 = [26]byte{}
	)

	if n2 < n1 {
		return false
	}

	for i := 0; i < n1; i++ {
		bitMap1[s1[i]-'a']++
		bitMap2[s2[i]-'a']++
	}
	if bitMap1 == bitMap2 {
		return true
	}
	// 窗口从n1开始，添加右侧的，移除左侧的
	for i := n1; i < n2; i++ {
		bitMap2[s2[i]-'a']++
		bitMap2[s2[i-n1]-'a']--
		if bitMap1 == bitMap2 {
			return true
		}
	}
	return false
}

// 滑动窗口优化
func checkInclusion21(s1, s2 string) bool {
	var (
		m, n   = len(s1), len(s2)
		bitMap = [26]int{} // 记录差异的字符数量
		diff   int         // 记录差异数量
	)
	if m > n {
		return false
	}

	for i := 0; i < m; i++ {
		bitMap[s1[i]-'a']--
		bitMap[s2[i]-'a']++
	}

	// 计算差异字符数量
	for i := range bitMap {
		v := bitMap[i]
		if v == 0 {
			continue
		} else if v > 0 {
			diff += v
		} else {
			diff -= v
		}
	}

	if diff == 0 {
		return true
	}
	for i := m; i < n; i++ {
		var x, y = s2[i-m] - 'a', s2[i] - 'a'
		if x == y {
			continue
		}

		bitMap[y]++
		if bitMap[y] <= 0 {
			diff--
		} else {
			diff++
		}

		bitMap[x]--
		if bitMap[x] >= 0 {
			diff--
		} else {
			diff++
		}
		if diff == 0 {
			return true
		}
	}

	return false
}

// 滑动窗口优化（todo)
func checkInclusion22(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for i, ch := range s1 {
		cnt[ch-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	for _, c := range cnt[:] {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := n; i < m; i++ {
		x, y := s2[i]-'a', s2[i-n]-'a'
		if x == y {
			continue
		}
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}

// 双指针（TODO）
func checkInclusion3(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]uint8{}
	for _, ch := range s1 {
		cnt[ch-'a']--
	}
	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		cnt[x]++
		for cnt[x] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}
