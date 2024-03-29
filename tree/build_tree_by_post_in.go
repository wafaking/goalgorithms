package tree

import "goalgorithms/common"

// 从中序与后序遍历序列构造二叉树(leetcode-106)
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
func buildTree21(inorder []int, postorder []int) *common.TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	var (
		root  = &common.TreeNode{Val: postorder[len(postorder)-1]}
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

// 使用循环(TODO-----)
func buildTree22(inorder []int, postorder []int) *common.TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &common.TreeNode{Val: postorder[len(postorder)-1]}
	stack := []*common.TreeNode{root}
	idx := len(inorder) - 1

	//中序遍历 inorder = [9,3,15,20,7]
	//后序遍历 postorder = [9,15,7,20,3]
	for i := len(postorder) - 2; i >= 0; i-- {
		postVal := postorder[i]
		curNode := stack[len(stack)-1]

		if curNode.Val != inorder[idx] { // 等于当前值
			curNode.Right = &common.TreeNode{Val: postVal}
			stack = append(stack, curNode.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[idx] {
				curNode = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				idx--
			}
			curNode.Left = &common.TreeNode{Val: postVal}
			stack = append(stack, curNode.Left)
		}
	}
	return root
}
