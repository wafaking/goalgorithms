package tree

import (
	"strconv"
	"strings"
)

//二叉树的所有路径(leetcode-257)
//给你一个二叉树的根节点root，按任意顺序，返回所有从根节点到叶子节点的路径。
//叶子节点是指没有子节点的节点。
//示例1：输入：root=[1,2,3,null,5]输出：["1->2->5","1->3"]
//示例2：输入：root=[1]输出：["1"]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 回溯+递归(前序遍历，深度优先遍历)
func binaryTreePaths1(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	var (
		numList = make([]string, 0)
		ans     = make([]string, 0)
		f       func(root *TreeNode)
	)

	f = func(root *TreeNode) {
		if root == nil {
			return
		}

		// 根节点先添加到记录中
		numList = append(numList, strconv.Itoa(root.Val))

		// 到达叶子节点，将这条路径添加到结果集中
		if root.Left == nil && root.Right == nil {
			var temp = make([]string, len(numList))
			copy(temp, numList)
			ans = append(ans, strings.Join(temp, "->"))
			return
		}

		// 子子节点不为空，则将结果添加到路径中，处理完成后回溯，将左节点弹出
		if root.Left != nil {
			f(root.Left)
			numList = numList[:len(numList)-1]
		}
		if root.Right != nil {
			f(root.Right)
			numList = numList[:len(numList)-1]
		}
	}
	f(root)
	return ans
}

// 深度优先搜索(节省内存版)
func binaryTreePaths2(root *TreeNode) []string {
	var (
		ans            = make([]string, 0)
		constructPaths func(root *TreeNode, path string)
	)

	constructPaths = func(root *TreeNode, path string) {
		if root != nil {
			pathSB := path
			pathSB += strconv.Itoa(root.Val)
			if root.Left == nil && root.Right == nil {
				ans = append(ans, pathSB)
			} else {
				pathSB += "->"
				constructPaths(root.Left, pathSB)
				constructPaths(root.Right, pathSB)
			}
		}
	}

	constructPaths(root, "")
	return ans
}

//宽度优先遍历
func binaryTreePaths3(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	var (
		ans      = make([]string, 0)
		queue    = []*TreeNode{root}
		ansQueue = []string{strconv.Itoa(root.Val)} // 存放临时结果,也queue一一对应
	)
	for len(queue) > 0 {
		var (
			curNode  = queue[0]
			curValue = ansQueue[0]
		)

		queue = queue[1:]
		ansQueue = ansQueue[1:]
		if curNode.Left == nil && curNode.Right == nil {
			ans = append(ans, curValue)
			continue
		}

		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
			ansQueue = append(ansQueue, curValue+"->"+strconv.Itoa(curNode.Left.Val))
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
			ansQueue = append(ansQueue, curValue+"->"+strconv.Itoa(curNode.Right.Val))
		}

	}

	return ans
}

//宽度优先遍历(节省内存版)
func binaryTreePaths4(root *TreeNode) []string {
	var paths = make([]string, 0)
	if root == nil {
		return paths
	}
	nodeQueue := []*TreeNode{}
	pathQueue := []string{}
	nodeQueue = append(nodeQueue, root)
	pathQueue = append(pathQueue, strconv.Itoa(root.Val))

	for i := 0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			continue
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	return paths
}
