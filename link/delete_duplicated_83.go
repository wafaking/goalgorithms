package link

// 83. Remove Duplicates from Sorted List  删除有序链表中的重复元素
// 解题思路：相同的只改变指针指向，不同的才会移动当前的cur指针，cur作为当前判断的指针
// Input: 1->1->2         Output: 1->2
// Input: 1->1->2->3->3     Output: 1->2->3
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head // cur指向头结点，而且改变cur指向不会影响到head的指向
	for cur.Next != nil {
		if cur.Val == cur.Next.Val { // 当前节点的值等于下个节点的值
			cur.Next = cur.Next.Next // cur指向下下个节点（cur指针不会移动）
		} else { // 当前节点的值不等于下个节点的值
			cur = cur.Next // 当前指针后移到下一个不同值的节点
		}
	}
	return head
}
