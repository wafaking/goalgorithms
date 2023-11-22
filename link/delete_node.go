package link

import "goalgorithms/common"

//删除链表中的节点(leetcode-237)
//请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为要被删除的节点。
//注：链表中节点的值唯一且给定节点不是属尾节点；
//示例1：输入：head=[4,5,1,9],node=5输出：[4,1,9]
//解释：给定你链表中值为5的第二个节点，那么在调用了你的函数之后，该链表应变为4->1->9.
//示例2：输入：head=[4,5,1,9],node=1输出：[4,5,9]
//解释：给定你链表中值为1的第三个节点，那么在调用了你的函数之后，该链表应变为4->5->9.

func deleteNode(node *common.ListNode) {
	if node.Next == nil {
		return
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	return
}
