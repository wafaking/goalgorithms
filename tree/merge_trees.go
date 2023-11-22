package tree

import "goalgorithms/common"

//合并二叉树(leetcode-617)
//给你两棵二叉树：root1和root2。
//将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。
//合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；否则，不为null的节点将直接作为新二叉树的节点。
//返回合并后的二叉树。注意:合并过程必须从两个树的根节点开始。
//示例1：输入：root1=[1,3,2,5],root2=[2,1,3,null,4,null,7]输出：[3,4,5,5,4,null,7]
//示例2：输入：root1=[1],root2=[1,2]输出：[2,2]

// 深度优先(合并到root1上)
func mergeTrees1(root1 *common.TreeNode, root2 *common.TreeNode) *common.TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val

	root1.Left = mergeTrees1(root1.Left, root2.Left)
	root1.Right = mergeTrees1(root1.Right, root2.Right)
	return root1
}

// 广度优先
func mergeTrees2(root1 *common.TreeNode, root2 *common.TreeNode) *common.TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	var (
		queue1 = []*common.TreeNode{root1}
		queue2 = []*common.TreeNode{root2}
	)

	for len(queue1) > 0 && len(queue2) > 0 {
		var (
			cur1 = queue1[0]
			cur2 = queue2[0]
		)

		queue1 = queue1[1:]
		queue2 = queue2[1:]

		// cur1，cur2有空的情况
		if cur1 == nil && cur2 == nil {
			continue
		} else if cur1 == nil {
			cur1 = cur2
			continue
		} else if cur2 == nil {
			continue
		}

		// cur1 != nil && cur2 !=nil
		cur1.Val += cur2.Val

		if cur1.Left != nil || cur2.Left != nil {
			if cur1.Left != nil && cur2.Left != nil {
				queue1 = append(queue1, cur1.Left)
				queue2 = append(queue2, cur2.Left)
			} else if cur1.Left == nil {
				cur1.Left = cur2.Left
			}
		}

		if cur1.Right != nil || cur2.Right != nil {
			if cur1.Right != nil && cur2.Right != nil {
				queue1 = append(queue1, cur1.Right)
				queue2 = append(queue2, cur2.Right)
			} else if cur1.Right == nil {
				cur1.Right = cur2.Right
			}
		}
	}
	return root1
}
