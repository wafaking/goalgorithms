package tree

import "goalgorithms/common"

//二叉搜索树的最近公共祖先(leetcode-235)
//给定一个二叉搜索树,找到该树中两个指定节点的最近公共祖先。
//最近公共祖先：对于有根树T的两个结点p、q，最近公共祖先表示为一个结点x，满足x是p、q的祖先且x的深度尽可能大（一个节点也可以是它自己的祖先）。
//注：所有节点的值都是唯一的。p、q 为不同节点且均存在于给定的二叉搜索树中。
//示例1:输入:root=[6,2,8,0,4,7,9,null,null,3,5],p=2,q=8,输出:6
//解释:节点2和节点8的最近公共祖先是6。
//示例2:输入:root=[6,2,8,0,4,7,9,null,null,3,5],p=2,q=4,输出:2
//解释:节点2和节点4的最近公共祖先是2,因为根据定义最近公共祖先节点可以为节点本身。

// 后序遍历
func lowestCommonAncestorBST1(root, p, q *common.TreeNode) *common.TreeNode {
	var f func(root *common.TreeNode) *common.TreeNode
	f = func(root *common.TreeNode) *common.TreeNode {
		if root == nil {
			return nil
		}

		// 当前节点比p、q都小，走左子节点
		if root.Val < p.Val && root.Val < q.Val {
			return f(root.Right)
		} else if root.Val > p.Val && root.Val > q.Val {
			// 当前节点比p、q都大，走右子节点
			return f(root.Left)
		}
		// 当前节点可能要不等于p或q，要么介于两者中间，一定是公共祖先节点
		return root
	}
	return f(root)
}

// 两次遍历(获取父节点列表再逐一匹配最末相等的父节点)
func lowestCommonAncestorBST2(root, p, q *common.TreeNode) *common.TreeNode {

	var (
		// 列出当前节点的所有的
		listAncestors = func(node *common.TreeNode) []*common.TreeNode {
			var (
				cur   = root
				queue = []*common.TreeNode{root}
			)
			for cur.Val != node.Val {
				if node.Val < cur.Val {
					cur = cur.Left
				} else {
					cur = cur.Right
				}
				queue = append(queue, cur)
			}
			return queue
		}
		queueP   = listAncestors(p)
		queueQ   = listAncestors(q)
		ancestor *common.TreeNode
	)
	for i := 0; i < len(queueQ) && i < len(queueP); i++ {
		// 找到最后一个相同的节点,就是公共祖先
		if queueP[i] == queueQ[i] {
			ancestor = queueP[i]
		}
	}

	return ancestor
}

// 一次遍历
func lowestCommonAncestorBST3(root, p, q *common.TreeNode) *common.TreeNode {
	for {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
}
