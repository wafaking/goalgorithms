package link

import "goalgorithms/common"

// 环形链表(leetcode-141)
// 给你一个链表的头节点head，判断链表中是否有环。
// 如果链表中有某个节点，可以通过连续跟踪next指针再次到达，则链表中存在环。为了表示给定链表中的环，
// 评测系统内部使用整数pos来表示链表尾连接到链表中的位置（索引从0开始）。注意：pos不作为参数进行传递。仅仅是为了标识链表的实际情况。
// 如果链表中存在环，则返回true。否则，返回false。
// 示例1：输入：head=[3,2,0,-4],pos=1输出：true
// 解释：链表中有一个环，其尾部连接到第二个节点。
// 示例2：输入：head=[1,2],pos=0输出：true
// 解释：链表中有一个环，其尾部连接到第一个节点。
// 示例3：输入：head=[1],pos=-1输出：false
// 解释：链表中没有环。

// 哈希表
func hasCycle1(head *common.ListNode) bool {
	var m = make(map[*common.ListNode]struct{})
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := m[cur]; !ok {
			m[cur] = struct{}{}
		} else {
			return true
		}
	}
	return false
}

// 快慢指针
func hasCycle2(head *common.ListNode) bool {
	// 1, 2, 3, 4, 5, 6, 1
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
