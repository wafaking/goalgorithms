package tree

import "goalgorithms/common"

//找树左下角的值(leetcode-513)
//给定一个二叉树的根节点root，请找出该二叉树的最底层最左边节点的值。
//假设二叉树中至少有一个节点。
//示例1:输入:root=[2,1,3]输出:1
//示例2:输入:[1,2,3,4,null,5,6,null,null,7]输出:7

// 宽度优先遍历(层序遍历)
func findBottomLeftValue1(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		queue = []*common.TreeNode{root}
		count = 1
		ans   = root.Val
	)
	for len(queue) > 0 {
		cur := queue[0]
		if count == 0 {
			ans = cur.Val
			count = len(queue)
		}

		queue = queue[1:]
		count--

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return ans
}

// 宽度优先遍历(层序遍历，先右后左)
func findBottomLeftValue2(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		queue = []*common.TreeNode{root}
		ans   = root.Val
	)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		//最后一个元素一定是最后一层最左侧元素
		ans = cur.Val

		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
	}
	return ans
}

// 深度优先遍历
func findBottomLeftValue3(root *common.TreeNode) int {
	var (
		maxHeight int
		ans       int
		f         func(root *common.TreeNode, curHeight int)
	)
	f = func(root *common.TreeNode, curHeight int) {
		if root == nil {
			return
		}

		curHeight++
		if curHeight > maxHeight {
			maxHeight = curHeight
			ans = root.Val
		}

		f(root.Left, curHeight)
		f(root.Right, curHeight)

		return
	}

	f(root, 0)
	return ans
}
