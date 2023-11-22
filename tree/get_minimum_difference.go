package tree

import (
	"goalgorithms/common"
	"math"
)

//二叉搜索树的最小绝对差(leetcode-530/783)
//给你一个二叉搜索树的根节点root，返回树中任意两不同节点值之间的最小差值。
//差值是一个正数，其数值等于两值之差的绝对值。注：节点值大于0
//示例1：输入：root=[4,2,6,1,3]输出：1
//示例2：输入：root=[1,0,48,null,null,12,49]输出：1

// 中序遍历
func getMinimumDifference1(root *common.TreeNode) int {
	// 由于是二叉搜索树，因此左子孙节点的值一定小于当前节点，右子孙节点的值一定大于当前节点
	// 因此，与当前节点差值最小的节点有几种情况：
	// 1. 与左、右子节点的差值：求左、中、右相邻差值即可;
	// 2. 与左子树的最右子节点的差值：左子树的最右节点的下一个节点一定是当前节点
	// 3. 与右子树的最左子节点的差值：右子树的最左节点的前一个元素一定是当前节点
	// 因此：中序遍历正好满足：

	var (
		f   func(root *common.TreeNode)
		min = math.MaxInt64
		pre = -1
	)

	f = func(root *common.TreeNode) {
		if root == nil {
			return
		}

		f(root.Left)
		if pre != -1 && min > root.Val-pre {
			min = root.Val - pre
		}
		pre = root.Val

		f(root.Right)
	}

	f(root)
	return min
}
