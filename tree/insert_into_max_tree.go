package tree

import "goalgorithms/common"

//最大二叉树II(leetcode-998)
//最大树定义：一棵树，并满足：其中每个节点的值都大于其子树中的任何其他值。
//给你最大树的根节点root和一个整数val。
//就像之前的问题那样，给定的树是利用Construct(a)例程从列表a（root=Construct(a)）递归地构建的：
//如果a为空，返回null。
//否则，令a[i]作为a的最大元素。创建一个值为a[i]的根节点root。
//root的左子树将被构建为Construct([a[0],a[1],...,a[i-1]])。
//root的右子树将被构建为Construct([a[i+1],a[i+2],...,a[a.length-1]])。
//返回root。
//请注意，题目没有直接给出a，只是给出一个根节点root=Construct(a)。
//假设b是a的副本，并在末尾附加值val。题目数据保证b中的值互不相同。
//返回Construct(b)。
//示例1：输入：root=[4,1,3,null,null,2],val=5输出：[5,4,null,1,3,null,null,2]
//解释：a=[1,4,2,3],b=[1,4,2,3,5]
//示例2：输入：root=[5,2,4,null,1],val=3输出：[5,2,4,null,1,null,3]
//解释：a=[2,1,5,4],b=[2,1,5,4,3]
//示例3：输入：root=[5,2,3,null,1],val=4输出：[5,2,4,null,1,3]
//解释：a=[2,1,5,3],b=[2,1,5,3,4]

// TODO ---
func insertIntoMaxTree(root *common.TreeNode, val int) *common.TreeNode {

	return nil
}
