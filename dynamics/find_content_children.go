package dynamics

import "sort"

// 分发饼干(leetcode-455)
// 一位家长想要给孩子们一些小饼干，但是每个孩子最多只能给一块饼干。
// 对于每个孩子i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；
// 并且每块饼干j，都有一个尺寸s[j]。如果s[j]>=g[i]，我们可以将这个饼干j分配给孩子i，这个孩子会得到满足。
// 你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

// 示例1: 输入:g=[1,2,3], s=[1,1], 输出:1, 解释:
// 有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
// 虽然有两块小饼干，由于尺寸都是1，只能让胃口值是1的孩子满足,所以输出1。

// 示例2: 输入:g=[1,2], s=[1,2,3], 输出:2, 解释:
// 有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
// 拥有的饼干数量和尺寸都足以让所有孩子满足, 所以该输出2.

func findContentChildren(g []int, s []int) int {
	var ans int
	sort.Ints(g)
	sort.Ints(s)
	for i, j := len(g)-1, len(s)-1; i >= 0 && j >= 0; i-- {
		if s[j] >= g[i] {
			ans++
			j--
		}
	}
	return ans
}