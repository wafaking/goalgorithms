package str

import (
	"sort"
)

// 字符串的排列(sword-38)
// 输入一个字符串，打印出该字符串中字符的所有排列。
// 示例: 输入：s = "abc", 输出：["abc","acb","bac","bca","cab","cba"]

// permutation1 回溯，全排列(剔除已选择的数据)
func permutation1(s string) []string {
	// 1. 将字符串转换为数组
	var (
		ans []string
		sli []string
	)

	for _, v := range s {
		sli = append(sli, string(v))
	}

	// 2. 排序，去重
	sort.Strings(sli)

	// 3. 回溯
	var (
		path      = ""
		backtrack func([]string)
	)

	backtrack = func(sli []string) {
		if len(sli) == 0 {
			ans = append(ans, path)
			return
		}
		for i := 0; i < len(sli); i++ {
			if i > 0 && i < len(sli) && sli[i] == sli[i-1] {
				continue
			}
			var cur = sli[i]
			path += cur
			sli = append(sli[:i], sli[i+1:]...)
			backtrack(sli)
			sli = append(sli[:i], append([]string{cur}, sli[i:]...)...)
			// 注：如果有中文，使用len(path)-1会出错,因此使用len(path)-len(cur)
			path = path[:len(path)-len(cur)]
		}
	}
	backtrack(sli)
	return ans
}

// permutation2 回溯，全排列()
func permutation2(s string) []string {
	var sli []string
	for _, v := range s {
		sli = append(sli, string(v))
	}

	sort.Strings(sli)

	var (
		ans       []string
		path      = ""
		visited   = make([]bool, len(sli))
		backtrace func()
	)

	backtrace = func() {
		// 注：如果有中文，使用len(path)==len(sli)永远不相等
		if len(path) == len(s) {
			ans = append(ans, path)
			return
		}

		for i := 0; i < len(sli); i++ {
			if visited[i] || (i > 0 && !visited[i-1] && sli[i] == sli[i-1]) {
				continue
			}
			path += sli[i]
			visited[i] = true
			backtrace()
			visited[i] = false
			// 注：如果有中文，使用len(path)-1会出错,因此使用len(path)-len(sli[i]))
			path = path[:len(path)-len(sli[i])]
		}
	}
	backtrace()
	return ans
}

// permutation3 回溯（交换位置）
func permutation3(s string) []string {
	var (
		res   []string
		bytes = []byte(s)
		dfs   func(idx int)
	)
	dfs = func(idx int) {
		if len(bytes)-1 == idx {
			res = append(res, string(bytes))
			return
		}
		var dict = make(map[byte]bool)
		for i := idx; i < len(bytes); i++ {
			if dict[bytes[i]] {
				continue
			}
			dict[bytes[i]] = true
			bytes[idx], bytes[i] = bytes[i], bytes[idx]
			dfs(idx + 1)
			bytes[idx], bytes[i] = bytes[i], bytes[idx]
		}
	}
	dfs(0)
	return res
}

func permutation4(s string) []string {
	var (
		mp     = make(map[string]int)
		pre    = make([]string, 0)
		result = []string{string(s[0])}
	)

	for i := 1; i < len(s); i++ {
		for _, str := range result {
			ll := len(str)
			for j := 0; j <= ll; j++ {
				temp := str[:j] + string(s[i]) + str[j:]
				if _, OK := mp[temp]; !OK {
					mp[temp]++
					pre = append(pre, temp)
				}
			}
		}
		result, pre = pre, make([]string, 0)
	}
	return result
}

func permutation5(s string) []string {
	res := make(map[string]bool)
	c := []byte(s)
	dfs(res, c, 0)
	RES := []string{}
	for key := range res {
		RES = append(RES, key)
	}
	return RES
}

func dfs(res map[string]bool, c []byte, i int) {
	if i == len(c)-1 {
		res[string(c)] = true
		return
	} else {
		for j := i; j < len(c); j++ {
			c[i], c[j] = c[j], c[i]
			dfs(res, c, i+1)
			c[i], c[j] = c[j], c[i]
		}
	}
}
