package tree

// 给定一个二叉树，返回它的右视图所能看到的节点值（按从顶部到底部的顺序）【leetcode199/sword2-46 】
/*
		  1              <---
		/   \
	   2     3          <---
	  / \    /\
	 5  6    2 4       <---
       / \
	  1  2		  <---

	输出: [1,3,4,2]
*/

// rightSideView1 法一：深度优先搜索
// 思路：对树进行深度优先搜索，在搜索过程中，我们总是先访问右子树。那么对于每一层来说，我们在这层见到的第一个结点一定是最右边的结点。
// 这样一来，我们可以存储在每个深度访问的第一个结点，一旦我们知道了树的层数，就可以得到最终的结果数组。
func rightSideView1(root *TreeNode) []int {
	var (
		res []int // 用来储存结果
		f   func(node *TreeNode, depth int)
	)

	f = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if len(res) == depth { // 每一层只放一个元素，因此只要深度等于元素个数，说明此层已经处理过
			res = append(res, node.Val)
		}
		f(node.Right, depth+1) // 深度遍历，先遍历右侧元素
		f(node.Left, depth+1)
	}
	f(root, 0)

	return res
}

// rightSideView2 法二：广度优先遍历
// 思路：一层一层遍历，每层的最后一个元素就是右视图中看到的元素
func rightSideView2(root *TreeNode) []int {
	var (
		res   []int       // 结果集
		queue []*TreeNode // 队列，用于存储每一层的元素
	)
	if root == nil {
		return res
	} else {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		var size = len(queue)
		res = append(res, queue[size-1].Val) // 最后一个元素添加到结果集中
		for i := 0; i < size; i++ {          // 遍历次数=上一层的元素个数
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil { // 先添加左侧元素
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}
