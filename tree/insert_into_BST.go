package tree

import "goalgorithms/common"

//二叉搜索树中的插入操作(leetcode-701)
//给定二叉搜索树（BST）的根节点root和要插入树中的值value，将值插入二叉搜索树。返回插入后二叉搜索树的根节点。
//输入数据保证，新值和原始二叉搜索树中的任意节点值都不同。
//注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。你可以返回任意有效的结果。
//示例1：输入：root=[4,2,7,1,3],val=5输出：[4,2,7,1,3,5]
//解释：另一个满足题目要求可以通过的树是：[5,2,7,1,3,4]
//示例2：输入：root=[40,20,60,10,30,50,70],val=25输出：[40,20,60,10,30,50,70,null,null,25]
//示例3：输入：root=[4,2,7,1,3,null,null,null,null,null,null],val=5输出：[4,2,7,1,3,5]

// 循环
func insertIntoBST1(root *common.TreeNode, val int) *common.TreeNode {
	if root == nil {
		return &common.TreeNode{Val: val}
	}
	var cur = root
	for cur != nil {
		if val <= cur.Val {
			if cur.Left == nil {
				cur.Left = &common.TreeNode{Val: val}
				break
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = &common.TreeNode{Val: val}
				break
			}
			cur = cur.Right
		}
	}

	return root
}

// 递归
func insertIntoBST2(root *common.TreeNode, val int) *common.TreeNode {
	if root == nil {
		return &common.TreeNode{Val: val}
	}
	if val <= root.Val {
		root.Left = insertIntoBST2(root.Left, val)
	} else {
		root.Right = insertIntoBST2(root.Right, val)
	}
	return root
}
