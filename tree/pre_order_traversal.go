package tree

//二叉树的前序遍历(leetcode-144)
//给你二叉树的根节点root，返回它节点值的前序遍历。
//示例1：输入：root=[1,null,2,3]输出：[1,2,3]
//示例2：输入：root=[]输出：[]
//示例3：输入：root=[1]输出：[1]
//示例4：输入：root=[1,2]输出：[1,2]
//示例5：输入：root=[1,null,2],输出：[1,2]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//前序遍历（根左右）

// 递归遍历
func preorderTraversal1(root *TreeNode) []int {
	var (
		ans = make([]int, 0)
		f   func(root *TreeNode)
	)

	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		f(root.Left)
		f(root.Right)
	}
	f(root)
	return ans
}

// 循环实现(使用栈先添加右子节点到栈中)
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		ans   = make([]int, 0)
		stack = make([]*TreeNode, 0)
	)

	var cur = root
	for cur != nil || len(stack) > 0 {
		if cur == nil {
			// 弹出栈顶元素
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		// 添加根节点元素
		ans = append(ans, cur.Val)

		// 先添加右子节点到栈中
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}

		// 顺着左子节点遍历
		if cur.Left != nil {
			cur = cur.Left
		} else {
			cur = nil
		}
	}

	return ans
}

// 循环遍历(添加当前节点到栈中)
func preorderTraversal3(root *TreeNode) []int {
	//一直沿着左子节点遍历，并将当前节点加入栈中,遍历到左子叶子节点时，再依次取出栈顶元素处理其右子节点
	var (
		cur   = root
		stack []*TreeNode // 1.初始化栈
		sli   []int
	)

	for cur != nil || len(stack) > 0 {
		if cur != nil { // 处理当前节点
			sli = append(sli, cur.Val) // 2.输出根节点
			stack = append(stack, cur) // 3.将当前节点放入栈中
			cur = cur.Left             // 4.沿着左子节点遍历
		} else { // 处理栈中元素
			cur = stack[len(stack)-1] // 栈顶元素(后进先出)
			stack = stack[:len(stack)-1]
			cur = cur.Right // 弹出栈顶元素后取其右子节点
		}
	}
	return sli
}
