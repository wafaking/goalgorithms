package tree

//对称二叉树(leetcode-101)
//给你一个二叉树的根节点root，检查它是否轴对称。
//示例1：输入：root=[1,2,2,3,4,4,3]输出：true
//示例2：输入：root=[1,2,2,null,3,null,3]输出：false

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
		 4
		/ \
	   2   7
	  / \  / \
	 1   3 6  9
*/

// 反转二叉树，再比对(TODO)
func isSymmetric1(root *TreeNode) bool {
	return false
}

// 层序遍历后，判断每层结果
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var (
		queue            = []*TreeNode{root}
		count            = 1
		cur              *TreeNode
		elements         = make([]int, 0)
		isSymmetricArray func(list []int) bool
	)

	isSymmetricArray = func(list []int) bool {
		var n = len(list)
		for i := 0; i < n>>1; i++ {
			if list[i] != list[n-1-i] {
				return false
			}
		}
		return true
	}

	for {
		if count == 0 {
			count = len(queue)
			if !isSymmetricArray(elements) {
				return false
			}
			elements = nil
			if len(queue) == 0 {
				break
			}
			continue
		}

		cur = queue[0]
		queue = queue[1:]
		if cur == nil { // 对称，不能忽略为nil的节点
			elements = append(elements, defaultNullTreeVal)
			count--
			continue
		}

		elements = append(elements, cur.Val)
		count--

		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
	}

	return true
}

// 递归
func isSymmetric3(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var cmp func(l, r *TreeNode) bool
	cmp = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		} else if l == nil && r != nil {
			return false
		} else if l != nil && r == nil {
			return false
		} else if l.Val != r.Val {
			return false
		}

		// l!=nil && r!=nil
		return cmp(l.Left, r.Right) && cmp(l.Right, r.Left)
	}

	return cmp(root.Left, root.Right)
}
