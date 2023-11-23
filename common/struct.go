package common

// ListNode 结构
type ListNode struct {
	Val  int
	Next *ListNode
}

type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
