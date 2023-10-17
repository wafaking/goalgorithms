package tree

import "sort"

//完全二叉树的节点个数(leetcode-222)
//给你一棵完全二叉树的根节点root，求出该树的节点个数。
//完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，
//	并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第h层，则该层包含1~2h个节点。
//示例1：输入：root=[1,2,3,4,5,6]输出：6
//示例2：输入：root=[]输出：0
//示例3：输入：root=[1]输出：1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 后序遍历
func countNodes1(root *TreeNode) int {
	var (
		count int
		f     func(root *TreeNode)
	)

	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		f(root.Left)
		f(root.Right)
		count++
	}
	f(root)
	return count
}

// 循环遍历(判断节点是否是满二叉树)
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var f func(root *TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		var (
			leftDepth, rightDepth int
			left                  = root.Left
			right                 = root.Right
		)
		for left != nil {
			left = left.Left
			leftDepth++
		}
		for right != nil {
			right = right.Right
			rightDepth++
		}
		// 当前节点是满二叉树的节点数量
		if leftDepth == rightDepth {
			return 2<<leftDepth - 1
		}
		// 当前节点不是满二叉树，使用后序遍历，分别求左、右、中节点数量
		return f(root.Left) + f(root.Right) + 1
	}
	return f(root)
}

// TODO(二分查找+位运算)
func countNodes3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}
