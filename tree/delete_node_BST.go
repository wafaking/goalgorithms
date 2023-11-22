package tree

import "goalgorithms/common"

//删除二叉搜索树中的节点(leetcode-450)
//给定一个二叉搜索树的根节点root和一个值key，删除二叉搜索树中的key对应的节点，并保证二叉搜索树的性质不变。
//返回二叉搜索树（有可能被更新）的根节点的引用。
//一般来说，删除节点可分为两个步骤：首先找到需要删除的节点；如果找到了，删除它。
//示例1:输入：root=[5,3,6,2,4,null,7],key=3,输出：[5,4,6,2,null,null,7]
//解释：给定需要删除的节点值是3，所以我们首先找到3这个节点，然后删除它。
//一个正确的答案是[5,4,6,2,null,null,7],如下图所示。
//另一个正确答案是[5,2,6,null,4,null,7]。
//示例2:输入:root=[5,3,6,2,4,null,7],key=0输出:[5,3,6,2,4,null,7]解释:二叉树不包含值为0的节点
//示例3:输入:root=[],key=0输出:[]
/*
		     10
			/  \
           6	14
	     /  \
		3    8
	  /  \  /  \
     1   4  7  9
*/

// 递归
func deleteNode1(root *common.TreeNode, key int) *common.TreeNode {
	var f func(root *common.TreeNode) *common.TreeNode
	f = func(root *common.TreeNode) *common.TreeNode {
		switch {
		case root == nil:
			return nil
		case key > root.Val:
			root.Right = f(root.Right)
		case key < root.Val:
			root.Left = f(root.Left)
		default:
			if root.Left == nil {
				return root.Right
			} else if root.Right == nil {
				return root.Left
			}
			//删除6,3上移，将4移至7左侧
			// 将左子节点B代替root节点A
			cur := root.Left
			// 找出root节点右子节点C的最左子节点M
			// 将左子节点B的右子节点放到M左子节点
			rootRight := root.Right
			for rootRight.Left != nil {
				rootRight = rootRight.Left
			}
			rootRight.Left = cur.Right
			cur.Right = root.Right
			return cur
		}
		return root
	}
	return f(root)
}

// 递归
func deleteNode2(root *common.TreeNode, key int) *common.TreeNode {
	switch {
	case root == nil:
		return nil
	case key < root.Val:
		root.Left = deleteNode2(root.Left, key)
	case key > root.Val:
		root.Right = deleteNode2(root.Right, key)
	default: // key == root.Val
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		// 删除6，7代替6的位置,将3移动到7左侧(注：需要删除7)
		rootRight := root.Right
		for rootRight.Left != nil {
			rootRight = rootRight.Left
		}
		// 删除7后将7的右子节点指定为8
		rootRight.Right = deleteNode2(root.Right, rootRight.Val)
		// 3仍在7的左侧
		rootRight.Left = root.Left
		return rootRight
	}
	return root
}

// 迭代(同法一的删除及移位方法)
func deleteNode3(root *common.TreeNode, key int) *common.TreeNode {
	var (
		cur = root
		pre *common.TreeNode
		// 用于记录当前节点是否是pre的左子节点
		// 可以用if pre != nil && pre.Left != nil && pre.Left.Val == cur.Val代替
		isLeft bool
	)

	for cur != nil {
		if cur.Val > key {
			pre = cur
			isLeft = true
			cur = cur.Left
			continue
		} else if cur.Val < key {
			pre = cur
			cur = cur.Right
			isLeft = false
			continue
		}

		// cur.Val == key
		if cur.Left == nil {
			if pre == nil {
				return cur.Right
			} else if isLeft {
				pre.Left = cur.Right
			} else {
				pre.Right = cur.Right
			}
		} else if cur.Right == nil {
			if pre == nil {
				return cur.Left
			} else if isLeft {
				pre.Left = cur.Left
			} else {
				pre.Right = cur.Left
			}
		} else {
			// neither left nor right is nil
			//让右子树的最左子节点代替删除的节点

			curLeft := cur.Left
			curRight := cur.Right
			for curRight.Left != nil {
				curRight = curRight.Left
			}
			curRight.Left = curLeft.Right
			curLeft.Right = cur.Right
			if pre == nil {
				return curLeft
			} else if isLeft {
				pre.Left = curLeft
			} else {
				pre.Right = curLeft
			}
		}
		break
	}

	return root
}

// 迭代(同法二的删除及移位方法)
func deleteNode4(root *common.TreeNode, key int) *common.TreeNode {
	var cur, curParent *common.TreeNode = root, nil
	for cur != nil && cur.Val != key {
		curParent = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if cur == nil {
		return root
	}
	if cur.Left == nil && cur.Right == nil {
		cur = nil
	} else if cur.Right == nil {
		cur = cur.Left
	} else if cur.Left == nil {
		cur = cur.Right
	} else {
		// 用右子节点的最左子节点代替当前节点
		successor, successorParent := cur.Right, cur
		for successor.Left != nil {
			successorParent = successor
			successor = successor.Left
		}
		if successorParent.Val == cur.Val {
			successorParent.Right = successor.Right
		} else {
			successorParent.Left = successor.Right
		}
		successor.Right = cur.Right
		successor.Left = cur.Left
		cur = successor
	}
	if curParent == nil {
		return cur
	}
	if curParent.Left != nil && curParent.Left.Val == key {
		curParent.Left = cur
	} else {
		curParent.Right = cur
	}
	return root
}
