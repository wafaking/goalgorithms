package tree

import "goalgorithms/common"

//平衡二叉树(leetcode-110)
//给定一个二叉树，判断它是否是高度平衡的二叉树。
//一棵高度平衡二叉树定义为：一个二叉树每个节点的左右两个子树的高度差的绝对值不超过1。
//示例1：输入：root=[3,9,20,null,null,15,7],输出：true
//示例2：输入：root=[1,2,2,3,3,null,null,4,4],输出：false
//示例3：输入：root=[]输出：true

// 自下而上递归(后序遍历)
func isBalanced1(root *common.TreeNode) bool {
	if root == nil {
		return true
	}

	var getHeight func(root *common.TreeNode) int
	getHeight = func(root *common.TreeNode) int {
		if root == nil {
			return 0
		}
		// 分别获取左、右子节点的最大深度
		leftHeight := getHeight(root.Left)
		rightHeight := getHeight(root.Right)

		// 不平衡则返回-1
		if leftHeight == -1 || rightHeight == -1 || (common.Abs(leftHeight-rightHeight) > 1) {
			return -1
		}

		// 当前节点平衡
		// 取当前节点的最大深度(即左、右子节点的最大深度+1)
		return common.MaxInTwo(leftHeight, rightHeight) + 1
	}

	//高度!=-1表明是平衡的
	return getHeight(root) >= 0
}

// 自上而下递归(前序遍历)
func isBalanced2(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	var getHeight func(root *common.TreeNode) int

	getHeight = func(root *common.TreeNode) int {
		if root == nil {
			return 0
		}
		leftHeight := getHeight(root.Left)
		rightHeight := getHeight(root.Right)
		return common.MaxInTwo(leftHeight, rightHeight) + 1
	}

	// 根、左、右
	return common.Abs(getHeight(root.Left)-getHeight(root.Right)) <= 1 && isBalanced2(root.Left) && isBalanced2(root.Right)
}
