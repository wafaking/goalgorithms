package tree

import "goalgorithms/common"

//二叉树的后序遍历(leetcode-145)
//给你一棵二叉树的根节点root，返回其节点值的后序遍历。

//示例1：输入：root=[1,null,2,3],输出：[3,2,1]
//示例2：输入：root=[],输出：[]
//示例3：输入：root=[1],输出：[1]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *common.TreeNode
 *     Right *common.TreeNode
 * }
 */

// 递归遍历(左右中)
func postOrderTraversal1(root *common.TreeNode) []int {
	var (
		ans = make([]int, 0)
		f   func(root *common.TreeNode)
	)
	f = func(root *common.TreeNode) {
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

// 循环遍历(左右中)：前序遍历再逆序
func postOrderTraversal2(root *common.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var (
		ans   = make([]int, 0)
		cur   = root
		stack = []*common.TreeNode{root}
	)

	// 1.使用前序遍历，但顺序为:中右左
	for len(stack) > 0 {
		if cur != nil {
			ans = append(ans, cur.Val) // 中
			stack = append(stack, cur)
			cur = cur.Right // 右
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = cur.Left // 左
		}
	}

	//2.数组逆序
	n := len(ans)
	for i := 0; i < (n >> 1); i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return ans
}

// postOrderTraversal3 后序遍历(左右根)
//增加辅助保存上一次打印结果数组的节点，
//当一个节点左右都是空的时候，就可以放入结果集
//当上一个放入结果集的节点是他的孩子节点的时候，
//注意:
// 当节点的左右不为空时， 要先加入右孩子，再加入左孩子，这样才能先访问左孩子。
func postOrderTraversal3(root *common.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		sli   []int
		stack = []*common.TreeNode{root}
		pre   *common.TreeNode //辅助节点存储上一次打印值的节点
	)

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		// 可以打印分两种情况：
		// 1. 当前节点是叶子节点；
		// 2. 上一个打印的节点有值，且上一个节点是当前节点的左子节点或右子节点(即当前节点的子节点已经打印过)
		if (cur.Left == nil && cur.Right == nil) || (pre != nil && (pre == cur.Left || pre == cur.Right)) {
			sli = append(sli, cur.Val)
			pre = cur
			stack = stack[:len(stack)-1]
		} else {
			// 这里添加，先添加右子节点，后添加左子节点
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}
	return sli
}
