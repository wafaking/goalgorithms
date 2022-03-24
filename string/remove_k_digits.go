package str

import (
	"fmt"
	"strings"
)

//移掉K位数字,使数值最小(402)
//给定以字符串表示的非负整数num和整数k，移除这个数中的k位数字，使得剩下的数字最小。
//示例 1 ： //输入：num = "1432219", k = 3, 输出："1219"
//示例 2 ： //输入：num = "10200", k = 1, 输出："200"
//示例 3 ： //输入：num = "10", k = 2, 输出："0"

// 法1：使用双端队列(注：依次与后面的元素相比，将已比较过的数字放入队列)
func removeKdigits1(num string, k int) string {
	length := len(num)
	if k <= 0 {
		return num
	}
	if length <= k {
		return "0"
	}
	// 注: 上面前两步校验已确定length>1

	//fmt.Printf("num: %s, k: %d\n", num, k)
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

// 借助栈实现
func removeKdigits2(num string, k int) string {
	if k == 0 {
		return num
	}
	if k >= len(num) {
		return "0"
	}

	// 1. 判断num是否有非法字符
	for i := 0; i < len(num); i++ {
		var val = num[i] - '0'
		if val > 9 || val < 0 {
			return ""
		}
	}

	var stack = make([]byte, 0, len(num))
	for i, _ := range num {
		// pop 100001 3
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}

	// 再从末尾截取一遍，防止出现: num=11111 k=2 最后未删除
	stack = stack[:len(stack)-k]
	res := strings.TrimLeft(string(stack), "0")
	if res == "" {
		res = "0"
	}
	return res
}

// 维护递增栈， 当出现 i < i - 1 位置的时候，就进行删除操作
// 要说贪心的话， 就是优先从高位进行删除，因为高位的数字决定整个数字的大小
// 高位数字越小，整个数字就越小
func removeKdigits3(num string, k int) string {
	stack := make([]byte, 0)
	var i int
	// 0. 首先：遍历到最后位置，或者k为0，都不再进行循环
	for i = 0; i < len(num) && k > 0; i++ {
		// 1. 空栈 或者递增 直接插入元素
		// ！ 这里要注意一个问题，前导0的问题，不能将0往空栈中插入
		if (len(stack) == 0 || num[i-1] < num[i]) &&
			num[i] != '0' {
			stack = append(stack, num[i])
			continue
		}
		// 2. 不满足递增，删除元素
		for len(stack) > 0 && stack[len(stack)-1] > num[i] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		// 3. 当前元素入栈, 如果栈为空，当前元素为 0（形成前导0了），不能入栈,否则都入栈
		if len(stack) == 0 && num[i] == '0' {
			continue
		}
		stack = append(stack, num[i])
	}
	fmt.Println(string(stack))
	// 4. 最后要判断，如果整个数字都是递增的， k 还没为 0, 那就接着删除（整个数字都到栈中了，直接从栈顶删除）
	// 针对 10 k = 2 的情况，需要全删, 栈为空，再进行 stack[:len(stack)-1] 下标越界错误，需要判断 len(stack) > 0
	fmt.Println()
	for k > 0 && len(stack) > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	// 5. 拼接最后的结果
	// 针对 "10001" k = 1 的情况， 剩余字符串可能含有前导 0 了，也要进行判断
	for i < len(num) && num[i] == '0' &&
		len(stack) == 0 {
		i++
	}
	rst := string(stack) + num[i:]
	if rst == "" {
		return "0"
	}
	return rst
}
