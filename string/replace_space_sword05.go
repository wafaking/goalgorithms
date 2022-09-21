package str

import (
	"strings"
	"unicode"
)

//替换空格(sword-05)
//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//示例1：输入：s = "We are happy.", 输出："We%20are%20happy."

func replaceSpace(s string) string {
	var count int
	for _, v := range s {
		if v == ' ' {
			count++
		}
	}
	if count == 0 {
		return s
	}

	var str = strings.Builder{}
	// 预分配内存, 如果是字符串长度可变，在原s基础上修改s，则可以倒序遍历s，减少移动
	str.Grow(len(s) + 2*count)
	for _, v := range s {
		if unicode.IsSpace(v) {
			str.WriteString("%20")
		} else {
			str.WriteRune(v)
		}
	}

	return str.String()
}
