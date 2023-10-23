package tree

//路径总和(leetcode-112)
//给你二叉树的根节点root和一个表示目标和的整数targetSum。判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和targetSum。如果存在，返回true；否则，返回false。
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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TODO ----
func hasPathSum(root *TreeNode, targetSum int) bool {

	return false
}
