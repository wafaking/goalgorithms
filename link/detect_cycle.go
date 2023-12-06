package link

import "goalgorithms/common"

//环形链表II(leetcode-142)
//给定一个链表的头节点head，返回链表开始入环的第一个节点。如果链表无环，则返回null。
//如果链表中有某个节点，可以通过连续跟踪next指针再次到达，则链表中存在环。为了表示给定链表中的环，
//评测系统内部使用整数pos来表示链表尾连接到链表中的位置（索引从0开始）。如果pos是-1，则在该链表中没有环。注意：pos不作为参数进行传递，仅仅是为了标识链表的实际情况。
//不允许修改链表。
//示例1：输入：head=[3,2,0,-4],pos=1输出：返回索引为1的链表节点
//解释：链表中有一个环，其尾部连接到第二个节点。
//示例2：输入：head=[1,2],pos=0输出：返回索引为0的链表节点
//解释：链表中有一个环，其尾部连接到第一个节点。
//示例3：输入：head=[1],pos=-1输出：返回null
//解释：链表中没有环。

// 哈希表
func detectCycle1(head *common.ListNode) *common.ListNode {
	var m = make(map[*common.ListNode]struct{})
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := m[cur]; ok {
			return cur
		} else {
			m[cur] = struct{}{}
		}
	}
	return nil
}

// 快慢指针
func detectCycle2(head *common.ListNode) *common.ListNode {
	var meet *common.ListNode
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// 找到快慢指针相遇的节点
			meet = fast
			break
		}
	}
	// 无环形
	if meet == nil {
		return nil
	}

	// 慢指针从头开始，与相遇节点每次都走一步，两者相遇的地方即为环形节点入口
	for slow := head; ; slow, meet = slow.Next, meet.Next {
		if slow == meet {
			return slow
		}
	}
}

//快慢指针
func detectCycle3(head *common.ListNode) *common.ListNode {
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // 快慢指针相遇
			for slow = head; ; slow, fast = slow.Next, fast.Next {
				if slow == fast {
					return slow
				}
			}
		}
	}
	return nil
}
