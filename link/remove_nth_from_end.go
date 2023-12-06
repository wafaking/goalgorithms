package link

import "goalgorithms/common"

//删除链表的倒数第N个结点(leetcode-19)
//给你一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。
//示例1：输入：head=[1,2,3,4,5],n=2输出：[1,2,3,5]
//示例2：输入：head=[1],n=1输出：[]
//示例3：输入：head=[1,2],n=1输出：[1]

// 双指针
func removeNthFromEnd1(head *common.ListNode, n int) *common.ListNode {
	var (
		dummy      = &common.ListNode{Next: head}
		slow, fast = dummy, dummy
		i          int
	)
	for fast.Next != nil {
		fast = fast.Next
		if i < n {
			i++
			continue
		}
		slow = slow.Next
	}
	if i == n { // 防止n大于head的长度
		slow.Next = slow.Next.Next
	}
	return dummy.Next
}

// 双指针(同1)
func removeNthFromEnd2(head *common.ListNode, n int) *common.ListNode {
	var (
		dummy      = &common.ListNode{Next: head}
		slow, fast = dummy, dummy
		i          = 0
	)
	// 如果长度小于n，则fast为nil
	for ; i < n && fast.Next != nil; i++ {
		fast = fast.Next
	}
	for ; fast != nil && fast.Next != nil; fast = fast.Next {
		slow = slow.Next
	}
	if i == n {
		slow.Next = slow.Next.Next
	}
	return dummy.Next
}

// 获取长度
func removeNthFromEnd3(head *common.ListNode, n int) *common.ListNode {
	var (
		dummy  = &common.ListNode{Next: head}
		length = 0
	)
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	if length < n {
		return head
	}
	for i, cur := 0, dummy; ; i++ {
		if i == length-n {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return dummy.Next
}

// 使用栈
func removeNthFromEnd4(head *common.ListNode, n int) *common.ListNode {
	var (
		stack = make([]*common.ListNode, 0)
		dummy = &common.ListNode{Next: head}
	)
	for cur := dummy; cur != nil; cur = cur.Next {
		stack = append(stack, cur)
	}
	//0, 1
	for i := 0; len(stack) > 0; i++ {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if i == n {
			cur.Next = cur.Next.Next
			break
		}
	}

	return dummy.Next
}
