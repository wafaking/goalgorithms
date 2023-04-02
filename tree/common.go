package tree

import (
	"fmt"
)

var root *TreeNode

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 添加节点
func Insert(num int) {
	node := &TreeNode{Val: num}
	if root == nil {
		root = node
		return
	}
	// 原sli：4, 3, 8, 1, 2, 5, 7, 9, 6, 10,
	// pre 4 3 1 2 8 5 7 6 9 10
	// in 1 2 3 4 5 6 7 8 9 10
	// post 2 1 3 6 7 5 10 9 8 4

	cur := root
	for {
		if num < cur.Val {
			if cur.Left == nil {
				cur.Left = node
				return
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = node
				return
			}
			cur = cur.Right
		}
	}
}

// 前序遍历的递归实现(先根遍历)
func PreOrder1(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)

	PreOrder1(root.Left)
	PreOrder1(root.Right)
	return
}

// 前序遍历的递归实现(先根遍历)
func PreOrder2(root *TreeNode) []int {
	var (
		preorder func(root *TreeNode)
		sli      []int
	)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		sli = append(sli, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return sli
}

// 前序遍历的循环实现(先根遍历)
func PreOrder3(root *TreeNode) []int {
	var (
		head  = root
		stack []*TreeNode // 1.初始化栈
		sli   []int
	)

	for head != nil || len(stack) > 0 {
		if head != nil { // 处理当前节点
			sli = append(sli, head.Val) // 2.输出根节点
			stack = append(stack, head) // 3.将当前节点放入栈中
			head = head.Left            // 4. 处理左子树
		} else { // 处理栈中元素
			head = stack[len(stack)-1] // 栈顶元素(后进先出)
			stack = stack[:len(stack)-1]
			head = head.Right
		}
	}
	return sli
}

// 中序遍历(先左子树，再根节点，再右子树)
func InOrder1(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder1(root.Left)
	fmt.Println(root.Val)
	InOrder1(root.Right)
}

// 中序遍历(先左子树，再根节点，再右子树)
func InOrder2(root *TreeNode) []int {
	var (
		sli     []int
		inOrder func(root *TreeNode)
	)

	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		sli = append(sli, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)
	return sli
}

// 中序遍历(先左子树，再根节点，再右子树)
func InOrder3(root *TreeNode) []int {
	var (
		sli   []int
		stack []*TreeNode
		head  = root
	)

	for head != nil || len(stack) > 0 {
		if head != nil {
			stack = append(stack, head)
			head = head.Left
		} else {
			head = stack[len(stack)-1]
			sli = append(sli, head.Val)
			stack = stack[:len(stack)-1]
			head = head.Right
		}
	}
	return sli
}

// 后序遍历（先左子树，再右子树，最后根节点）
func PostOrder1(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder1(root.Left)
	PostOrder1(root.Right)
	fmt.Println(root.Val)
}

// 后序遍历（先左子树，再右子树，最后根节点）
func PostOrder2(root *TreeNode) []int {
	var (
		sli       []int
		postOrder func(root *TreeNode)
	)
	postOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postOrder(root.Left)
		postOrder(root.Right)
		sli = append(sli, root.Val)
	}
	postOrder(root)
	return sli
}

// 后序遍历（先左子树，再右子树，最后根节点）
//增加辅助保存上一次打印结果数组的节点，
//当一个节点左右都是空的时候，就可以放入结果集
//当上一个放入结果集的节点是他的孩子节点的时候，
//注意:
// 当节点的左右不为空时， 要先加入右孩子，再加入左孩子，这样才能先访问左孩子。
func PostOrder3(root *TreeNode) []int {
	var (
		sli   []int
		stack []*TreeNode
		pre   *TreeNode //辅助节点存储上一次打印值的节点
	)
	if root == nil {
		return sli
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		// 可以打印分两种情况：
		// 1. 当前节点是叶子节点；
		// 2. 上一个打印的节点有值，且上一个节点是当前节点的左子节点或右子节点(即当前节点的子节点已经打印过)
		if (cur.Left == nil && cur.Right == nil) ||
			(pre != nil && (pre == cur.Left || pre == cur.Right)) {
			sli = append(sli, cur.Val)
			pre = cur
			stack = stack[:len(stack)-1]
		} else {
			// 这里添加，先添加右子节点，后添加左子节点
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}
	return sli
}

// 层序遍历(宽度优先，借助队列实现)
func LevelOrder1(root *TreeNode) []int {
	var (
		sli   []int
		queue []*TreeNode
	)
	if root == nil {
		return sli
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0] // 先进先出
		sli = append(sli, cur.Val)
		queue = queue[1:]

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return sli
}
