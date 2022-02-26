package str

import (
	"fmt"
	"strings"
)

//移掉 K 位数字,使数值最小(402)

//给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。

//示例 1 ：
//输入：num = "1432219", k = 3
//输出："1219"
//解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。

//示例 2 ：
//输入：num = "10200", k = 1
//输出："200"
//解释：移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。

//示例 3 ：
//输入：num = "10", k = 2
//输出："0"
//解释：从原数字移除所有的数字，剩余为空就是 0 。

// 法1：使用双端队列(注：依次与后面的元素相比，将已比较过的数字放入队列)
func removeKdigits(num string, k int) string {
	length := len(num)
	if k <= 0 {
		return num
	}

	if length <= k {
		return "0"
	}
	// 注: 上面前两步校验已确定length>1

	fmt.Printf("num: %s, k: %d\n", num, k)
	// 判断是否有不合法字符
	for _, v := range num {
		val := v - '0'
		if val < 0 || val > 9 {
			return "0"
		}
	}

	// 1、依次与后面元素比较，将比较过的数字放入栈
	// 1 4 5 5 4 7 8 9 1 2  即1与4、4与5、5与4比较
	//cur := 0   // 当前游标
	times := 0 // 移除的次数
	sli := make([]byte, 0)
	leftStr := ""
	for cur := 0; cur < length; cur++ {
		if times == k {
			// 如果删除元素的个数已达到指定数目，将后续元素添加到列表中
			leftStr = num[cur:]
			break
		}

		curNum := num[cur] - '0'
		if len(sli) == 0 {
			sli = append(sli, curNum)
			continue
		}

		if curNum >= sli[len(sli)-1] {
			sli = append(sli, curNum)
			continue
		}

		// 注：等于号要放此处，防止数据不够，如上面数据只删除1位，则必须删除5
		for times < k {
			if len(sli) > 0 && curNum < sli[len(sli)-1] {
				sli = sli[:len(sli)-1]
				times++
				continue
			}
			break
		}
		sli = append(sli, curNum) // 注意需要把当前元素添加进去
	}

	// 再检查一遍，防止删除的元素数<k
	if times < k {
		sli = sli[:len(sli)-(k-times)]
	}

	// 2、将栈内元素依次弹出组成字符串并返回结果
	res := ""
	for _, v := range sli {
		res = fmt.Sprintf("%s%d", res, v)
	}
	res = strings.TrimLeft(fmt.Sprintf("%s%s", res, leftStr), "0")
	if len(res) == 0 {
		res = "0"
	}
	return res
}
