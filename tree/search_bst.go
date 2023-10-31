package tree

//二叉搜索树中的搜索(leetcode-700)
//给定二叉搜索树（BST）的根节点root和一个整数值val。
//你需要在BST中找到节点值等于val的节点。返回以该节点为根的子树。如果节点不存在，则返回null。
//示例1:输入：root=[4,2,7,1,3],val=2输出：[2,1,3]
//示例2:输入：root=[4,2,7,1,3],val=5输出：[]

// 深度遍历
func searchBST1(root *TreeNode, val int) *TreeNode {
	var f func(root *TreeNode) *TreeNode
	f = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val == val {
			return root
		} else if root.Val > val {
			return f(root.Left)
		} else {
			return f(root.Right)
		}
	}

	return f(root)
}

// 广度优先
func searchBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Val == val {
			return cur
		} else if cur.Val > val && cur.Left != nil {
			queue = append(queue, cur.Left)
		} else if cur.Val < val && cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return nil
}

// 迭代
func searchBST3(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		}
	}
	return nil
}
