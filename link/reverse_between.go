package link

import "goalgorithms/common"

// 反转链表II(leetcode-92)
// 给定单链表的头指针head和两个整数left和right，其中left<=right 。请你反转从位置left到位置right的链表节点，返回反转后的链表 。
// 示例1：输入：head=[1,2,3,4,5], left=2, right=4, 输出：[1,4,3,2,5]
// 示例2：输入：head=[5], left=1, right=1, 输出：[5]

// reverseBetween1 穿针引线(截取left-right节点并反转，再串联)
func reverseBetween1(head *common.ListNode, left int, right int) *common.ListNode {
	var (
		// 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
		dummy = &common.ListNode{Val: -1, Next: head}
		pre   = dummy
	)

	//1->2->3->4->5->null
	// 1. 找取left节点
	for count := 1; count < left && pre != nil; count++ {
		pre = pre.Next
	}

	// 2. 找出right位置节点
	var rightNode = pre.Next // rightNode记录right位置节点
	for count := 1; count < right-left+1; count++ {
		rightNode = rightNode.Next
	}

	// 3.切断出一个子链表（截取链表）
	rightRemain := rightNode.Next // 右侧剩余节点
	rightNode.Next = nil          // 截取right位置后面的节点
	leftNode := pre.Next          // 获取取left位置节点
	pre.Next = nil                // 截断left-right节点

	// 3. 反转链表(left-right节点)
	var (
		temp *common.ListNode
		cur  = leftNode
	)
	for cur != nil {
		next := cur.Next
		cur.Next = temp
		temp = cur
		cur = next
	}

	// 4. 拼接剩余链表
	leftNode.Next = rightRemain
	pre.Next = rightNode

	return dummy.Next
}

// reverseBetween2 一次遍历，将头节点插到前面(不考虑left及right超限问题)
func reverseBetween2(head *common.ListNode, left int, right int) *common.ListNode {
	var (
		dummyHead = &common.ListNode{Val: -1, Next: head} // 新的头节点
		pre       = dummyHead                             // left的前一节点
	)

	// 1. 找到pre节点
	for count := 1; count < left && pre != nil; count++ {
		pre = pre.Next
	}
	// cur指向left位置节点:3
	cur := pre.Next

	//1->2->3->4->5->null
	for count := 1; count < right-left+1; count++ {
		// cur: 3
		next := cur.Next     // 保存next节点：4
		temp := pre.Next     // 保存pre的下一节点：3
		cur.Next = next.Next // 将当前节点指向下一节点的下一节点:3->5
		next.Next = temp     // 将next节点插到头部:4->3
		pre.Next = next      // 将pre的下一节点保存成next节点：2->4
	}

	return dummyHead.Next
}
