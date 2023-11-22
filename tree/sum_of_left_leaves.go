package tree

import "goalgorithms/common"

//左叶子之和(leetcode-404)
//给定二叉树的根节点root，返回所有左叶子之和。
//示例1：输入:root=[3,9,20,null,null,15,7]输出:24
//解释:在这个二叉树中，有两个左叶子，分别是9和15，所以返回24
//示例2:输入:root=[1]输出:0

// 深度优先递归+前序遍历
func sumOfLeftLeaves1(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		ans int
		f   func(pre, cur *common.TreeNode)
	)
	f = func(pre, cur *common.TreeNode) {
		if cur == nil {
			return
		}
		// 中
		if cur.Left == nil && cur.Right == nil {
			if pre != nil && pre.Left == cur {
				ans += cur.Val
			}
		}

		f(cur, cur.Left)  // 左
		f(cur, cur.Right) // 右
	}
	f(nil, root)
	return ans
}

// 深度优先递归(后序遍历)
func sumOfLeftLeaves2(root *common.TreeNode) int {
	var (
		f func(root *common.TreeNode) int
	)

	f = func(root *common.TreeNode) int {
		if root == nil { // 当前节点不存在
			return 0
		} else if root.Left == nil && root.Right == nil {
			// 当前节点是叶子节点
			return 0
		}

		// 当前节点的左子节点是叶子节点
		leftLeave := f(root.Left) // 左
		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			leftLeave = root.Left.Val
		}

		rightLeave := f(root.Right)   //右
		return leftLeave + rightLeave //中
	}

	return f(root)
}

// 宽度优先遍历
func sumOfLeftLeaves3(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		queue = []*common.TreeNode{root}
		ans   int
	)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		// 当前节点的左子节点是叶子节点
		if cur.Left != nil && cur.Left.Left == nil && cur.Left.Right == nil {
			ans += cur.Left.Val
		}

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return ans
}
