package str

import (
	"sort"
)

// 字母异位词分组(leetcode-49)
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
// 示例 1:
//     输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//     输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
// 示例 2:
//     输入: strs = [""]
//     输出: [[""]]
// 示例 3:
// 输入: strs = ["a"]
// 输出: [["a"]]
// 提示：
//     1 <= strs.length <= 10^4
//     0 <= strs[i].length <= 100
//     strs[i] 仅包含小写字母

func groupAnagrams1(strs []string) [][]string {
	var (
		m   = make(map[string][]string, 0)
		res = make([][]string, 0)
	)

	for _, str := range strs {
		// 1.[]string+sort.Strings()
		// var arr = make([]string, 0, len(str))
		// for _, c := range str {
		// 	arr = append(arr, string(c))
		// }
		// var arr = make([]byte, 0, len(str))
		// sort.Strings(arr)
		// newStr := strings.Join(arr, "")

		// 2.使用[]byte+sort.Slice()
		arr := []byte(str)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		newStr := string(arr)

		if list, ok := m[newStr]; ok {
			m[newStr] = append(list, str)
		} else {
			m[newStr] = []string{str}
		}
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func groupAnagrams2(strs []string) [][]string {
	var (
		m   = make(map[[26]int][]string, 0)
		res = make([][]string, 0)
	)
	for _, str := range strs {
		var bt [26]int
		for _, v := range str {
			bt[v-'a']++
		}
		m[bt] = append(m[bt], str)
	}

	for _, v := range m {
		res = append(res, v)
	}
	return res
}
