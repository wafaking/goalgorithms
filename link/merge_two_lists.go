package link

// 合并两个有序链表(leetcode-21)
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 示例1：
// 输入：l1 = [1,2,4], l2 = [1,3,4],  输出：[1,1,2,3,4,4]

// 示例 3：
// 输入：l1 = [], l2 = [0], 输出：[0]

// 法一：使用循环
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var (
		newHead = &ListNode{Val: -1}
		cur     = newHead
	)
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = &ListNode{Val: list2.Val}
			cur = cur.Next
			list2 = list2.Next
			continue
		} else if list2 == nil {
			cur.Next = &ListNode{Val: list1.Val}
			cur = cur.Next
			list1 = list1.Next
			continue
		}

		if list1.Val > list2.Val {
			cur.Next = &ListNode{
				Val: list2.Val,
			}
			list2 = list2.Next

		} else if list1.Val < list2.Val {
			cur.Next = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next

		} else {
			cur.Next = &ListNode{
				Val: list1.Val,
			}
			cur = cur.Next

			cur.Next = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
			list2 = list2.Next
		}
		cur = cur.Next
	}
	return newHead.Next
}

// 法一：使用循环(直接使用list1)
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {

	var (
		newHead = &ListNode{Val: -1}
		cur     = newHead
	)
	for list1 != nil || list2 != nil {

		if list1 != nil && list2 != nil {
			if list1.Val > list2.Val {
				cur.Next = list2
				list2 = list2.Next
			} else if list1.Val < list2.Val {
				cur.Next = list1
				list1 = list1.Next
			} else {
				cur.Next = list1
				cur = cur.Next
				list1 = list1.Next
				cur.Next = list2
				list2 = list2.Next
			}
			cur = cur.Next

		} else if list1 == nil {
			cur.Next = list2
			break
		} else if list2 == nil {
			cur.Next = list1
			break
		}

	}

	return newHead.Next
}

// 法三：使用递归
func mergeTwoLists3(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		list2.Next = mergeTwoLists3(list1, list2.Next)
		return list2
	} else {
		list1.Next = mergeTwoLists3(list1.Next, list2)
		return list1
	}
}
