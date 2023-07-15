package dynamics

// 打家劫舍III(leetcode-337)
// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为root。
// 除了root之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，
// 聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
// 给定二叉树的root。返回在不触动警报的情况下，小偷能够盗取的最高金额。

//示例1: 输入:root=[3,2,3,null,3,null,1],输出:7; 解释:小偷一晚能够盗取的最高金额3+3+1=7
//示例2: 输入:root=[3,4,5,1,3,null,1], 输出:9; 解释:小偷一晚能够盗取的最高金额4+5=9
//示例3: 输入:root=[4,1,6,null,3,null,1],输出:9; 解释:小偷一晚能够盗取的最高金额6+3=9

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob31(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil { // 边界条件
			return 0, 0 // 不存在，怎么选都是0
		}

		lRob, lNotRob := dfs(node.Left)  // 递归左子树
		rRob, rNotRob := dfs(node.Right) // 递归右子树

		// 选当前节点(则左、右子节点都取不选的值)
		rob := lNotRob + rNotRob + node.Val

		// 不选当前节点,则可以选子节点(则左子树选与不选取大值)
		// 即选左子树不一定是最大值，因为选左子树时，左子树的子节点不能选，会影响结果
		// 右子树同理
		notRob := maxInTwo(lRob, lNotRob) + maxInTwo(rRob, rNotRob)
		return rob, notRob
	}

	return maxInTwo(dfs(root))
}
