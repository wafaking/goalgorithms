package str

//反转字符串中的元音字母(leetcode-345)
//给你一个字符串s，仅反转字符串中的所有元音字母，并返回结果字符串。
//元音字母包括'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。
//示例1：输入：s="hello"输出："holle"
//示例2：输入：s="leetcode"输出："leotcede"
//提示：
//	1<=s.length<=3*105
//	s由可打印的ASCII字符组成

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
