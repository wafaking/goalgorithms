package str

// 基于KMP算法实现查找字符串子串函数
func strStrV2(haystack, needle string) int {
	// 子串长度=0
	if len(needle) == 0 {
		return 0
	}
	//主串长度=0，或者主串长度小于子串长度
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}
	// 子串长度=1时单独判断
	if len(needle) == 1 {
		i := 0
		for ; i < len(haystack); i++ {
			if haystack[i] == needle[0] {
				return i
			}
		}
		return -1
	}

	// 其他情况走 KMP 算法
	return kmpSearch(haystack, needle)
}

// KMP 算法实现函数
func kmpSearch(s, p string) int {
	var (
		next = generateNext(p)
		i, j = 0, 0 // 主串、子串的起始位置
	)

	for i < len(s) && j < len(p) {
		//①如果j = -1，或者当前字符匹配成功（即S[i] == P[j]），都令i++，j++
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			//②如果j != -1，且当前字符匹配失败（即S[i] != P[j]），则令i不变，j=next[j]
			//next[j]即为j应该返回的位置
			j = next[j]
		}
	}

	if j == len(p) { // 完全匹配，返回下标位置
		return i - j
	}

	return -1
}

// getNext 构造用于可直接获取到移动到子串位置的索引数组
// 即如果不匹配了，模式串移动的位置就
func generateNext(p string) []int {
	var (
		i, j = 0, 1                //前缀、后缀子串起始位置
		next = make([]int, len(p)) //
	)
	next[0] = -1 // 初始为-1
	for j < len(p)-1 {
		// 如果匹配，则i,j都会后移，
		// 如果不匹配，则取i位置的next数组中的值(即i重新应该上次匹配的位置)
		/*
			ABCDABD（后缀） 起始：j=1
			 ABCDABD(前缀) 起始：i=0
			因为是通过最长可匹配前缀子串计算，所以 j 的值最大不会超过 m-1
		*/

		if i == -1 || p[i] == p[j] {
			i++
			j++
			// 此时i即为最大前后缀公共子串长度
			// 回到上一个最长可匹配前缀子串结尾字符下标
			next[j] = i
		} else {
			i = next[i]
		}
	}

	return next
}
