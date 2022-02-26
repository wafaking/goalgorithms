package link

// GetLinkLength1 获取链表长
func GetLinkLength1(node *ListNode) int {
	var length int
	for node != nil {
		length++
		node = node.Next
	}
	return length
}

// GetLinkLength2 获取链表长度
func GetLinkLength2(node *ListNode) int {
	if node == nil {
		return 0
	}
	return 1 + GetLinkLength2(node.Next)
}
