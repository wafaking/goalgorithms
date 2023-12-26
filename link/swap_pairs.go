package link

import "goalgorithms/common"

//两两交换链表中的节点(leetcode-24)
//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//示例1：输入：head=[1,2,3,4]输出：[2,1,4,3]
//示例2：输入：head=[]输出：[]
//示例3：输入：head=[1]输出：[1]
//提示：
//	链表中节点的数目在范围[0,100]内
//	0<=Node.val<=100

func swapPairs1(head *common.ListNode) *common.ListNode {
	var (
		dummy = &common.ListNode{Next: head}
		pre   = dummy
		cur   = head
	)

	// 1,2,3,4
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur
		pre.Next = next

		pre = cur
		cur = cur.Next
	}
	return dummy.Next
}
