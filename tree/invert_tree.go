package tree

import "goalgorithms/common"

//翻转二叉树(leetcode-226)
//给你一棵二叉树的根节点root，翻转这棵二叉树，并返回其根节点。
//示例1：输入：root=[4,2,7,1,3,6,9]输出：[4,7,2,9,6,3,1]
//示例2：输入：root=[2,1,3]输出：[2,3,1]
//示例3：输入：root=[]输出：[]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *common.TreeNode
 *     Right *common.TreeNode
 * }
 */

/*
		 4
		/ \
	   2   7
	  / \  / \
	 1   3 6  9

		 4
		/ \
	   7   2
	  / \  / \
	 9  6  3  1
*/

// 递归遍历
func invertTree1(root *common.TreeNode) *common.TreeNode {
	var f func(cur *common.TreeNode)
	f = func(cur *common.TreeNode) {
		if cur == nil || (cur.Left == nil && cur.Right == nil) {
			return
		}
		cur.Left, cur.Right = cur.Right, cur.Left
		f(cur.Left)
		f(cur.Right)
	}

	f(root)
	return root
}

// 迭代法(TODO)
func invertTree2(root *common.TreeNode) *common.TreeNode {
	return nil
}
