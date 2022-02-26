package str

import "fmt"

func reverseString(s []byte) {
	var start, end = 0, len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	fmt.Println(s)
}

// 将字符串中的元音字母翻转
// 	Example 1:
// 	Input: "hello"
// Output: "holle"
// 	Example 2:
// 	Input: "leetcode"
// Output: "leotcede"
func reverseVowels(s string) string {
	var vowelMap = map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	var sb = []byte(s)
	var start, end = 0, len(sb) - 1

	for start < end {
		if _, ok := vowelMap[sb[start]]; !ok {
			start++
			continue
		}

		if _, ok := vowelMap[sb[end]]; !ok {
			end--
			continue
		}

		sb[start], sb[end] = sb[end], sb[start]
		start++
		end--
	}

	return string(sb)
}
