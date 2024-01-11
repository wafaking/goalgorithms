package link

import "goalgorithms/common"

//排序链表(leetcode-148)
//给你链表的头结点head，请将其按升序排列并返回排序后的链表。
//示例1：输入：head=[4,2,1,3]输出：[1,2,3,4]
//示例2：输入：head=[-1,5,3,4,0]输出：[-1,0,3,4,5]
//示例3：输入：head=[]输出：[]
//提示：
//	链表中节点的数目在范围[0,5*10^4]内
//	-10^5<=Node.val<=10^5
//进阶：你可以在O(nlogn)时间复杂度和常数级空间复杂度下，对链表进行排序吗？

// 遍历+重新生成链表
func sortList1(head *common.ListNode) *common.ListNode {
	var (
		list      = make([]int, 0)
		quickSort func(list []int)
	)
	for cur := head; cur != nil; cur = cur.Next {
		list = append(list, cur.Val)
	}
	quickSort = func(list []int) {

	}
	quickSort(list)

	var (
		dummy = new(common.ListNode)
		cur   = dummy
	)
	for _, v := range list {
		node := &common.ListNode{Val: v}
		cur.Next = node
		cur = cur.Next
	}

	return dummy.Next
}
