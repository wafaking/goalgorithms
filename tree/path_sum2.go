package tree

import (
	"goalgorithms/common"
)

//路径总和II(leetcode-113/LCR-153)
//给你二叉树的根节点root和一个整数目标和targetSum，找出所有从根节点到叶子节点路径总和等于给定目标和的路径。
//叶子节点是指没有子节点的节点。
//示例1：输入：root=[5,4,8,11,null,13,4,7,2,null,null,5,1],targetSum=22输出：[[5,4,11,2],[5,8,4,5]]
//示例2：输入：root=[1,2,3],targetSum=5输出：[]
//示例3：输入：root=[1,2],targetSum=0输出：[]
/*
		 5
		/ \
	   4   8
	  /   / \
	 11  13  4
	/  \    / \
   7    2  5   1
给定，root和targetSum=22,返回:
	[
		[5,4,11,2],
		[5,8,4,5]
	]
*/

// 深度优先遍历
func pathSum21(root *common.TreeNode, targetSum int) [][]int {

	var (
		path = make([]int, 0)
		ans  = make([][]int, 0)
		f    func(root *common.TreeNode, targetSum int)
	)
	f = func(root *common.TreeNode, targetSum int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		targetSum -= root.Val
		if root.Left == nil && root.Right == nil && targetSum == 0 {
			var temp = make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		if root.Left != nil {
			f(root.Left, targetSum)
			path = path[:len(path)-1]
		}

		if root.Right != nil {
			f(root.Right, targetSum)
			path = path[:len(path)-1]
		}
		targetSum += root.Val
	}
	f(root, targetSum)
	return ans
}

// 广度优先遍历
func pathSum22(root *common.TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		ans         = make([][]int, 0)
		queue       = []*common.TreeNode{root} // 存储节点元素的队列
		targetQueue = []int{targetSum}         // 对应每个节点的target
		pathQueue   = [][]int{{root.Val}}      // 对应每个节点的path
	)

	for len(queue) > 0 {
		var (
			cur       = queue[0]
			curTarget = targetQueue[0]
			curPath   = pathQueue[0]
		)
		queue = queue[1:]
		targetQueue = targetQueue[1:]
		pathQueue = pathQueue[1:]

		curTarget -= cur.Val
		if cur.Left == nil && cur.Right == nil {
			if curTarget == 0 {
				ans = append(ans, curPath)
			}
			continue
		}

		if cur.Left != nil {
			queue = append(queue, cur.Left)
			targetQueue = append(targetQueue, curTarget)
			// 注：此处不能直接在curPath添加元素,cap(curPath)=4,len(curPath)=3,
			//    则添加后curPath会变,多次添加最后一个元素会被覆盖
			pathQueue = append(pathQueue, append(append([]int{}, curPath...), cur.Left.Val))
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			targetQueue = append(targetQueue, curTarget)
			pathQueue = append(pathQueue, append(append([]int{}, curPath...), cur.Right.Val))
		}
	}

	return ans
}

// 广度优先遍历
func pathSum23(root *common.TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		ans         = make([][]int, 0)
		queue       = []*common.TreeNode{root}    // 存储节点元素的队列
		targetQueue = []int{targetSum - root.Val} // 对应每个节点的target
		pathQueue   = [][]int{{root.Val}}         // 对应每个节点的path
	)
	for len(queue) > 0 {
		var tempQueue = queue
		queue = nil
		for _, cur := range tempQueue {
			var (
				curTarget = targetQueue[0]
				curPath   = pathQueue[0]
			)
			targetQueue = targetQueue[1:]
			pathQueue = pathQueue[1:]

			//curTarget -= cur.Val
			if cur.Left == nil && cur.Right == nil {
				if curTarget == 0 {
					ans = append(ans, curPath)
				}
				continue
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
				targetQueue = append(targetQueue, curTarget-cur.Left.Val)
				pathQueue = append(pathQueue, append(append([]int{}, curPath...), cur.Left.Val))
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
				targetQueue = append(targetQueue, curTarget-cur.Right.Val)
				pathQueue = append(pathQueue, append(append([]int{}, curPath...), cur.Right.Val))
			}
		}
	}
	return ans
}
