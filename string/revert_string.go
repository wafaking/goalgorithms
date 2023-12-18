package str

//反转字符串(leetcode-344)
//编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组s的形式给出。
//不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用O(1)的额外空间解决这一问题。
//示例1：输入：s=["h","e","l","l","o"]输出：["o","l","l","e","h"]
//示例2：输入：s=["H","a","n","n","a","h"]输出：["h","a","n","n","a","H"]

func reverseString1(s []byte) {
	var start, end = 0, len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func reverseString2(s []byte) {
	for i, n := 0, len(s)-1; i <= n/2; i++ {
		s[i], s[n-i] = s[n-i], s[i]
	}
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
