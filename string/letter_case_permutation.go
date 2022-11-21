package str

import "unicode"

// todo ---------------
// 字母大小写全排列(leetcode-784)
// 给定字符串s，通过将字符串s中的每个字母转变大小写，可以获得一个新的字符串，返回所有可能得到的字符串集合。
// 示例1:输入：s="a1b2", 输出：["a1b2", "a1B2", "A1b2", "A1B2"],
// 示例2:输入: s="3z4", 输出: ["3z4","3Z4"]

func letterCasePermutation11(s string) []string {
	var (
		ans   = []string{s}
		bytes = []byte(s)
	)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 'a' && bytes[i] <= 'z' {
			bytes[i] = bytes[i] - 32
			ans = append(ans, string(bytes))

		} else if bytes[i] >= 'A' && bytes[i] <= 'Z' {
			bytes[i] = bytes[i] + 32
			ans = append(ans, string(bytes))
		} else {
			continue
		}
		for j := i + 1; j < len(bytes); j++ {
			if bytes[j] >= 'a' && bytes[j] <= 'z' {
				bytes[j] = bytes[j] - 32
				ans = append(ans, string(bytes))
			} else if bytes[j] >= 'A' && bytes[j] <= 'Z' {
				bytes[j] = bytes[j] + 32
				ans = append(ans, string(bytes))
			}
		}

	}
	return ans
}

func letterCasePermutation12(s string) []string {
	var (
		ans       []string
		path      []byte
		bytes     = []byte(s)
		backTrace func(idx int)
	)

	backTrace = func(idx int) {
		if len(path) == len(bytes) {
			ans = append(ans, string(path))
			return
		}
		if unicode.IsLetter(rune(bytes[idx])) {

		} else {
			//path = append()
		}

		for i := idx; i < len(bytes); i++ {
			var temp = bytes[i]

			if temp >= 'a' && temp <= 'z' {
				path = append(path, bytes[i])
				backTrace(i + 1)

				// 恢复
				path = path[:len(path)-1]

				//变形
				//bytes[i] = bytes[i] - 32
				path = append(path, bytes[i]-32)
				backTrace(i + 1)

				path = path[:len(path)-1]
				//bytes[i] = bytes[i] + 32

			} else if temp >= 'A' && temp <= 'Z' {
				path = append(path, bytes[i])
				backTrace(i + 1)

				path = path[:len(path)-1]
				//bytes[i] = bytes[i] + 32
				path = append(path, bytes[i]+3)
				backTrace(i + 1)

				path = path[:len(path)-1]
				//bytes[i] = bytes[i] - 32
			} else {
				path = append(path, bytes[i])
				//backTrace(i + 1)
			}

		}
	}

	backTrace(0)
	return ans
}

// letterCasePermutation13 使用位数组枚举
//func letterCasePermutation13(s string) []string {
//	var (
//		ans   [][]string
//		bytes = []byte(s)
//	)
//	for mask := 0; mask < 1<<len(bytes); mask++ {
//		var path []string
//		for i := 0; i < len(bytes); i++ {
//			if unicode.IsLetter(bytes[i])
//			if (1 << i & mask) > 0 {
//				path = append(path, bytes[i])
//			}
//		}
//		if len(path) == len(bytes) {
//			ans = append(ans, path)
//		}
//	}
//	return ans
//}
