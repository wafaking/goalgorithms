package link

import "goalgorithms/common"

// 从尾到头打印链表(sword-6)
// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
// 示例1：输入：head = [1,3,2], 输出：[2,3,1]

//法一： 顺序遍历，将元素添加至栈中，再将栈中元素输出；
func reversePrint1(head *common.ListNode) []int {
	var (
		dummy = head
		sli   []int
	)
	for dummy != nil {
		sli = append(sli, dummy.Val)
		dummy = dummy.Next
	}

	if len(sli) <= 1 {
		return sli
	}
	n := len(sli)
	for i := 0; i < n>>1+n&1; i++ { // 注意此处需要判别奇偶数
		sli[i], sli[n-1-i] = sli[n-1-i], sli[i]
	}
	//交换方法2
	//var start, end = 0, n - 1
	//for start < end {
	//	sli[start], sli[end] = sli[end], sli[start]
	//	start++
	//	end--
	//}
	// 或将链表反转再输出
	return sli
}

//法二：使用defer打印(正确)
func reversePrint2(head *common.ListNode) []int {
	var (
		cur        = head
		ans        []int
		handleNode func(node *common.ListNode)
	)
	handleNode = func(cur *common.ListNode) {
		for cur != nil {
			defer func(v int) {
				ans = append(ans, v)
			}(cur.Val)
			cur = cur.Next
		}
	}
	handleNode(cur)
	return ans
}

// 法四：使用递归
func reversePrint3(head *common.ListNode) []int {
	var (
		cur       = head
		printList = func(node *common.ListNode) {}
		ans       []int
	)
	printList = func(cur *common.ListNode) {
		if cur == nil {
			return
		}
		var v = cur.Val
		cur = cur.Next
		printList(cur)
		ans = append(ans, v)
	}
	printList(cur)
	return ans
}

// reversePrint4 递归
func reversePrint4(head *common.ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(reversePrint4(head.Next), head.Val)
}
