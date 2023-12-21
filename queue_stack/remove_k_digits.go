package queue_stack

import (
	"strings"
	"unicode"
)

//移掉K位数字,使数值最小(leetcode-402)
//给定以字符串表示的非负整数num和整数k，移除这个数中的k位数字，使得剩下的数字最小。
//示例1：//输入：num="1432219",k=3,输出："1219"
//示例2：//输入：num="10200",k=1,输出："200"
//示例3：//输入：num="10",k=2,输出："0"

// 法1：单调栈+贪心
func removeKdigits1(num string, k int) string {
	if k <= 0 {
		return num
	}
	if len(num) <= k {
		return "0"
	}

	var isNumAllDigit = func(num string) bool {
		for k := range num {
			if !unicode.IsDigit(rune(num[k])) {
				return false
			}
		}
		return true
	}

	// 判断是否有不合法字符
	if !isNumAllDigit(num) {
		return "0"
	}

	var sli []byte // 栈
	for i := 0; i < len(num); i++ {
		// 先添加进栈中
		sli = append(sli, num[i])

		// 第一个数直接添加
		if i == 0 {
			continue
		}

		for { // 回顾单调栈
			var i = len(sli) - 1
			if (i > 0) && (sli[i] < sli[i-1]) && k > 0 {
				// 将栈内i<i-1位置的数据都弹出
				sli = append(sli[:i-1], sli[i:]...)
				k--
				continue
			}
			// 所有i-1<i都成立，单调栈回顾完成
			break
		}
	}

	// 去掉所有前缀0，注：使用trimleft能删除所有前缀0，用trimPrefix只能删除一个0
	sli = []byte(strings.TrimLeft(string(sli), "0"))
	if len(sli) > k {
		return string(sli[:len(sli)-k])
	}
	return "0"
}

// 法1：单调栈+贪心
func removeKdigits2(num string, k int) string {
	if k == 0 {
		return num
	}
	if k >= len(num) {
		return "0"
	}

	// 判断num是否有非法字符
	for i := 0; i < len(num); i++ {
		var val = num[i] - '0'
		if val > 9 || val < 0 {
			return ""
		}
	}

	var stack = make([]int32, 0, len(num))
	for _, v := range num {
		// 移除栈中比v大的元素
		for len(stack) > 0 && k > 0 && (v < stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
			k--
		}

		// 将v添加到栈中
		if v != '0' || len(stack) > 0 {
			stack = append(stack, v)
		}
	}

	// 特殊情况
	if len(stack) == 0 || k >= len(stack) {
		return "0"
	}
	// 排队特殊情况后，只剩下len(stack)>k的情况，此时stack中不可能有0,因此直接返回
	return string(stack[:len(stack)-k])
}
