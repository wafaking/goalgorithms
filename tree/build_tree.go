package tree

import "log"

// leetcode 105. 从前序与中序遍历序列构造二叉树(同sword 07重建二叉树)
// 根据前序遍历和中序遍历结果，重建该二叉树(假设输入的前序遍历和中序遍历的结果中都不含重复的数字)。
/*
例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
		 3
		/ \
	   9  20
		 /  \
		15   7
*/
// Output: [3,9,20,null,null,15,7]

func buildTree11(preorder []int, inorder []int) *TreeNode {
	var (
		root     *TreeNode
		endIndex int
	)
	if len(preorder) == 0 || len(inorder) == 0 {
		return root
	}

	// [1,2,4,7,3,5,6,8]
	// [4,7,2,1,5,3,8,6]
	root = &TreeNode{Val: preorder[0]}
	for i := 0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			endIndex = i // 获取中序根节点的索引位置
			break
		}
	}
	root.Left = buildTree11(preorder[1:len(inorder[:endIndex])+1], inorder[:endIndex])
	root.Right = buildTree11(preorder[len(inorder[:endIndex])+1:], inorder[endIndex+1:]) // 注意此处preorder的索引
	return root
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

// 从后序序与中序遍历序列构造二叉树(leetcode106)
// 根据后序遍历和中序遍历结果，构造二叉树(假设输入的前序遍历和中序遍历的结果中都不含重复的数字)。
/*
例如，给出
中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：
		 3
		/ \
	   9  20
		 /  \
		15   7
*/
// Output: [3,9,20,null,null,15,7

func buildTree21(inorder []int, postorder []int) *TreeNode {

	var (
		root     *TreeNode
		endIndex int
	)
	if len(inorder) == 0 || len(postorder) == 0 {
		return root
	}

	// post [2 1 3 6 7 5 10 9 8 4]
	//   in [1 2 3 4 5 6 7 8 9 10]

	root = &TreeNode{Val: postorder[len(postorder)-1]}
	for i := 0; i < len(inorder); i++ { // 找出后序根节点在中序序列中的位置
		if postorder[len(postorder)-1] == inorder[i] {
			endIndex = i
			break
		}
	}

	root.Left = buildTree21(inorder[:endIndex], postorder[:endIndex])
	root.Right = buildTree21(inorder[endIndex+1:], postorder[endIndex:len(postorder)-1])

	return root
}

func buildTree22(inorder []int, postorder []int) *TreeNode {
	var (
		idxMap    = map[int]int{}
		treeBuild func(leftIndex, rightIndex int) *TreeNode
	)
	// post [2 1 3 6 7 5 10 9 8 4]
	//   in [1 2 3 4 5 6 7 8 9 10]

	for k, v := range postorder {
		idxMap[v] = k // 节点元素不会有重复
	}
	treeBuild = func(inOrderLIndex, inOrderRIndex int) *TreeNode {
		if inOrderLIndex > inOrderRIndex { // 无剩余节点
			return nil
		}
		log.Println("lI, rI, postOrder:", inOrderLIndex, inOrderRIndex, postorder)
		val := postorder[len(postorder)-1]       // 根节点值
		postorder = postorder[:len(postorder)-1] // 注：每次要将上次的根节点去除
		root := &TreeNode{Val: val}
		i := idxMap[val]

		// 注：根据val的值将中序遍历划分成左、右子树
		// 由于每次都会取postorder的不末尾元素，因此要先遍历右子树，再遍历左子树
		root.Right = treeBuild(i+1, inOrderRIndex)
		root.Left = treeBuild(inOrderLIndex, i-1)
		return root
	}
	return treeBuild(0, len(inorder)-1)
}

// 使用循环
func buildTree23(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	stack := []*TreeNode{root}
	inorderIndex := len(inorder) - 1

	// post [2 1 3 6 7 5 10 9 8 4]
	//   in [1 2 3 4 5 6 7 8 9 10]

	for i := len(postorder) - 2; i >= 0; i-- {
		postorderVal := postorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Right = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex--
			}
			node.Left = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Left)
		}
	}
	return root
}
