package tree

import "goalgorithms/common"

//二叉树的最大深度(leetcode-104)
//给定一个二叉树root，返回其最大深度。
//二叉树的最大深度是指从根节点到最远叶子节点的最长路径上的节点数。
//示例1：输入：root=[3,9,20,null,null,15,7],输出：3
//示例2：输入：root=[1,null,2],输出：2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 最大深度：从根节点到叶子节点的最长路径，使用前序遍历
// 最大高度：从叶子节点到根节点的最长路径，使用后序遍历

// 前序遍历(深度优先遍历)
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var getMaxDepth func(root *TreeNode) int
	getMaxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		var leftDepth, rightDepth int
		if root.Left != nil {
			leftDepth = getMaxDepth(root.Left) + 1
		}
		if root.Right != nil {
			rightDepth = getMaxDepth(root.Right) + 1
		}
		return common.MaxInTwo(leftDepth, rightDepth)
	}

	return getMaxDepth(root) + 1
}

// 前序遍历(深度优先)
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var getMaxDepth func(root *TreeNode) int
	getMaxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		var leftDepth, rightDepth int
		if root.Left != nil {
			leftDepth = getMaxDepth(root.Left)
		}
		if root.Right != nil {
			rightDepth = getMaxDepth(root.Right)
		}
		return common.MaxInTwo(leftDepth, rightDepth) + 1
	}

	return common.MaxInTwo(getMaxDepth(root.Left), getMaxDepth(root.Right)) + 1
}

// 宽度优先遍历
func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}
