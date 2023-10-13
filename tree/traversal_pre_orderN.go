package tree

//N叉树的前序遍历(leetcode-589)
//给定一个n叉树的根节点root，返回其节点值的前序遍历。
//n叉树在输入中按层序遍历进行序列化表示，每组子节点由空值null分隔（请参见示例）。
//示例1：输入：root=[1,null,3,2,4,null,5,6]输出：[1,3,5,6,2,4]
//示例2：输入：root=[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
//输出：[1,2,3,6,7,11,14,4,8,12,5,9,13,10]

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
		1
      /  |  \
	 3 	 2   4
	/ \
   2   4
*/

// TODO ----
func preorder(root *NTreeNode) []int {
	return []int{}
}
