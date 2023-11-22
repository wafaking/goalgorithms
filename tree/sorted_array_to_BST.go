package tree

import (
	"goalgorithms/common"
	"math/rand"
	"time"
)

//将有序数组转换为二叉搜索树(leetcode-108)
//给你一个整数数组nums，其中元素已经按升序排列，请你将其转换为一棵高度平衡二叉搜索树。
//高度平衡二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过1」的二叉树。
//示例1：输入：nums=[-10,-3,0,5,9]输出：[0,-3,9,-10,null,5]
//解释：[0,-10,5,null,-3,null,9]也将被视为正确答案：
//示例2：输入：nums=[1,3]输出：[3,1]
//解释：[1,null,3]和[3,1]都是高度平衡二叉搜索树。

// 中序遍历(选择中间位置作为根节点)
func sortedArrayToBST1(nums []int) *common.TreeNode {

	// l,r为左闭右开区间
	var f func(l, r int) *common.TreeNode
	f = func(l, r int) *common.TreeNode {
		if l >= r {
			return nil
		}
		var (
			mid  = l + (r-l)>>1
			root = &common.TreeNode{Val: nums[mid]}
		)
		root.Left = f(l, mid)
		root.Right = f(mid+1, r)
		return root
	}

	return f(0, len(nums))
}

// 中序遍历(随机选择中间左边或右边位置作为根节点)
func sortedArrayToBST2(nums []int) *common.TreeNode {
	rand.Seed(time.Now().UnixNano())
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *common.TreeNode {
	if left > right {
		return nil
	}

	// 选择任意一个中间位置数字作为根节点
	mid := (left + right + rand.Intn(2)) / 2
	root := &common.TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}
