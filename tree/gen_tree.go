package tree

// GenerateTree 根据层序遍历结果的数组生成树，其中-1表示节点不存在
func GenerateTree(nums []int) *TreeNode {
	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}

	var (
		root  = &TreeNode{Val: nums[0]} // root节点
		treeQ = make([]*TreeNode, 0)
	)
	// 将root节点添加到队列头部
	treeQ = append(treeQ, root)

	// 因每次要创建左、右子节点，因此i+2
	for i := 1; i < len(nums); i += 2 {
		var leftNode *TreeNode
		if nums[i] != -1 { // 非空节点
			leftNode = &TreeNode{Val: nums[i]}
		}

		// 弹出当前节点
		cur := treeQ[0]
		treeQ = treeQ[1:]

		cur.Left = leftNode

		// 创建右节点
		if i+1 < len(nums) {
			var rightNode *TreeNode
			if nums[i+1] != -1 {
				rightNode = &TreeNode{Val: nums[i+1]}
			}
			cur.Right = rightNode
		}

		// 判断左右子节点有没有必要再添加子节点
		if cur.Left != nil {
			treeQ = append(treeQ, cur.Left)
		}
		if cur.Right != nil {
			treeQ = append(treeQ, cur.Right)
		}
	}

	return root
}
