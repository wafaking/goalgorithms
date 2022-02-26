package str

import (
	"fmt"
	"sort"
)

//剑指 Offer 38. 字符串的排列
//输入一个字符串，打印出该字符串中字符的所有排列。
//你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

//示例:
//输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]

func permutation(s string) []string {
	return nil
}

func permutation1(s string) []string {
	res := []string{}
	bytes := []byte(s)
	var dfs func(x int)
	dfs = func(x int) {
		if x == len(bytes)-1 {
			res = append(res, string(bytes))
		}
		dict := map[byte]bool{}
		for i := x; i < len(bytes); i++ {
			if !dict[bytes[i]] {
				bytes[x], bytes[i] = bytes[i], bytes[x]
				dict[bytes[x]] = true
				dfs(x + 1)
				bytes[x], bytes[i] = bytes[i], bytes[x]
			}
		}
	}
	dfs(0)
	return res
}

func permutation2(s string) (ans []string) {
	newStr := make([]string, len(s))
	for k, v := range s {
		newStr[k] = string(v)
	}
	sort.Strings(newStr)
	n := len(s)
	perm := ""
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, perm)
		}
		for i, v := range newStr {
			if vis[i] || i > 0 && !vis[i-1] && v == newStr[i-1] {
				continue
			}
			perm += v
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

func permutation3(s string) []string {
	length := len(s)
	mp := make(map[string]int)
	pre := make([]string, 0)
	result := []string{string(s[0])}

	for i := 1; i < length; i++ {
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

func permutation4(s string) []string {
	res := make(map[string]bool)
	c := []byte(s)
	fmt.Println("first : -----: ", c)
	dfs(res, c, 0)
	RES := []string{}
	fmt.Printf("res: %#v\n", res)
	for key := range res {
		RES = append(RES, key)
	}
	return RES
}

func dfs(res map[string]bool, c []byte, i int) {
	if i == len(c)-1 {
		fmt.Println("----final: ", i, res)
		res[string(c)] = true
		return
	} else {
		for j := i; j < len(c); j++ {
			c[i], c[j] = c[j], c[i]
			fmt.Println("----------: ", i, j, string(c))
			dfs(res, c, i+1)
			c[i], c[j] = c[j], c[i]
			fmt.Println("---111-------: ", i, j, string(c))
		}
	}
}
