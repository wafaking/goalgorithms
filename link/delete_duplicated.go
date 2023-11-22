package link

import "goalgorithms/common"

//删除排序链表中的重复元素(leetcode-83)
//给定一个已排序的链表的头head，删除所有重复的元素，使每个元素只出现一次。返回已排序的链表。
//示例1：输入：head=[1,1,2]输出：[1,2]
//示例2：输入：head=[1,1,2,3,3]输出：[1,2,3]

// 一次遍历
func deleteDuplicates11(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	cur := head // cur指向头结点，而且改变cur指向不会影响到head的指向
	for cur.Next != nil {
		if cur.Val == cur.Next.Val { // 当前节点的值等于下个节点的值
			cur.Next = cur.Next.Next // cur指向下下个节点（cur指针不会移动）
		} else { // 当前节点的值不等于下个节点的值
			cur = cur.Next // 当前指针后移到下一个不同值的节点
		}
	}
	return head
}

// 一次遍历
func deleteDuplicates12(head *common.ListNode) *common.ListNode {
	for cur := head; cur != nil && cur.Next != nil; {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
