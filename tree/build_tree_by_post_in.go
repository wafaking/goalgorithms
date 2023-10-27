package tree

// 从后序序与中序遍历序列构造二叉树(leetcode-106)
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
Output: [3,9,20,null,null,15,7]
*/

// 递归法
func buildTree21(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	var (
		root  = &TreeNode{Val: postorder[len(postorder)-1]}
		pivot int
	)
	// 查找根节点在中序序列中的位置
	for ; pivot < len(inorder); pivot++ {
		if inorder[pivot] == root.Val {
			break
		}
	}

	root.Left = buildTree21(inorder[:pivot], postorder[:pivot])
	root.Right = buildTree21(inorder[pivot+1:], postorder[pivot:len(postorder)-1])
	return root
}

// 递归(双指针)
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
		//log.Println("lI, rI, postOrder:", inOrderLIndex, inOrderRIndex, postorder)
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

// 迭代法 todo
func buildTree24(inorder []int, postorder []int) *TreeNode {
	//if len(inorder) == 0 || len(postorder) == 0 {
	//	return nil
	//}
	//var (
	//	postIdx = len(postorder) - 1
	//	root    = &TreeNode{Val: postorder[postIdx]}
	//)
	//for postIdx >= 0 {
	//	var idx int
	//	for ; idx < len(inorder); idx++ {
	//		if inorder[idx] == postorder[postIdx] {
	//			break
	//		}
	//	}
	//	if idx == 1 {
	//
	//	}
	//	break
	//}
	//return root
	return nil
}

// 使用循环(TODO-----)
func buildTree23(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	stack := []*TreeNode{root}
	idx := len(inorder) - 1

	// post [2 1 3 6 7 5 10 9 8 4]
	//   in [1 2 3 4 5 6 7 8 9 10]

	//中序遍历 inorder = [9,3,15,20,7]
	//后序遍历 postorder = [9,15,7,20,3]
	for i := len(postorder) - 2; i >= 0; i-- {
		postVal := postorder[i]
		curNode := stack[len(stack)-1]

		if curNode.Val != inorder[idx] { // 等于当前值
			curNode.Right = &TreeNode{Val: postVal}
			stack = append(stack, curNode.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[idx] {
				curNode = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				idx--
			}
			curNode.Left = &TreeNode{Val: postVal}
			stack = append(stack, curNode.Left)
		}
	}
	return root
}
