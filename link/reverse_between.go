package link

import "goalgorithms/common"

//反转链表II(leetcode-92)
//给定单链表的头指针head和两个整数left和right，其中left<=right。请你反转从位置left到位置right的链表节点，返回反转后的链表。
//示例1：输入：head=[1,2,3,4,5],left=2,right=4,输出：[1,4,3,2,5]
//示例2：输入：head=[5],left=1,right=1,输出：[5]

func reverseBetween1(head *common.ListNode, left, right int) *common.ListNode {
	// 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
	dummyNode := &common.ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	// 第 1 步：从虚拟头节点走 left - 1 步，来到 left 节点的前一个节点
	// 建议写在 for 循环里，语义清晰
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 第 2 步：从 pre 再走 right - left + 1 步，来到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 第 3 步：切断出一个子链表（截取链表）
	leftNode := pre.Next
	curr := rightNode.Next

	// 注意：切断链接
	pre.Next = nil
	rightNode.Next = nil

	// reverseBetween1 穿针引线(截取left-right节点并反转，再串联)
	var reverseLinkedList = func(head *common.ListNode) {
		var pre *common.ListNode
		cur := head
		for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
	}

	// 第 4 步：同第 206 题，反转链表的子区间
	reverseLinkedList(leftNode)

	// 第 5 步：接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
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

// 一次遍历(考虑left>len(head),left>right,righ>len(head)的情况)
func reverseBetween3(head *common.ListNode, left int, right int) *common.ListNode {
	var dummyPre, partHead *common.ListNode
	for cur, start := head, 1; cur != nil && start <= right; cur, start = cur.Next, start+1 {
		if start == left-1 {
			dummyPre = cur
		} else if start == left {
			var (
				newCur = cur
				pre    *common.ListNode
			)
			// 1,2,3,4,5
			for ; newCur != nil && start <= right; start++ {
				//next := newCur.Next
				//newCur.Next = pre
				//pre = newCur
				//newCur = next
				newCur.Next, pre, newCur = pre, newCur, newCur.Next
			}
			partHead = pre
			cur.Next = newCur
			break
		}
	}
	if dummyPre == nil && partHead == nil {
		return head
	} else if dummyPre == nil {
		return partHead
	} else if partHead != nil {
		dummyPre.Next = partHead
	}
	return head
}
