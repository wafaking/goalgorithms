package tree

import "goalgorithms/common"

//剑指 Offer 34. 二叉树中和为某一值的路径
//输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
//示例:
//给定如下二叉树，以及目标和 target = 22，
/*
		 5
		/ \
	   4   8
	  /   / \
	 11  13  4
	/  \    / \
   7    2  5   1

*/
//返回:
//	[
//		[5,4,11,2],
//		[5,8,4,5]
//	]

//func pathSum(root *common.TreeNode, target int) [][]int {
//	var res [][]int
//	if root == nil {
//		return res
//	}
//	var sum int
//	//for {
//	//	var elem []int
//	//}
//	if sum(root, sum) == target {
//
//	}
//
//	return res
//}
// TODO ----
func sum(node *common.TreeNode, sum int) int {
	if node == nil {
		return sum
	}
	return node.Val + sum
}
