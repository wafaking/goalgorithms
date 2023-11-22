package tree

import "goalgorithms/common"

//最长同值路径(leetcode-687)
//给定一个二叉树的root，返回最长的路径的长度，这个路径中的每个节点具有相同值。这条路径可以经过也可以不经过根节点。
//两个节点之间的路径长度由它们之间的边数表示。
//示例1:输入：root=[5,4,5,1,1,5]输出：2
//示例2:输入：root=[1,4,5,4,4,5]输出：2

func longestUnivaluePath(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		ans int
		f   func(root *common.TreeNode) int
	)
	f = func(root *common.TreeNode) int {
		if root == nil {
			return 0
		}

		leftVal := f(root.Left)   // 左子节点相同值边数
		rightVal := f(root.Right) // 右子节点相同值边数

		// 记录当前左、右子节点相同值边数
		var curLeft, curRight int
		if root.Left != nil && root.Val == root.Left.Val {
			curLeft = leftVal + 1
		}

		if root.Right != nil && root.Val == root.Right.Val {
			curRight = rightVal + 1
		}

		// 当前节点相同值边数和与结果取最大值：4，4，4为两条边数
		ans = common.MaxInTwo(ans, curLeft+curRight)
		// 返回当前节点左、右子节点相同值边数的最大值，如：4，4，4则返回1
		return common.MaxInTwo(curLeft, curRight)
	}

	_ = f(root)
	return ans
}
