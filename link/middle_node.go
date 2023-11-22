package link

import "goalgorithms/common"

//链表的中间结点(leetcode-876)
//给定一个头结点为head的非空单链表，返回链表的中间结点。
//如果有两个中间结点，则返回第二个中间结点。
//示例1： 输入：[1,2,3,4,5],输出：此列表中的结点3(序列化形式：[3,4,5]),返回的结点值为3。 (测评系统对该结点序列化表述是 [3,4,5])。
//示例2： 输入：[1,2,3,4,5,6], 输出：此列表中的结点4(序列化形式：[4,5,6])

// middleNode11 计数器，奇数走两步，偶数走一步
func middleNode1(head *common.ListNode) *common.ListNode {
	slow := head
	fast := head
	i := 1
	for fast != nil {
		if i&1 == 0 { // 判断奇偶
			slow = slow.Next
		}
		fast = fast.Next
		i++
	}
	return slow
}

func middleNode2(head *common.ListNode) *common.ListNode {
	var slow, fast = head, head
	// 可以添加虚拟头部元素
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 每次移动一个
		fast = fast.Next.Next // 每次移动两个
	}
	return slow
}
