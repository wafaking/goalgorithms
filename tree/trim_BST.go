package tree

import "goalgorithms/common"

//修剪二叉搜索树(leetcode-669)
//给你二叉搜索树的根节点root，同时给定最小边界low和最大边界high。通过修剪二叉搜索树，
//使得所有节点的值在[low,high]中。修剪树不应该改变保留在树中的元素的相对结构(如果没有被移除，原有的父代子代关系都应当保留)。
//可以证明，存在唯一的答案。 所以结果应当返回修剪好的二叉搜索树的新的根节点。
//注意，根节点可能会根据给定的边界发生改变。注：树中每个节点的值都是唯一的;
//示例1：输入：root=[1,0,2],low=1,high=2输出：[1,null,2]
//示例2：输入：root=[3,0,4,null,2,null,null,1],low=1,high=3输出：[3,2,null,1]

// 递归(前序)
func trimBST1(root *common.TreeNode, low int, high int) *common.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST1(root.Right, low, high)
	} else if root.Val > high {
		return trimBST1(root.Left, low, high)
	}
	root.Left = trimBST1(root.Left, low, high)
	root.Right = trimBST1(root.Right, low, high)
	return root
}

// 迭代
func trimBST2(root *common.TreeNode, low int, high int) *common.TreeNode {
	// 找到处于low与high之间的节点
	for {
		switch {
		case root == nil:
			return nil
		case root.Val < low:
			root = root.Right
			continue
		case root.Val > high:
			root = root.Left
			continue
		}
		break
	}

	for cur := root; cur.Left != nil; {
		// 如果cur.Left小于low,cur.Left.Left一定小于low，不需要再处理
		// 但cur.Left.Right仍需要处理，让其顶替cur.Left
		if cur.Left.Val < low {
			// 右子节点有可能满足条件,右子节点代替父级节点
			cur.Left = cur.Left.Right
		} else {
			//cur.Left.Right一定小于当前节点，因此不需要处理
			cur = cur.Left
		}
	}

	for cur := root; cur.Right != nil; {
		// 如果cur.Right大于high,则cur.Right.Right一定不满足条件
		if cur.Right.Val > high {
			// 用cur.Right.Left代替cur.Right节点
			cur.Right = cur.Right.Left
		} else {
			cur = cur.Right
		}
	}
	return root
}
