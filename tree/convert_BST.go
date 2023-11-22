package tree

//把二叉搜索树转换为累加树(leetcode-538/1038)
//给出二叉搜索树的根节点，该树的节点值各不相同，请你将其转换为累加树（GreaterSumTree），
//使每个节点node的新值等于原树中大于或等于node.val的值之和,
//即将它的每个节点的值替换成树中大于或者等于该节点值的所有节点值之和
//提醒:
//	节点的左子树仅包含键小于节点键的节点。
//	节点的右子树仅包含键大于节点键的节点。
//	左右子树也必须是二叉搜索树。
//示例1：输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
//输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
//示例2：输入：root=[0,null,1]输出：[1,null,1]
//示例3：输入：root=[1,0,2]输出：[3,3,2]
//示例4：输入：root=[3,2,4,1]输出：[7,9,4,10]

// 先右节点的中序遍历
func convertBST1(root *TreeNode) *TreeNode {
	var (
		f   func(root *TreeNode)
		pre *TreeNode
	)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		f(root.Right)
		if pre != nil {
			root.Val += pre.Val
		}
		pre = root
		f(root.Left)
	}
	f(root)
	return root
}

// Morris 遍历(TODO)
func convertBST2(root *TreeNode) *TreeNode {
	sum := 0
	node := root
	for node != nil {
		if node.Right == nil {
			sum += node.Val
			node.Val = sum
			node = node.Left
		} else {
			succ := getSuccessor(node)
			if succ.Left == nil {
				succ.Left = node
				node = node.Right
			} else {
				succ.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}
		}
	}
	return root
}

func getSuccessor(node *TreeNode) *TreeNode {
	succ := node.Right
	for succ.Left != nil {
		if succ.Left == node {
			break
		}
		succ = succ.Left
	}
	return succ
}
