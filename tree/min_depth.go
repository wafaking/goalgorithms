package tree

import "goalgorithms/common"

//二叉树的最小深度(leetcode-111)
//给定一个二叉树，找出其最小深度。
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//说明：叶子节点是指没有子节点的节点。
//示例1：输入：root=[3,9,20,null,null,15,7]输出：2
//示例2：输入：root=[2,null,3,null,4,null,5,null,6]输出：5

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 深度遍历(后序遍历)
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var getMinDepth func(root *TreeNode) int
	getMinDepth = func(root *TreeNode) int {
		if root == nil || (root.Left == nil && root.Right == nil) {
			return 0
		}

		// 注：要找到叶子节点到根节点的最短路径
		if root.Left != nil && root.Right == nil {
			return getMinDepth(root.Left) + 1
		} else if root.Right != nil {
			return getMinDepth(root.Right) + 1
		}
		return common.MinInTwo(getMinDepth(root.Left), getMinDepth(root.Right)) + 1
	}
	return getMinDepth(root) + 1
}

// 宽度遍历(后序遍历)
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		queue = []*TreeNode{root}
		ans   = 1
	)
	for len(queue) > 0 {
		var sz = len(queue)
		for sz > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left == nil && cur.Right == nil {
				return ans
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			sz--
		}
		ans++
	}

	return ans
}
