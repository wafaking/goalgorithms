package tree

import "goalgorithms/common"

//二叉搜索树中的众数(leetcode-501)
//给你一个含重复值的二叉搜索树（BST）的根节点root，找出并返回BST中的所有众数（即，出现频率最高的元素）。
//如果树中有不止一个众数，可以按任意顺序返回。注：树中节点的数目在范围[1, 10^4] 内;
//假定BST满足如下定义：
//	结点左子树中所含节点的值小于等于当前节点的值
//	结点右子树中所含节点的值大于等于当前节点的值
//	左子树和右子树都是二叉搜索树
//示例1：输入：root=[1,null,2,2]输出：[2]
//示例2：输入：root=[0]输出：[0]

// 中序遍历（记录前一个节点的指针）
func findMode1(root *common.TreeNode) []int {
	var (
		pre             *common.TreeNode //用于记录前一个节点的值
		count, maxCount int
		f               func(root *common.TreeNode)
		ans             = make([]int, 0)
	)

	f = func(root *common.TreeNode) {
		if root == nil {
			return
		}

		f(root.Left)
		if pre == nil {
			// 直接添加当前节点
			ans = append(ans, root.Val)
			count = 1
			maxCount = 1
		} else {

			if root.Val == pre.Val {
				count++
			} else {
				count = 1 // 计数器重置
			}

			if count == maxCount {
				ans = append(ans, root.Val)
			} else if count > maxCount {
				maxCount = count
				ans = []int{pre.Val}
			}

		}

		pre = root
		f(root.Right)
	}
	f(root)
	return ans
}

// 中序遍历（记录前一个节点的值）
func findMode2(root *common.TreeNode) []int {
	var (
		ans                   = make([]int, 0)
		f                     func(root *common.TreeNode)
		base, count, maxCount int
	)

	f = func(root *common.TreeNode) {
		if root == nil {
			return
		}
		f(root.Left)

		if root.Val == base {
			count++
		} else {
			count = 1
			base = root.Val
		}

		if count > maxCount {
			maxCount = count
			ans = []int{root.Val}
		} else if count == maxCount {
			ans = append(ans, root.Val)
		}

		f(root.Right)
	}
	f(root)
	return ans
}
