package link

import (
	"goalgorithms/common"
	"math/rand"
	"time"
)

//排序链表(leetcode-148)
//给你链表的头结点head，请将其按升序排列并返回排序后的链表。
//示例1：输入：head=[4,2,1,3]输出：[1,2,3,4]
//示例2：输入：head=[-1,5,3,4,0]输出：[-1,0,3,4,5]
//示例3：输入：head=[]输出：[]
//提示：
//	链表中节点的数目在范围[0,5*10^4]内
//	-10^5<=Node.val<=10^5
//进阶：你可以在O(nlogn)时间复杂度和常数级空间复杂度下，对链表进行排序吗？

// 遍历+快排+重新生成链表(易超时)
func sortList1(head *common.ListNode) *common.ListNode {
	var (
		list      = make([]int, 0)
		quickSort func(list []int) []int
	)
	for cur := head; cur != nil; cur = cur.Next {
		list = append(list, cur.Val)
	}
	rand.Seed(time.Now().UnixNano())
	quickSort = func(list []int) []int {
		var n = len(list)
		if n <= 1 {
			return list
		}

		pivot := rand.Intn(n)
		list[0], list[pivot] = list[pivot], list[0]

		var left, right = make([]int, 0), make([]int, 0)
		for i := 1; i < len(list); i++ {
			if list[0] <= list[i] {
				right = append(right, list[i])
			} else {
				left = append(left, list[i])
			}
		}
		return append(append(quickSort(left), list[0]), quickSort(right)...)
	}
	list = quickSort(list)

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

// 快慢指针+归并排序
func sortList2(head *common.ListNode) *common.ListNode {
	var (
		// 拆分
		mergeSortCore func(head, tail *common.ListNode) *common.ListNode
		// 合并
		merge func(head1, head2 *common.ListNode) *common.ListNode
	)
	// head,tail为左闭右开区间
	mergeSortCore = func(head, tail *common.ListNode) *common.ListNode {
		if head == nil {
			return nil
		}
		// 左闭右开不能超出右边界
		if head.Next == tail {
			head.Next = nil
			return head
		}

		// 快慢指针，找出中间节点slow
		var slow, fast = head, head
		for fast != tail {
			slow = slow.Next
			fast = fast.Next
			if fast != tail {
				fast = fast.Next
			}
		}
		return merge(mergeSortCore(head, slow), mergeSortCore(slow, tail))
	}

	merge = func(head1, head2 *common.ListNode) *common.ListNode {
		var (
			dummy      = new(common.ListNode)
			cur        = dummy
			cur1, cur2 = head1, head2
		)
		for cur1 != nil || cur2 != nil {
			if cur1 == nil {
				cur.Next = cur2
				break
			} else if cur2 == nil {
				cur.Next = cur1
				break
			}

			if cur1.Val <= cur2.Val {
				temp := cur1.Next
				cur1.Next = nil
				cur.Next = cur1
				cur1 = temp
			} else {
				temp := cur2.Next
				cur2.Next = nil
				cur.Next = cur2
				cur2 = temp
			}
			cur = cur.Next
		}
		return dummy.Next
	}
	return mergeSortCore(head, nil)
}

// 迭代(TODO)
func sortList3(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}

	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	var merge = func(head1, head2 *common.ListNode) *common.ListNode {
		dummyHead := &common.ListNode{}
		temp, temp1, temp2 := dummyHead, head1, head2
		for temp1 != nil && temp2 != nil {
			if temp1.Val <= temp2.Val {
				temp.Next = temp1
				temp1 = temp1.Next
			} else {
				temp.Next = temp2
				temp2 = temp2.Next
			}
			temp = temp.Next
		}
		if temp1 != nil {
			temp.Next = temp1
		} else if temp2 != nil {
			temp.Next = temp2
		}
		return dummyHead.Next
	}

	dummyHead := &common.ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *common.ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}
