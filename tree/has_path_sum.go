package tree

//路径总和(leetcode-112)
//给定二叉树的根节点root和目标和的整数targetSum。判断该树中是否存在根节点到叶子节点的路径，
//这条路径上所有节点值相加等于目标和targetSum。如果存在，返回true；否则，返回false。
//叶子节点是指没有子节点的节点。
//示例1：输入：root=[5,4,8,11,null,13,4,7,2,null,null,null,1],targetSum=22输出：true
//	解释：等于目标和的根节点到叶节点路径如上图所示。
//示例2：输入：root=[1,2,3],targetSum=5输出：false,解释：
//	树中存在两条根节点到叶子节点的路径：
//	(1-->2):和为3
//	(1-->3):和为4
//不存在sum=5的根节点到叶子节点的路径。
//示例3：输入：root=[],targetSum=0输出：false
//解释：由于树是空的，所以不存在根节点到叶子节点的路径。

// 深度优先遍历
func hasPathSum1(root *TreeNode, targetSum int) bool {

	var f func(root *TreeNode) bool
	f = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		// 减去当前节点的值
		targetSum -= root.Val

		// 是叶子节点，判断和是不是满足targetSum
		if root.Left == nil && root.Right == nil {
			if targetSum == 0 { // 满足则结束
				return true
			}
			// 不满足则回溯
			targetSum += root.Val
			return false
		}

		// 不是叶子节点，则继续处理左、右子节点
		var okLeft, okRight bool
		okLeft = f(root.Left)
		if okLeft { // 若左子节点已满足条件，则直接返回
			return true
		}

		okRight = f(root.Right)

		// 回溯
		targetSum += root.Val

		return okRight
	}
	return f(root)
}

// 递归深度优先(简便写法)
func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum2(root.Left, sum-root.Val) || hasPathSum2(root.Right, sum-root.Val)
}

// 广度优先遍历
func hasPathSum3(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var (
		nodeQueue  = []*TreeNode{root}
		valueQueue = []int{targetSum - root.Val}
	)
	for len(nodeQueue) > 0 {
		var (
			// curNode与curValue一一对应
			curNode  = nodeQueue[0]
			curValue = valueQueue[0]
		)
		nodeQueue = nodeQueue[1:]
		valueQueue = valueQueue[1:]

		// 叶子节点
		if curNode.Left == nil && curNode.Right == nil && curValue == 0 {
			return true
		}

		if curNode.Left != nil {
			nodeQueue = append(nodeQueue, curNode.Left)
			valueQueue = append(valueQueue, curValue-curNode.Left.Val)

		}
		if curNode.Right != nil {
			nodeQueue = append(nodeQueue, curNode.Right)
			valueQueue = append(valueQueue, curValue-curNode.Right.Val)
		}
	}
	return false
}
