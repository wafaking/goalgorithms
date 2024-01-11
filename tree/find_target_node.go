package tree

import "goalgorithms/common"

//寻找二叉搜索树中的目标节点(leetcode-LCR174)
//某公司组织架构以二叉搜索树形式记录，节点值为处于该职位的员工编号。请返回第cnt大的员工编号。
//示例1：输入：root=[7,3,9,1,5],cnt=2输出：7
//示例2：输入:root=[10,5,15,2,7,null,20,1,null,6,8],cnt=4输出:8
//提示： 1 ≤ cnt ≤ 二叉搜索树元素个数

func findTargetNode1(root *common.TreeNode, cnt int) int {
	var (
		list     = make([]int, 0)
		preOrder func(root *common.TreeNode)
	)
	preOrder = func(root *common.TreeNode) {
		if root == nil {
			return
		}
		preOrder(root.Left)
		list = append(list, root.Val)
		preOrder(root.Right)
	}
	preOrder(root)

	if len(list) < cnt {
		return -1
	}
	return list[len(list)-cnt]
}
