package tree

import (
	"goalgorithms/common"
)

//有序链表转换二叉搜索树(leetcode-109)
//给定一个单链表的头节点head，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
//本题中，一个高度平衡二叉树是指一个二叉树每个节点的左右两个子树的高度差不超过1。
//示例1:输入:head=[-10,-3,0,5,9]输出:[0,-3,9,-10,null,5]
//解释:一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。
//示例2:输入:head=[]输出:[]

// 将链表转换成有序数组，再参考sortedArrayToBST1转换成二叉搜索树
func sortedListToBST1(head *common.ListNode) *TreeNode {
	var list = make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		list = append(list, cur.Val)
	}
	var f func(l, r int) *TreeNode
	f = func(l, r int) *TreeNode {
		if l >= r {
			return nil
		}
		mid := l + (r-l)>>1
		root := &TreeNode{Val: list[mid]}
		root.Left = f(l, mid)
		root.Right = f(mid+1, r)
		return root
	}
	return f(0, len(list))
}

// 分治(方法同法一，但每次都查找中间节点)
func sortedListToBST2(head *common.ListNode) *TreeNode {
	var (
		f       func(head, end *common.ListNode) *TreeNode
		findMid = func(head, end *common.ListNode) *common.ListNode {
			// head,end为前闭后开区间，因此不能包含end
			slow, fast := head, head
			for fast != end && fast.Next != end {
				fast = fast.Next.Next
				slow = slow.Next
			}
			return slow
		}
	)

	f = func(head, end *common.ListNode) *TreeNode {
		if head == end {
			return nil
		}

		mid := findMid(head, end)
		root := &TreeNode{Val: mid.Val}
		root.Left = f(head, mid)
		root.Right = f(mid.Next, end)
		return root
	}
	return f(head, nil)
}

// 分治+中序遍历优化(todo)
func sortedListToBST3(head *common.ListNode) *TreeNode {
	var (
		globalHead = head
		f          func(l, r int) *TreeNode
		getLength  = func(head *common.ListNode) int {
			var count int
			for ; head != nil; head = head.Next {
				count++
			}
			return count
		}
	)
	f = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		mid := (l + r + 1) / 2
		root := &TreeNode{}
		root.Left = f(l, mid-1)
		root.Val = globalHead.Val
		globalHead = globalHead.Next
		root.Right = f(mid+1, r)
		return root
	}

	return f(0, getLength(head)-1)
}
