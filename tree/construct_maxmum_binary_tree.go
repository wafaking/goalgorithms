package tree

//最大二叉树(leetcode-654)
//给定一个不重复的整数数组nums。最大二叉树可以用下面的算法从nums递归地构建:
//创建一个根节点，其值为nums中的最大值。
//递归地在最大值左边的子数组前缀上构建左子树。
//递归地在最大值右边的子数组后缀上构建右子树。
//返回nums构建的最大二叉树。
//示例1：输入：nums=[3,2,1,6,0,5]输出：[6,3,5,null,2,0,null,null,1]
//解释：递归调用如下所示：
//	-[3,2,1,6,0,5]中的最大值是6，左边部分是[3,2,1]，右边部分是[0,5]。
//	-[3,2,1]中的最大值是3，左边部分是[]，右边部分是[2,1]。
//	-空数组，无子节点。
//	-[2,1]中的最大值是2，左边部分是[]，右边部分是[1]。
//	-空数组，无子节点。
//	-只有一个元素，所以子节点是一个值为1的节点。
//	-[0,5]中的最大值是5，左边部分是[0]，右边部分是[]。
//	-只有一个元素，所以子节点是一个值为0的节点。
//	-空数组，无子节点。
//示例2：输入：nums=[3,2,1]输出：[3,null,2,null,1]

// 传递数组
func constructMaximumBinaryTree1(nums []int) *TreeNode {
	var f func(nums []int) *TreeNode
	f = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		var idx int
		for i := 1; i < len(nums); i++ {
			if nums[i] > nums[idx] {
				idx = i
			}
		}
		var root = &TreeNode{Val: nums[idx]}
		root.Left = f(nums[:idx])
		root.Right = f(nums[idx+1:])
		return root
	}

	return f(nums)
}

// 传递指针
func constructMaximumBinaryTree2(nums []int) *TreeNode {
	// l<r 左闭右开区间
	var f func(l, r int) *TreeNode
	f = func(l, r int) *TreeNode {
		if l >= r {
			return nil
		}
		var idx = l // 记录最大值位置
		for i := l + 1; i < r; i++ {
			if nums[i] > nums[idx] {
				idx = i
			}
		}

		var root = &TreeNode{Val: nums[idx]}
		root.Left = f(l, idx)
		root.Right = f(idx+1, r)
		return root
	}

	return f(0, len(nums))
}
