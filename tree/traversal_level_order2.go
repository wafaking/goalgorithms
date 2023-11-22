package tree

import "goalgorithms/common"

//二叉树的层序遍历(leetcode-102)
//给你二叉树的根节点root，返回其节点值的层序遍历。（即逐层地，从左到右访问所有节点）。
//示例1：输入：root=[3,9,20,null,null,15,7]输出：[[3],[9,20],[15,7]]
//示例2：输入：root=[1]输出：[[1]]
//示例3：输入：root=[]输出：[]

// 层序遍历(使用队列)
func levelOrder21(root *common.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		ans      = make([][]int, 0)
		queue    = []*common.TreeNode{root}
		elements = make([]int, 0)
		count    = 1
		cur      *common.TreeNode
	)

	for {
		if count == 0 { // 当前层遍历完成
			ans = append(ans, elements)
			count = len(queue)
			elements = nil
			if len(queue) == 0 {
				break
			}
			continue
		}

		//当前层还没遍历完
		//将当前层元素添加到当前层结果集
		cur = queue[0]
		queue = queue[1:]
		elements = append(elements, cur.Val)
		count--

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return ans
}

// 层序遍历(TODO)
func levelOrder22(root *common.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		ret   = make([][]int, 0)
		queue = []*common.TreeNode{root}
	)

	for i := 0; len(queue) > 0; i++ {
		ret = append(ret, []int{})
		var p = make([]*common.TreeNode, 0)
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		queue = p
	}
	return ret
}
