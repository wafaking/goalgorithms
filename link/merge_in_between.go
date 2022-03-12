package link

//给你两个链list1和list2，它们包含的元素分别为n个和m个。
//请你将list1中下标从a到b的全部节点都删除，并将list2接在被删除节点的位置。
//下图中蓝色边和节点展示了操作后的结果：
//请你返回结果链表的头指针。

//示例 1：
//输入：list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
//输出：[0,1,2,1000000,1000001,1000002,5]
//解释：我们删除 list1 中下标为 3 和 4 的两个节点，并将 list2 接在该位置。上图中蓝色的边和节点为答案链表。

//示例 2：
//输入：list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
//输出：[0,1,1000000,1000001,1000002,1000003,1000004,6]
//解释：上图中蓝色的边和节点为答案链表。

// 双指针法，记录pre, next位置，将list2.Next=next, pre.Next=list2
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	//1. 特殊案例
	// 1.1 list1为空链表时，直接返回list2s
	if list1 == nil {
		return list2
	}

	// 1.2 当输入的起始、截止位置大小相反时，交换位置
	if a > b {
		a, b = b, a
	}

	// 1.3 当a,b输入的值小于0的情况下的判断
	if a <= 0 && b < 0 { // 都小于0，返回list1
		return list1
	} else if a < 0 { // a小于0时，让a=0
		a = 0
	}

	// 2. 双指针法，开始截取
	var (
		count   int                               // 记录第二个指针移动的位置
		newHead = &ListNode{Val: -1, Next: list1} // 虚拟头节点，防止从0开始截取
		pre     = newHead                         // 双指针的前一个指针
		cur     = newHead.Next                    // 双指针的第二个指针（从list1第一个元素开始）
	)
	for cur != nil {
		if count == a-1 { // 记录截取位置的前一个节点
			pre = cur
		}
		if count == b { // 要截取的截止位置
			newList2 := addTailList(list2, cur.Next) // 将当前节点之后节点添加到list2的尾部
			pre.Next = newList2                      // 将前面的链表与后面的链表重新串起来
			break
		}
		cur = cur.Next
		count++
	}

	// list1的长度小于b
	if count < b {
		pre.Next = list2
	}
	return newHead.Next
}

// 将list链表添加到head链表的尾部
func addTailList(head, list *ListNode) *ListNode {
	if head == nil {
		return list
	}
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = list
	return head
}
