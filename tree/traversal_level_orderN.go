package tree

import (
	"goalgorithms/common"
)

//N叉树的层序遍历(leetcode-429)
//给定一个N叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
//树的序列化输入是用层序遍历，每组子节点都由null值分隔（参见示例）。
//示例1：输入：root=[1,null,3,2,4,null,5,6]输出：[[1],[3,2,4],[5,6]]
//示例2：输入：root=[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
//输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]

// 广度优先遍历
func levelOrder1(root *common.NTreeNode) [][]int {
	var ans = make([][]int, 0)
	if root == nil {
		return ans
	}
	var (
		queue   = []*common.NTreeNode{root}
		count   = 1
		tempAns = make([]int, 0)
	)
	for {
		if count == 0 {
			ans = append(ans, tempAns)
			tempAns = nil
			count = len(queue)
			if count == 0 {
				break
			}
			continue
		}
		count--
		cur := queue[0]
		queue = queue[1:]
		tempAns = append(tempAns, cur.Val)
		for _, v := range cur.Children {
			queue = append(queue, v)
		}
	}

	return ans
}

// 广度优先
func levelOrder2(root *common.NTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		ans   = make([][]int, 0)
		queue = []*common.NTreeNode{root}
	)
	for len(queue) > 0 {
		var (
			levelVal   = make([]int, 0)
			levelQueue = make([]*common.NTreeNode, 0)
		)
		// 将当前queue中元素处理完
		for _, cur := range queue {
			levelVal = append(levelVal, cur.Val)
			for _, v := range cur.Children {
				levelQueue = append(levelQueue, v)
			}
		}

		ans = append(ans, levelVal)
		queue = append([]*common.NTreeNode{}, levelQueue...)
		levelQueue = nil
	}

	return ans
}

// 广度优先
func levelOrder3(root *common.NTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		ans   = make([][]int, 0)
		queue = []*common.NTreeNode{root}
	)
	for len(queue) > 0 {
		var (
			tempQueue = queue
			levelVal  = make([]int, 0)
		)

		queue = nil // 清空并存储下一层节点
		for _, cur := range tempQueue {
			levelVal = append(levelVal, cur.Val)
			queue = append(queue, cur.Children...)
		}
		ans = append(ans, levelVal)
	}

	return ans
}

//func levelOrder(root *Node) (ans [][]int) {
//	if root == nil {
//		return
//	}
//	q := []*Node{root}
//	for q != nil {
//		level := []int{}
//		tmp := q
//		q = nil
//		for _, node := range tmp {
//			level = append(level, node.Val)
//			q = append(q, node.Children...)
//		}
//		ans = append(ans, level)
//	}
//	return
//}
