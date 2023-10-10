package tree

//二叉树的后序遍历(leetcode-145)
//给你一棵二叉树的根节点root，返回其节点值的后序遍历。

//示例1：输入：root=[1,null,2,3],输出：[3,2,1]
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

// 递归遍历(左右中)
func postorderTraversal1(root *TreeNode) []int {
	var (
		ans = make([]int, 0)
		f   func(root *TreeNode)
	)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		f(root.Left)
		f(root.Right)
		ans = append(ans, root.Val)
	}
	f(root)
	return ans
}

// 循环遍历(TODO)
func postorderTraversal2(root *TreeNode) []int {
	var (
		ans = make([]int, 0)
		//cur = root
	)

	//for cur != nil {
	//
	//}
	return ans
}
