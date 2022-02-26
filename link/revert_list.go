package link

//剑指 Offer 24. 反转链表
//定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//示例:
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL

func reverseList() *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
