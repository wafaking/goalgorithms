package tree

import "goalgorithms/common"

var (
	//root               *common.TreeNode
	defaultNullTreeVal = -99999
)

// BuildBinaryTree 构建二叉树(使用队列)
func BuildBinaryTree(nums []int) *common.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if nums[0] == defaultNullTreeVal {
		return nil
	}

	var (
		root  = &common.TreeNode{Val: nums[0]}
		queue = []*common.TreeNode{root}
	)

	for i := 1; i < len(nums) && len(queue) > 0; i++ {
		var (
			r           = queue[0]
			left, right *common.TreeNode
		)
		queue = queue[1:]
		if nums[i] != defaultNullTreeVal {
			left = &common.TreeNode{Val: nums[i]}
			r.Left = left
			queue = append(queue, left)
		}
		i++

		if i == len(nums) {
			continue
		}

		if nums[i] != defaultNullTreeVal {
			right = &common.TreeNode{Val: nums[i]}
			r.Right = right
			queue = append(queue, right)
		}
	}

	return root
}

// Insert 添加节点
func Insert(root *common.TreeNode, num int) {
	var node *common.TreeNode
	if num == defaultNullTreeVal {
		node = nil
	} else {
		node = &common.TreeNode{Val: num}
	}
	if root == nil {
		root = node
		return
	}
	// 原sli：4, 3, 8, 1, 2, 5, 7, 9, 6, 10,
	// pre 4 3 1 2 8 5 7 6 9 10
	// in 1 2 3 4 5 6 7 8 9 10
	// post 2 1 3 6 7 5 10 9 8 4

	cur := root
	for {
		if num < cur.Val {
			if cur.Left == nil {
				cur.Left = node
				return
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = node
				return
			}
			cur = cur.Right
		}
	}
}
