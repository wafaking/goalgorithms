package link

import "goalgorithms/common"

//删除排序链表中的重复元素II(leetcode-82)
//给定一个已排序的链表的头head，删除原始链表中所有重复数字的节点，只留下不同的数字。返回已排序的链表。
//示例1：输入：head=[1,2,3,3,4,4,5]输出：[1,2,5]
//示例2：输入：head=[1,1,1,2,3]输出：[2,3]

// 一次遍历
func deleteDuplicates21(head *common.ListNode) *common.ListNode {
	var (
		newHead = &common.ListNode{Val: -1, Next: head}
		pre     = newHead
		cur     = head
	)

	for cur != nil && cur.Next != nil {
		if cur.Val != cur.Next.Val {
			cur = cur.Next
			pre = pre.Next
			continue
		}
		for {
			if cur.Next != nil && cur.Next.Val == cur.Val {
				cur = cur.Next
			} else {
				cur = cur.Next
				pre.Next = cur
				break
			}
		}
	}
	return newHead.Next
}

// 一次遍历
func deleteDuplicates22(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}

	var (
		dummy = &common.ListNode{Val: 0, Next: head}
		cur   = dummy
	)

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
