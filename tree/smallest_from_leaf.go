package tree

import "goalgorithms/common"

//从叶结点开始的最小字符串(leetcode-988)
//给定一颗根结点为root的二叉树，树中的每一个结点都有一个[0,25]范围内的值，分别代表字母'a'到'z'。
//返回按字典序最小的字符串，该字符串从这棵树的一个叶结点开始，到根结点结束。
//注：字符串中任何较短的前缀在字典序上都是较小的：
//例如，在字典序上"ab"比"aba"要小。叶结点是指没有子结点的结点。
//节点的叶节点是没有子节点的节点。
//示例1：输入：root=[0,1,2,3,4,3,4]输出："dba"
//示例2：输入：root=[25,1,3,1,3,0,2]输出："adz"
//示例3：输入：root=[2,2,1,null,1,0,null,0]输出："abc"

func smallestFromLeaf1(root *common.TreeNode) string {
	return ""
}
