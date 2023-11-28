package tree

import (
	"goalgorithms/common"
)

//路径总和III(leetcode-437)
//给定一个二叉树的根节点root，和一个整数targetSum，求该二叉树里节点值之和等于targetSum的路径的数目。
//路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//示例1：输入：root=[10,5,-3,3,2,null,11,3,-2,null,1],targetSum=8输出：3
//解释：和等于8的路径有3条，如图所示。
//示例2：输入：root=[5,4,8,11,null,13,4,7,2,null,null,5,1],targetSum=22输出：3

// 后序遍历
func pathSum31(root *common.TreeNode, targetSum int) int {
	var (
		ans       int
		postOrder func(root *common.TreeNode)
		// 记录每个节点可能的和的值
		nodeSumMap = make(map[*common.TreeNode][]int, 0)
	)

	postOrder = func(root *common.TreeNode) {
		if root == nil {
			return
		}
		postOrder(root.Left)
		postOrder(root.Right)
		if root.Val == targetSum {
			ans++
		}
		valList := []int{root.Val}

		if root.Left != nil {
			for _, v := range nodeSumMap[root.Left] {
				valList = append(valList, root.Val+v)
				if root.Val+v == targetSum {
					ans++
				}
			}
		}
		if root.Right != nil {
			for _, v := range nodeSumMap[root.Right] {
				valList = append(valList, root.Val+v)
				if root.Val+v == targetSum {
					ans++
				}
			}
		}
		nodeSumMap[root] = valList
	}
	postOrder(root)

	return ans
}

// 深度优先遍历
func pathSum32(root *common.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var (
		ans int
		f   func(root *common.TreeNode, targetSum int)
	)

	// 计算当前节点和为targetSum的数量
	f = func(root *common.TreeNode, targetSum int) {
		if root == nil {
			return
		}
		targetSum -= root.Val
		if targetSum == 0 {
			ans += 1
		}
		f(root.Left, targetSum)
		f(root.Right, targetSum)
	}

	// 获取当前节点的结果数
	f(root, targetSum)
	// 递归获取当前节点子节点的结果数
	ans += pathSum32(root.Left, targetSum)
	ans += pathSum32(root.Right, targetSum)
	return ans
}

// 前缀和(TODO)
func pathSum33(root *common.TreeNode, targetSum int) int {
	var ans int
	var (
		preSum = map[int]int{0: 1} // 默认与targetSum相等时值为1
		dfs    func(root *common.TreeNode, cur int)
	)
	dfs = func(root *common.TreeNode, cur int) {
		if root == nil {
			return
		}
		cur += root.Val
		ans += preSum[cur-targetSum]
		preSum[cur]++
		dfs(root.Left, cur)
		dfs(root.Right, cur)
		preSum[cur]--
	}
	dfs(root, 0)
	return ans
}
