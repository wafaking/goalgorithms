package tree

import (
	"goalgorithms/common"
)

//N叉树的后序遍历(leetcode-590)
//给定一个n叉树的根节点root，返回其节点值的后序遍历。
//n叉树在输入中按层序遍历进行序列化表示，每组子节点由空值null分隔（请参见示例）。
//示例1：输入：root=[1,null,3,2,4,null,5,6]输出：[5,6,3,2,4,1]
//示例2：输入：root=[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
//输出：[2,6,14,11,7,3,12,8,4,13,9,10,5,1]

// 递归
func postorder1(root *common.NTreeNode) []int {
	var (
		ans = make([]int, 0)
		f   func(root *common.NTreeNode)
	)
	f = func(root *common.NTreeNode) {
		if root == nil {
			return
		}
		for _, child := range root.Children {
			f(child)
		}
		ans = append(ans, root.Val)
	}
	f(root)
	return ans
}

// TODO -- 迭代
func postorder2(root *common.NTreeNode) []int {
	var ans = make([]int, 0)
	if root == nil {
		return nil
	}
	var (
		stack   = make([]*common.NTreeNode, 0)
		cur     = root
		nextIdx = make(map[*common.NTreeNode]int, 0)
	)
	for cur != nil || len(stack) > 0 {
		// 将每一层的首个节点添加到栈中
		for cur != nil {
			stack = append(stack, cur)
			if len(cur.Children) == 0 {
				break
			}
			nextIdx[cur] = 1
			cur = cur.Children[0]
			continue
		}
		cur = stack[len(stack)-1]
		idx := nextIdx[cur]
		if idx < len(cur.Children) {
			nextIdx[cur] = idx + 1
			cur = cur.Children[idx]
		} else {
			ans = append(ans, cur.Val)
			stack = stack[:len(stack)-1]
			delete(nextIdx, cur)
			cur = nil
		}

	}
	return ans
}

// 迭代
func postorder3(root *common.NTreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		ans   = make([]int, 0)
		stack = []*common.NTreeNode{root}
		// 记录节点的子节点是否被添加到栈里
		visited = make(map[*common.NTreeNode]struct{}, 0)
	)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]

		// 当前子节点是否被添加过 || 是否是叶子节点
		if _, ok := visited[cur]; ok || len(cur.Children) == 0 {
			ans = append(ans, cur.Val)
			stack = stack[:len(stack)-1]
			continue
		}

		for i := len(cur.Children) - 1; i >= 0; i-- {
			stack = append(stack, cur.Children[i])
		}
		visited[cur] = struct{}{}
	}

	return ans
}

// 迭代(前序遍历逆序-右左中)
func postorder4(root *common.NTreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		ans        = make([]int, 0)
		stack      = []*common.NTreeNode{root}
		reverSlice = func(ans []int) {
			n := len(ans) - 1
			for i := 0; i <= n/2; i++ {
				ans[i], ans[n-i] = ans[n-i], ans[i]
			}
		}
	)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, cur.Val)
		stack = append(stack, cur.Children...)
	}

	reverSlice(ans)
	return ans
}
