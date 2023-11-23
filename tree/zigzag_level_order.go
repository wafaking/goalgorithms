package tree

import (
	"goalgorithms/common"
)

//二叉树的锯齿形层序遍历(leetcode-103)
//给你二叉树的根节点root，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//示例1：输入：root=[3,9,20,null,null,15,7]输出：[[3],[20,9],[15,7]]
//示例2：输入：root=[1]输出：[[1]]
//示例3：输入：root=[]输出：[]

// 层序遍历+偶层逆序
func zigzagLevelOrder1(root *common.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		ans        = make([][]int, 0)
		queue      = []*common.TreeNode{root}
		count      int
		reverseSli = func(sli []int) {
			n := len(sli) - 1
			for i := 0; i <= (n >> 1); i++ {
				sli[i], sli[n-i] = sli[n-i], sli[i]
			}
		}
	)
	for len(queue) > 0 {
		var (
			levelVal = make([]int, 0)
			size     = len(queue)
		)

		for i := 0; i < size; i++ {
			var cur = queue[0]
			queue = queue[1:]
			levelVal = append(levelVal, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		if count&1 == 1 {
			reverseSli(levelVal)
		}

		ans = append(ans, levelVal)
		count++
	}

	return ans
}
