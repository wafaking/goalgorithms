package str

import "unicode"

// 字母大小写全排列(leetcode-784)
// 给定字符串s，通过将字符串s中的每个字母转变大小写，可以获得一个新的字符串，返回所有可能得到的字符串集合。
// 示例1:输入：s="a1b2", 输出：["a1b2", "a1B2", "A1b2", "A1B2"],
// 示例2:输入: s="3z4", 输出: ["3z4","3Z4"]

// 注：1000001, A->65,
// 100000,32=1<<5
// 1100001,a->97
// 因此：a^(1<<5)=A, A^(1<<5)=a
// letterCasePermutation1 使用深度遍历，map去重
func letterCasePermutation1(s string) []string {
	var (
		ans       []string
		used      = make(map[string]struct{})
		bytes     = []byte(s)
		backTrace func(idx int)
	)

	backTrace = func(idx int) {
		for ; idx < len(bytes); idx++ {
			if (bytes[idx] >= 'a' && bytes[idx] <= 'z') ||
				(bytes[idx] >= 'A' && bytes[idx] <= 'Z') {
				break
			}
		}

		if idx >= len(bytes) {
			return
		}

		if _, ok := used[string(bytes)]; !ok {
			ans = append(ans, string(bytes))
			used[string(bytes)] = struct{}{}
		}
		backTrace(idx + 1)

		bytes[idx] = bytes[idx] ^ 32
		if _, ok := used[string(bytes)]; !ok {
			ans = append(ans, string(bytes))
			used[string(bytes)] = struct{}{}
		}
		backTrace(idx + 1)
	}
	ans = append(ans, s)
	used[s] = struct{}{}
	backTrace(0)
	return ans
}

func letterCasePermutation2(s string) []string {
	var (
		ans       = []string{}
		bytes     = []byte(s)
		backTrace func(idx int)
		path      []byte
	)
	backTrace = func(idx int) {
		if idx >= len(bytes) {
			ans = append(ans, string(path))
			return
		}
		path = append(path, bytes[idx])
		backTrace(idx + 1)
		path = path[:len(path)-1]
		if unicode.IsLetter(rune(bytes[idx])) {
			bytes[idx] = bytes[idx] ^ 32
			path = append(path, bytes[idx])
			backTrace(idx + 1)
			path = path[:len(path)-1]
		}

	}
	backTrace(0)
	return ans
}

// letterCasePermutation3 广度优先遍历（TODO）
func letterCasePermutation3(s string) []string {
	var (
		ans []string
		q   = []string{""}
	)
	for len(q) > 0 {
		cur := q[0]
		pos := len(cur)
		if pos == len(s) {
			ans = append(ans, cur)
			q = q[1:]
		} else {
			if unicode.IsLetter(rune(s[pos])) {
				q = append(q, cur+string(s[pos]^32))
			}
			q[0] += string(s[pos])
		}
	}
	return ans
}

// letterCasePermutation4 使用循环，广度优先遍历(先在结果集中添加字母，再在结果集中拼接)
// s = "a1b2"
// [a,A]->[a1,A1]->[a1b,A1b,a1B,A1B]->[a1b2,A1b2,a1B2,A1B2]
func letterCasePermutation4(s string) []string {
	var ans []string
	for idx := 0; idx < len(s); idx++ {
		if len(ans) == 0 { // 先添加一个
			ans = append(ans, string(s[idx]))
			if unicode.IsLetter(rune(s[idx])) { // 字母就多添加一个
				ans = append(ans, string(s[idx]^32))
			}
			continue
		}

		for i, v := range ans {
			ans[i] = v + string(s[idx])
			if unicode.IsLetter(rune(s[idx])) {
				ans = append(ans, v+string(s[idx]^32))
			}
		}
	}
	return ans
}

// letterCasePermutation5 深度遍历（推荐）
func letterCasePermutation5(s string) (ans []string) {
	var (
		dfs func([]byte, int)
	)
	dfs = func(s []byte, pos int) {
		for pos < len(s) && unicode.IsDigit(rune(s[pos])) {
			pos++
		}
		if pos == len(s) {
			ans = append(ans, string(s))
			return
		}
		dfs(s, pos+1) // s[pos]保持原样，深度遍历
		s[pos] ^= 32  // s[pos]由a->A
		dfs(s, pos+1)
		s[pos] ^= 32 // s[pos]恢复
	}
	dfs([]byte(s), 0)
	return
}

// letterCasePermutation6 使用位数组枚举
//假设字符串s有m个字母，那么全排列就有2^m个字符串序列，且可以用位掩码bits唯一地表示一个字符串。
//bits的第i为0表示字符串s中从左往右第i个字母为小写形式；
//bits的第i为1表示字符串s中从左往右第i个字母为大写形式；
//我们采用的位掩码只计算字符串s中的字母，对于数字则直接跳过，
//通过位图计算从而构造正确的全排列。我们依次检测字符串第i个字符串c：
//如果字符串c为数字，则我们直接在当前的序列中添加字符串c；
//如果字符串c为字母，且c为字符串中的第k个字母，如果掩码bits中的第k位为0，
//则添加字符串c的小写形式；如果掩码bits中的第k位为1，则添加字符串c的大写形式；
func letterCasePermutation6(s string) (ans []string) {
	m := 0
	for _, c := range s {
		if unicode.IsLetter(c) {
			m++
		}
	}
	for mask := 0; mask < 1<<m; mask++ {
		t, k := []rune(s), 0
		for i, c := range t {
			if unicode.IsLetter(c) {
				if mask>>k&1 > 0 {
					t[i] = unicode.ToUpper(c)
				} else {
					t[i] = unicode.ToLower(c)
				}
				k++
			}
		}
		ans = append(ans, string(t))
	}
	return
}
