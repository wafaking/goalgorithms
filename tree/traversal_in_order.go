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
func inOrderTraversal1(root *TreeNode) []int {
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

// 循环遍历
func inOrderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		ans   = make([]int, 0)
		cur   = root
		stack = make([]*TreeNode, 0)
	)

	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left // 左
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, cur.Val) // 中
			cur = cur.Right            // 右
		}
	}
	return ans
}

func inorderTraversal3(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left != nil {
			// predecessor 节点表示当前 root 节点向左走一步，然后一直向右走至无法走为止的节点
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				// 有右子树且没有设置过指向 root，则继续向右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				// 将 predecessor 的右指针指向 root，这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
				predecessor.Right = root
				// 遍历左子树
				root = root.Left
			} else { // predecessor 的右指针已经指向了 root，则表示左子树 root.Left 已经访问完了
				res = append(res, root.Val)
				// 恢复原样
				predecessor.Right = nil
				// 遍历右子树
				root = root.Right
			}
		} else { // 没有左子树
			res = append(res, root.Val)
			// 若有右子树，则遍历右子树
			// 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
			root = root.Right
		}
	}
	return
}
