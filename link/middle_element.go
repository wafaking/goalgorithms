package link

// 876. Middle of the Linked List   找到链表的中间点
// 解题思路：快慢指针，从head开始跑，快指针结束时，返回慢指针
// Input: 1->2->3->4->5     Output: 3
// Input: 1->2->3->4->5->6   Output: 4

// 给定非空单链表，返回中间节点元素。如果有两个中间节点，返回第二个中间节点。
// 如：输入：[1,2,3,4,5]，返回3 [3,4,5])
// 	输入：[1,2,3,4,5,6]，返回4 [4,5,6](此链表有两个中间节点3和4,只返回第二个4；

// MiddleNode 获取链表中间节点
func MiddleNode(head *ListNode) int {
	var slow, fast = head, head
	// 可以添加虚拟头部元素
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 每次移动一个
		fast = fast.Next.Next // 每次移动两个
	}
	return slow.Val
}
