package str

// 字符串的排列(leetcode-567)
// 给定两个字符串s1和s2，写一个函数来判断s2是否包含s1的排列。
// 换句话说，第一个字符串的排列之一是第二个字符串的子串。
// 示例1： 输入: s1="ab" s2="eidbaooo", 输出:True, 即s2包含s1的排列之一("ba").
// 示例2：输入: s1="ab", s2="eidboaoo", 输出: False

func checkInclusion(s1 string, s2 string) bool {
	// 1. 校验
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}

	//"ab" "eidboaoo"
	// 2. 初始化两个数组，以字母ascii码相对于'a'值为索引，出现的次数为count
	array1 := [26]int{}
	array2 := [26]int{}
	for k, v := range s1 {
		array1[v-'a']++
		v2 := s2[k] // s2也只取两位
		array2[v2-'a']++
	}

	// 先比较一次，如果不相等，使用滑动窗口向右滑动[数组可以比较，切片不可以]
	if array2 == array1 {
		return true
	}

	// 使用滑动窗口，逐渐向右滑动s2，
	//右侧新进入的字母次数+1,左侧滑出的字母次数-1
	for i := m; i < n; i++ { // 从未比较过的地方开始,如d开始（即窗口右侧）
		v2 := s2[i] - 'a' // 即d所对应的ascii值
		array2[v2]++      // d的索引的次数+1

		v2Pre := s2[i-m] - 'a' // 窗口左侧的e对应的ascii码值
		array2[v2Pre]--        // e索引对应的次数-1（窗口左侧）
		if array1 == array2 {
			return true
		}
	}
	return false
}
