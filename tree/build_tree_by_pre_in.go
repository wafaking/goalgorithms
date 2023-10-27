package tree

//从前序与中序遍历序列构造二叉树(leetcode-105/sword-07)
// 根据前序遍历和中序遍历结果，重建该二叉树(假设输入的前序遍历和中序遍历的结果中都不含重复的数字)。

/*
例如，给出
前序遍历:preorder = [3,9,20,15,7]
中序遍历:inorder = [9,3,15,20,7]
返回如下的二叉树：
		 3
		/ \
	   9  20
		 /  \
		15   7
Output: [3,9,20,null,null,15,7]
*/

// 递归
func buildTree11(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	var (
		root = &TreeNode{Val: preorder[0]}
		idx  int // 根节点在中序序列中的位置
	)

	// 查找根节点在中序序列中的位置
	for ; idx < len(inorder); idx++ {
		if inorder[idx] == preorder[0] {
			break
		}
	}
	//idx同时也表明了中序序列中根节点前有idx个左节点，
	//因此在前序序列中根节点(第一个)的后面的idx个元素都是左节点
	root.Left = buildTree11(preorder[1:1+idx], inorder[:idx])
	root.Right = buildTree11(preorder[1+idx:], inorder[idx+1:])
	return root
}

// TODO ---迭代法
func buildTree12(preorder []int, inorder []int) *TreeNode {
	return nil
}

func buildTree13(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// [1,2,4,7,3,5,6,8]
	// [4,7,2,1,5,3,8,6]
	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{}
	stack = append(stack, root)
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}
