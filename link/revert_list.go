package link

import "goalgorithms/common"

// 反转链表(leet-206/sword-24)
// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
// 示例: 输入: 1->2->3->4->5->NULL, 输出: 5->4->3->2->1->NULL

// 使用栈
func reverseList1(head *common.ListNode) *common.ListNode {
	var (
		stack = make([]int, 0)
		dummy = &common.ListNode{}
	)
	for cur := head; cur != nil; cur = cur.Next {
		stack = append(stack, cur.Val)
	}
	for cur := dummy; len(stack) > 0; cur = cur.Next {
		cur.Next = &common.ListNode{Val: stack[len(stack)-1]}
		stack = stack[:len(stack)-1]
	}
	return dummy.Next
}

// 循环
func reverseList2(head *common.ListNode) *common.ListNode {
	var pre *common.ListNode
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 循环(省去中间变量)
func reverseList3(head *common.ListNode) *common.ListNode {
	var pre *common.ListNode
	for cur := head; cur != nil; {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}

//1,2,3,4
// 递归
func reverseList4(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 假设链表为：N1->N2->...->N(k-1)->N(k)->N(k+1)->N(m)
	// 如果N(k+1)之前的节点已经翻转，则N(k+1)的下一个节点即为N(k),即head.Next.Next = head
	newHead := reverseList4(head.Next)
	head.Next.Next = head
	//newHead.Next = head
	// 断开，防止循环
	head.Next = nil
	return newHead
}
