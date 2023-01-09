package link

// 二进制链表转整数(leetcode-1290)
// 给定单链表的头结点head。链表中每个结点的值不是0就是1。已知此链表是一个整数数字的二进制表示形式。
// 请你返回该链表所表示数字的十进制值 。
// 示例1：输入：head=[1,0,1] 输出：5 解释：二进制数 (101) 转化为十进制数 (5)
// 示例2：输入：head=[0,0] 输出：0
// 示例3：输入：head=[1] 输出：1
// 示例4：输入：head=[1,0,0,1,0,0,1,1,1,0,0,0,0,0,0] 输出：18880

func getDecimalValue1(head *ListNode) int {
	var (
		sli []int
		sum int
		mod int
	)

	for i := 0; head != nil; i, head = i+1, head.Next {
		sli = append(sli, head.Val)
	}

	for i := len(sli) - 1; i >= 0; i-- {
		if i == len(sli)-1 {
			mod = 1
		}
		if sli[i] > 0 {
			sum += sli[i] * mod
		}
		mod *= 2
	}

	return sum
}

// getDecimalValue2 使用二进制的方式解决
func getDecimalValue2(head *ListNode) int {
	var res int
	for temp := head; temp != nil; temp = temp.Next {
		res = res<<1 | temp.Val
	}
	return res
}

func getDecimalValue3(head *ListNode) int {
	var res int
	for temp := head; nil != temp; temp = temp.Next {
		res = res*2 + temp.Val
	}
	return res
}
