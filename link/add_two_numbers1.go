package link

// 两数相加(leetcode-2)
// 给定两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//示例 1： 输入：l1 = [2,4,3], l2 = [5,6,4], 输出：[7,0,8]，即342 + 465 = 807.
//示例 2： 输入：l1 = [0], l2 = [0],输出：[0]
//示例 3： 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9],输出：[8,9,9,9,0,0,0,1]

func addTwoNumbers11(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head               = &ListNode{Val: -1}
		cur1, cur2, curNew = l1, l2, head
		sign               int
	)
	for cur1 != nil || cur2 != nil || sign != 0 {
		var num1, num2 int
		if cur1 != nil {
			num1 = cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			num2 = cur2.Val
			cur2 = cur2.Next
		}
		temp := num1 + num2 + sign
		if temp > 9 {
			sign = 1
			temp %= 10
		} else {
			sign = 0
		}

		curNew.Next = &ListNode{Val: temp}
		curNew = curNew.Next
	}
	return head.Next
}
