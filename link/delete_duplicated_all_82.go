package link

func deleteDuplicatesAll(head *ListNode) *ListNode {

	cur := head

	if cur == nil {
		return nil
	}
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head
}
