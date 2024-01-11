package tree

import "goalgorithms/common"

//前序遍历构造二叉搜索树(leetcode-1008)
//给定一个整数数组，它表示BST(即二叉搜索树)的先序遍历，构造树并返回其根。
//保证对于给定的测试用例，总是有可能找到具有给定需求的二叉搜索树。
//二叉搜索树是一棵二叉树，其中每个节点，Node.left的任何后代的值严格小于Node.val,Node.right的任何后代的值严格大于Node.val。
//二叉树的前序遍历首先显示节点的值，然后遍历Node.left，最后遍历Node.right。
//示例1：输入：preorder=[8,5,1,7,10,12]输出：[8,5,10,1,7,null,12]
//示例2:输入:preorder=[1,3]输出:[1,null,3]
//提示：
//	1<=preorder.length<=100
//	1<=preorder[i]<=10^8
//	preorder中的值互不相同

func bstFromPreorder1(preorder []int) *common.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 左闭右闭
	var dfs func(l, r int) *common.TreeNode
	dfs = func(l, r int) *common.TreeNode {
		if l > r {
			return nil
		}

		root := &common.TreeNode{Val: preorder[l]}
		var pivot = l + 1
		for ; pivot <= r; pivot++ {
			if preorder[pivot] > preorder[l] {
				break
			}
		}
		root.Left = dfs(l+1, pivot-1)
		root.Right = dfs(pivot, r)
		return root
	}
	return dfs(0, len(preorder)-1)
}
