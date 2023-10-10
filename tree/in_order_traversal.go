package tree

//二叉树的中序遍历(leetcode-94)
//给定一个二叉树的根节点root，返回它的中序遍历。
//示例1：输入：root=[1,null,2,3],输出：[1,3,2]
//示例2：输入：root=[],输出：[]
//示例3：输入：root=[1],输出：[1]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归遍历(左中右)
func inorderTraversal1(root *TreeNode) []int {
	var (
		ans = make([]int, 0)
		f   func(root *TreeNode)
	)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		f(root.Left)
		ans = append(ans, root.Val)
		f(root.Right)
	}
	f(root)
	return ans
}

// 循环遍历 (TODO)
func inorderTraversal2(root *TreeNode) []int {
	var (
		ans = make([]int, 0)
	)
	return ans
}
