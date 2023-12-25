package queue_stack

//有效的括号(leetcode-20)
//给定一个只包括'('，')'，'{'，'}'，'['，']'的字符串s，判断字符串是否有效。
//有效字符串需满足：
//	左括号必须用相同类型的右括号闭合。
//	左括号必须以正确的顺序闭合。
//	每个右括号都有一个对应的相同类型的左括号。
//	s仅由括号'()[]{}'组成
//示例1：输入：s="()"输出：true
//示例2：输入：s="()[]{}"输出：true
//示例3：输入：s="(]"输出：false

// 使用栈
func isValid1(s string) bool {
	var stack = make([]int32, 0)
	for _, v := range s {
		switch v {
		case '(':
			stack = append(stack, v)
		case '{':
			stack = append(stack, v)
		case '[':
			stack = append(stack, v)
		case ')':
			if len(stack) < 1 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[0 : len(stack)-1]

		case '}':
			if len(stack) < 1 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[0 : len(stack)-1]
		case ']':
			if len(stack) < 1 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 使用栈
func isValid2(s string) bool {
	var n = len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack = make([]byte, 0)
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
