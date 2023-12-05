package link

import "goalgorithms/common"

//移除链表元素(leetcode-203)
//给你一个链表的头节点{head{和一个整数{val{，请你删除链表中所有满足{Node.val{=={val{的节点，并返回{新的头节点{。
//示例1：输入：head=[1,2,6,3,4,5,6],val=6输出：[1,2,3,4,5]
//示例2：输入：head=[],val=1输出：[]
//示例3：输入：head=[7,7,7,7],val=7输出：[]

// 迭代
func removeElements1(head *common.ListNode, val int) *common.ListNode {
	var (
		dummy = &common.ListNode{Next: head}
		cur   = dummy
	)
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}

// TODO 递归
func removeElements2(head *common.ListNode, val int) *common.ListNode {
	return nil
}
