package tree

type BinaryTreeNode struct {
	val    int
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
	Father *BinaryTreeNode
}

/*
		 3
		/ \
	   9  20
 	 	 /  \
		15   7
*/

//给定二叉树其中的一个结点，请找出此节点中序遍历的下一个结点并且返回。
//注意，树中的结点不仅包含左右子结点，而且包含指向父结点的指针。

func getNext(node *BinaryTreeNode) *BinaryTreeNode {
	// 判断当前节点有无右子节点
	// 1.有右子节点，next:右子节点的最左子节点(3:15, 20:7)
	// 2.无右子节点，分两种情况：
	// 2.1 当前节点为父节点的左子节点，next:父节点
	// 2.2 当前节点为父节点的右子节点，next:向上找父节点，
	//     直到父节点b为某个父节点a的左子树，则next为a,否则为nil；
	if node == nil {
		return nil
	}
	if node.Right != nil { // 1. 有右子节点
		nextNode := node.Right     // 找此右子节点的左子节点
		for nextNode.Left != nil { //
			nextNode = nextNode.Left
		}
		return nextNode
	}

	// 2. 无右子节点
	var pNode = node
	for pNode.Father != nil {
		if pNode.Father.Left == pNode { // 是父节点的左子节点，则返回父节点
			return pNode.Father
		}
		// 是父节点的右子节点，
		pNode = pNode.Father // 找第一个当前节点是父节点左孩子的节点
	}
	return nil // 退回到根节点仍未找到
}
