package tree

import "goalgorithms/common"

//N叉树的前序遍历(leetcode-589)
//给定一个n叉树的根节点root，返回其节点值的前序遍历。
//n叉树在输入中按层序遍历进行序列化表示，每组子节点由空值null分隔（请参见示例）。
//示例1：输入：root=[1,null,3,2,4,null,5,6]输出：[1,3,5,6,2,4]
//示例2：输入：root=[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
//输出：[1,2,3,6,7,11,14,4,8,12,5,9,13,10]

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
		1
      /  |  \
	 3 	 2   4
	/ \
   2   4
*/

// 递归
func preorder1(root *common.Node) []int {
	var (
		ans = make([]int, 0)
		f   func(root *common.Node)
	)
	f = func(root *common.Node) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		for i := 0; i < len(root.Children); i++ {
			f(root.Children[i])
		}
	}
	f(root)
	return ans
}

// 迭代(使用栈)
func preorder2(root *common.Node) []int {
	if root == nil {
		return []int{}
	}
	var (
		ans   = make([]int, 0)
		stack = []*common.Node{root}
	)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, cur.Val)
		for i := len(cur.Children) - 1; i >= 0; i-- {
			stack = append(stack, cur.Children[i])
		}
	}
	return ans
}

// TODO 迭代(使用map记录访问过的历史)
func preorder3(root *common.Node) (ans []int) {
	if root == nil {
		return
	}
	st := make([]*common.Node, 0)
	nextIndex := map[*common.Node]int{}
	node := root
	for len(st) > 0 || node != nil {
		for node != nil {
			ans = append(ans, node.Val)
			st = append(st, node)
			if len(node.Children) == 0 {
				break
			}
			nextIndex[node] = 1
			node = node.Children[0]
		}
		node = st[len(st)-1]
		i := nextIndex[node]
		if i < len(node.Children) {
			nextIndex[node] = i + 1
			node = node.Children[i]
		} else {
			st = st[:len(st)-1]
			delete(nextIndex, node)
			node = nil
		}
	}
	return
}
