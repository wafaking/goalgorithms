package link

// 链表中的两数相加(sword2-25/leetcode-445)
// 给定两个非空链表l1和l2来代表两个非负整数。数字最高位位于链表开始位置。
// 它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

//示例1：输入：l1 = [7,2,4,3], l2 = [5,6,4], 输出：[7,8,0,7]
//  7->2->4->3
//	   5->6->4
//示例2：输入：l1 = [2,4,3], l2 = [5,6,4], 输出：[8,0,7]
//示例3：输入：l1 = [0], l2 = [0], 输出：[0]

//进阶：如果输入链表不能修改该如何处理？换句话说，不能对列表中的节点进行翻转。
//注意：本题与主站 445 题相同：https://leetcode-cn.com/problems/add-two-numbers-ii/

// 法一：翻转链表
func addTwoNumbers21(l1 *ListNode, l2 *ListNode) *ListNode {
	ll1 := reverseLink(l1)
	ll2 := reverseLink(l2)

	var (
		sign    int
		newHead = &ListNode{Val: -1}
		cur     = newHead
	)
	for ll1 != nil || ll2 != nil || sign > 0 {
		var x, y int
		if ll1 != nil {
			x = ll1.Val
			ll1 = ll1.Next
		}
		if ll2 != nil {
			y = ll2.Val
			ll2 = ll2.Next
		}
		temp := x + y + sign
		if temp > 9 {
			temp %= 10
			sign = 1
		} else {
			sign = 0
		}
		cur.Next = &ListNode{Val: temp}
		cur = cur.Next
	}
	return reverseLink(newHead.Next)
}

func reverseLink(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var newHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	head = newHead
	return head
}

// 法二：借助列表实现(不改变原链表结果,用列表存放每个相加得出的结果)
// 法二与法三相似，但一个是列表的概念，一个是栈的概念
func addTwoNumbers22(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		sli1, sli2, sli []int
		head1           = l1
		head2           = l2
		sign            int
	)

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	for head1 != nil {
		sli1 = append(sli1, head1.Val)
		head1 = head1.Next
	}
	for head2 != nil {
		sli2 = append(sli2, head2.Val)
		head2 = head2.Next
	}

	for lth1, lth2 := len(sli1)-1, len(sli2)-1; lth1 >= 0 || lth2 >= 0 || sign > 0; lth1, lth2 = lth1-1, lth2-1 {
		var x, y int
		if lth1 >= 0 {
			x = sli1[lth1]
		}
		if lth2 >= 0 {
			y = sli2[lth2]
		}
		temp := x + y + sign
		if temp > 9 {
			sign = 1
			temp %= 10
		} else {
			sign = 0
		}
		sli = append(sli, temp)
	}

	var (
		head = &ListNode{Val: -1}
		cur  = head
	)
	for lth := len(sli) - 1; lth >= 0; lth-- {
		cur.Next = &ListNode{Val: sli[lth]}
		cur = cur.Next
	}

	return head.Next
}

// 法二：借助栈实现
func addTwoNumbers23(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		stackA, stackB, stack []int
		sign                  = 0
		list                  = &ListNode{Val: -1}
		cur                   = list
	)

	stackA = appendListToStack(l1, stackA)
	stackB = appendListToStack(l2, stackB)

	for !isEmpty(stackA) || !isEmpty(stackB) || sign > 0 {
		var x, y, temp int
		if !isEmpty(stackA) {
			x, _ = pop(&stackA)
		}
		if !isEmpty(stackB) {
			y, _ = pop(&stackB)
		}
		temp = x + y + sign
		if temp > 9 {
			sign = 1
			temp %= 10
		} else {
			sign = 0
		}
		stack = append(stack, temp)
	}
	for !isEmpty(stack) {
		val, _ := pop(&stack)
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	return list.Next
}

func appendListToStack(list *ListNode, stack []int) []int {
	for list != nil {
		stack = append(stack, list.Val)
		list = list.Next
	}
	return stack
}

func isEmpty(stack []int) bool {
	return len(stack) == 0
}

func pop(stack *[]int) (int, bool) {
	if stack == nil {
		return 0, false
	}

	if len(*stack) == 0 {
		return 0, false
	}
	s := *stack
	popVal := s[len(s)-1]
	*stack = s[:len(s)-1]
	return popVal, true
}
