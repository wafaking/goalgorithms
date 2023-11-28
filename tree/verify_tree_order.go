package tree

//验证二叉搜索树的后序遍历序列(leetcode-LCR-152)
//请实现一个函数来判断整数数组postorder是否为二叉搜索树的后序遍历结果。
//postorder中无重复数字
//示例1：输入:postorder=[4,9,6,9,8]输出:false
//示例2：输入:postorder=[4,6,5,9,8]输出:true

// 递归
func verifyTreeOrder1(postorder []int) bool {
	if len(postorder) <= 2 {
		return true
	}
	// 左闭右闭
	var f func(l, r int) bool
	f = func(l, r int) bool {
		if l > r {
			return false
		} else if l == r {
			return true
		}

		var (
			root     = postorder[r]
			pivot    = r - 1
			hasPivot bool
		)
		for i := l; i < r; i++ {
			if postorder[i] >= root && !hasPivot {
				pivot = i
				hasPivot = true
				continue
			}
			if hasPivot && postorder[i] < root {
				return false
			}
		}
		return f(l, pivot) && f(pivot, r-1)
	}

	return f(0, len(postorder)-1)
}

// TODO 递归
func verifyTreeOrder2(postorder []int) bool {
	if len(postorder) < 2 {
		return true
	}

	index := len(postorder) - 1   // 区分左右子树：左子树上的值全都比根节点小，右子树上的值全都比根节点大
	rootValue := postorder[index] // 用来记录根节点的值

	for k, v := range postorder {
		// 当出现第一个大于根节点的值时，这个值往后全是右子树
		if index == len(postorder)-1 && v > rootValue {
			index = k
		}
		// 在右子树中出现小于根节点的值时，则该树不是二叉搜索树
		if index != len(postorder)-1 && rootValue > v {
			return false
		}
	}
	return verifyTreeOrder2(postorder[:index]) && verifyTreeOrder2(postorder[index:len(postorder)-1])
}
