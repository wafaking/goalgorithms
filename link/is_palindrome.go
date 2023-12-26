package link

import "goalgorithms/common"

//回文链表(leetcode-234)
//给你一个单链表的头节点head，请你判断该链表是否为回文链表。如果是，返回true；否则，返回false。
//示例1：输入：head=[1,2,2,1]输出：true
//示例2：输入：head=[1,2]输出：false
//注：0<=Node.val<=9
//进阶：你能否用O(n)时间复杂度和O(1)空间复杂度解决此题？

// 使用数组
func isPalindrome1(head *common.ListNode) bool {
	var list = make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		list = append(list, cur.Val)
	}
	for i := 0; i <= (len(list)-1)>>1; i++ {
		if list[i] != list[len(list)-1-i] {
			return false
		}
	}
	return true
}

// 递归
func isPalindrome2(head *common.ListNode) bool {
	var (
		cur              = head
		recursivelyCheck func(head *common.ListNode) bool
	)
	// 1,2,3,4,5
	recursivelyCheck = func(head *common.ListNode) bool {
		if head == nil {
			return true
		}

		if recursivelyCheck(head.Next) {
			if head.Val == cur.Val {
				cur = cur.Next
				return true
			}
		}
		return false
	}
	return recursivelyCheck(head)
}

// 递归
func isPalindrome3(head *common.ListNode) bool {
	cur := head
	var recursivelyCheck func(*common.ListNode) bool
	recursivelyCheck = func(curNode *common.ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != cur.Val {
				return false
			}
			cur = cur.Next
		}
		return true
	}
	return recursivelyCheck(head)
}

// 快慢指针(使用O(n)时间复杂度和O(1)空间复杂度)
func isPalindrome4(head *common.ListNode) bool {
	var slow, fast = head, head
	// 找到中间节点slow,slow为后半部分开头
	// 1,2,3,4,5,6
	// 1,1-->2,3-->3,5-->4,nil
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}

	// 1->2->3->4->5->6
	// 返转后半部分链表
	var pre *common.ListNode
	for cur := slow; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	// 与前半部分链表对比(此处借用fast变量指向前半部分开头)
	for fast = head; pre != nil; pre, fast = pre.Next, fast.Next {
		if pre.Val != fast.Val {
			return false
		}
	}
	return true
}
