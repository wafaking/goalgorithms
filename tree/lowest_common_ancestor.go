package tree

import "goalgorithms/common"

//二叉树的最近公共祖先(leetcode-236)
//给定一个二叉树,找到该树中两个指定节点的最近公共祖先。
//最近公共祖先定义：对于有根树T的两个节点p、q，最近公共祖先表示为一个节点x，满足x是p、q的祖先且x的深度尽可能大(一个节点也可以是它自己的祖先)。
//注：树中节点数目在范围[2,10^5]内,且节点值不重复；
//示例1：输入：root=[3,5,1,6,2,0,8,null,null,7,4],p=5,q=1,输出：3
//解释：节点5和节点1的最近公共祖先是节点3。
//示例2：输入：root=[3,5,1,6,2,0,8,null,null,7,4],p=5,q=4,输出：5
//解释：节点5和节点4的最近公共祖先是节点5。因为根据定义最近公共祖先节点可以为节点本身。
//示例3：输入：root=[1,2],p=1,q=2输出：1

// 后序遍历(向上查找公共父节点)
func lowestCommonAncestor1(root, p, q *common.TreeNode) *common.TreeNode {
	var f func(root *common.TreeNode) *common.TreeNode
	f = func(root *common.TreeNode) *common.TreeNode {
		if root == nil {
			return nil
		} else if root.Val == p.Val || root.Val == q.Val {
			return root
		}

		leftNode := f(root.Left)
		rightNode := f(root.Right)

		if leftNode != nil && rightNode != nil {
			return root
		} else if leftNode != nil {
			return leftNode
		} else if rightNode != nil {
			return rightNode
		}

		return nil
	}

	return f(root)
}

// 存储父节点
func lowestCommonAncestor2(root, p, q *common.TreeNode) *common.TreeNode {
	var (
		m       = map[int]*common.TreeNode{root.Val: nil} // 存储节点的父节点
		visited = make(map[int]struct{}, 0)               // 存储p节点查找父节点的记录
		f       func(root *common.TreeNode)
	)
	f = func(root *common.TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			m[root.Left.Val] = root
		}
		if root.Right != nil {
			m[root.Right.Val] = root
		}
		f(root.Left)
		f(root.Right)
	}
	f(root)

	// 注：此题假设p与q在root中一定有公共祖先
	for p != nil {
		parent, _ := m[p.Val]
		// 记录节点p查找父节点的记录
		visited[p.Val] = struct{}{}
		p = parent
	}

	for q != nil {
		// 查找q节点的父节点
		if _, ok := visited[q.Val]; ok {
			return q
		} else {
			q = m[q.Val]
		}
	}
	return nil
}
