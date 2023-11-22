package link

import "goalgorithms/common"

// 反转链表(leet-206/sword-24)
// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
// 示例: 输入: 1->2->3->4->5->NULL, 输出: 5->4->3->2->1->NULL

// reverseList1 法一：循环赋值
func reverseList1(head *common.ListNode) *common.ListNode {
	var (
		cur = head
		pre *common.ListNode
	)
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// reverseList2 法二：递归
func reverseList2(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 假设链表为：N1->N2->...->N(k-1)->N(k)->N(k+1)->N(m)
	// 如果N(k+1)之前的节点已经翻转，则N(k+1)的下一个节点即为N(k),即head.Next.Next = head
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseList3(head *common.ListNode) *common.ListNode {
	var (
		cur = head
		pre *common.ListNode
	)
	//var new *Student //  5,4,3,2,1
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

func reverseList4(head *common.ListNode) *common.ListNode {
	var (
		cur = head
		pre *common.ListNode
	)

	// 1->2->3->4->5->null
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
