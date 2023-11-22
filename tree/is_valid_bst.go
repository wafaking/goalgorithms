package tree

import (
	"goalgorithms/common"
	"math"
)

//验证二叉搜索树(leetcode-98)
//给你一个二叉树的根节点root，判断其是否是一个有效的二叉搜索树。
//有效二叉搜索树定义如下：
//	节点的左子树只包含小于当前节点的数。
//	节点的右子树只包含大于当前节点的数。
//	所有左子树和右子树自身必须也是二叉搜索树。
//示例1：输入：root=[2,1,3]输出：true
//示例2：输入：root=[5,1,4,null,null,3,6]输出：false,解释：
//	根节点的值是5，但是右子节点的值是4。

// 中序遍历(再判断有序数组)
func isValidBST1(root *common.TreeNode) bool {
	var (
		f    func(root *common.TreeNode) bool
		nums = make([]int, 0, 1) // 用来存储前一个节点的值
	)
	f = func(root *common.TreeNode) bool {
		if root == nil {
			return true
		}
		if !f(root.Left) {
			return false
		}
		if len(nums) == 0 {
			nums = append(nums, root.Val)
		} else if nums[0] > root.Val {
			return false
		} else if nums[0] == root.Val {
			return false
		} else if nums[0] < root.Val {
			nums[0] = root.Val
		}

		return f(root.Right)
	}

	return f(root)
}

// 中序遍历(记录前一个节点，方法同1)
func isValidBST2(root *common.TreeNode) bool {
	var (
		pre *common.TreeNode
		f   func(root *common.TreeNode) bool
	)

	// pre节点用于记录前一个节点的值
	f = func(root *common.TreeNode) bool {
		if root == nil {
			return true
		}

		if !f(root.Left) { // 左
			return false
		}

		if pre != nil && pre.Val >= root.Val { // 中
			return false
		}

		// 中
		pre = root
		return f(root.Right) // 右
	}

	return f(root)
}

// 前序遍历
func isValidBST3(root *common.TreeNode) bool {
	// left,right用于标记当前节点值的范围
	var f func(root *common.TreeNode, left, right int) bool

	f = func(root *common.TreeNode, left, right int) bool {
		if root == nil {
			return true
		}

		if root.Val <= left || root.Val >= right {
			return false
		}

		return f(root.Left, left, root.Val) && f(root.Right, root.Val, right)
	}

	return f(root, math.MinInt64, math.MaxInt64)
}

// 中序遍历(迭代实现)
func isValidBST4(root *common.TreeNode) bool {
	var (
		stack = make([]*common.TreeNode, 0)
		min   = math.MinInt64
	)
	for len(stack) > 0 || root != nil {
		// 找到最左叶子节点
		for root != nil {
			// 先把自身加进去
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= min {
			return false
		}
		min = root.Val

		//if root.Right != nil {
		//	stack = append(stack, root.Right)
		//}
		// 注：此处root要变化
		root = root.Right
	}
	return true
}
