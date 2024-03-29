package tree

import "goalgorithms/common"

// levelOrder11 层序遍历(宽度优先，借助队列实现)
func levelOrder11(root *common.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var (
		sli   = make([]int, 0)
		queue = make([]*common.TreeNode, 0)
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0] // 先进先出
		queue = queue[1:]
		if cur == nil {
			sli = append(sli, defaultNullTreeVal)
			continue
		}

		sli = append(sli, cur.Val)

		if cur.Left == nil && cur.Right == nil {
			continue
		}
		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
	}
	return sli
}
