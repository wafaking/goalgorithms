package tree

import "goalgorithms/common"

//二叉树的层序遍历II(leetcode-107)
//给你二叉树的根节点root，返回其节点值自底向上的层序遍历。（即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//示例1：输入：root=[3,9,20,null,null,15,7]输出：[[15,7],[9,20],[3]]
//示例2：输入：root=[1]输出：[[1]]
//示例3：输入：root=[]输出：[]

// 宽度优先遍历(将正常结果逆序)
func levelOrderBottom1(root *common.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		ans      = make([][]int, 0)
		queue    = []*common.TreeNode{root}
		elements = make([]int, 0)
		cur      *common.TreeNode
		count    = 1
	)
	for {
		if count == 0 {
			ans = append(ans, elements)
			count = len(queue)
			elements = nil
			if len(queue) == 0 {
				break
			}
			continue
		}

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
	n := len(ans)
	for i := 0; i < n>>1; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return ans
}
