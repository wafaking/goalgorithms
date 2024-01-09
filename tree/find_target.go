package tree

import "goalgorithms/common"

//两数之和IV-输入二叉搜索树(leetcode-653)
//给定一个二叉搜索树root和一个目标结果k，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回true。
//示例1：输入:root=[5,3,6,2,4,null,7],k=9输出:true
//示例2：输入:root=[5,3,6,2,4,null,7],k=28输出:false
//提示:
//	二叉树的节点个数的范围是  [1, 104].
//	-104 <= Node.val <= 104
//	题目数据保证，输入的 root 是一棵 有效 的二叉搜索树
//	-105 <= k <= 105

// 深度优先搜索+哈希表
func findTarget1(root *common.TreeNode, k int) bool {
	var (
		m   = make(map[int]struct{}, 0)
		dfs func(root *common.TreeNode) bool
	)
	dfs = func(root *common.TreeNode) bool {
		if root == nil {
			return false
		}
		if _, ok := m[k-root.Val]; ok {
			return true
		}
		m[root.Val] = struct{}{}
		return dfs(root.Left) || dfs(root.Right)
	}
	return dfs(root)
}

// 广度优先搜索+哈希表
func findTarget2(root *common.TreeNode, k int) bool {
	var (
		m     = make(map[int]struct{}, 0)
		queue = []*common.TreeNode{root}
	)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if _, ok := m[k-cur.Val]; ok {
			return true
		}
		m[cur.Val] = struct{}{}
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return false
}

//深度优先搜索+中序遍历+双指针
func findTarget3(root *common.TreeNode, k int) bool {
	var (
		list             = make([]int, 0)
		inorderTraversal func(root *common.TreeNode)
	)
	inorderTraversal = func(root *common.TreeNode) {
		if root == nil {
			return
		}
		inorderTraversal(root.Left)
		list = append(list, root.Val)
		inorderTraversal(root.Right)
	}
	inorderTraversal(root)

	for i, j := 0, len(list)-1; i < j; {
		sum := list[i] + list[j]
		if sum == k {
			return true
		} else if sum > k {
			j--
		} else {
			i++
		}
	}

	return false
}
