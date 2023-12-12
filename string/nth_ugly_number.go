package str

import (
	"container/heap"
	"goalgorithms/common"
	"sort"
)

//丑数II(leetcode-264)
//给你一个整数n，请你找出并返回第n个丑数。
//丑数就是质因子只包含2、3和5的正整数。
//示例1：输入：n=10输出：12解释：[1,2,3,4,5,6,8,9,10,12]是由前10个丑数组成的序列。
//示例2：输入：n=1输出：1解释：1通常被视为丑数。
//提示：
//	1<=n<=1690

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

// 最小堆
func nthUglyNumber1(n int) int {
	if n == 1 {
		return 1
	}

	var (
		factors = []int{2, 3, 5}
		m       = map[int]struct{}{1: {}}
		ans     int
		h       = &hp{IntSlice: []int{1}}
	)

	for i := 1; i <= n; i++ {
		ans = heap.Pop(h).(int)
		for _, v := range factors {
			if _, ok := m[ans*v]; !ok {
				m[ans*v] = struct{}{}
				heap.Push(h, ans*v)
			}
		}
	}
	return ans
}

// 动态规划
func nthUglyNumber2(n int) int {
	var (
		//dp[i]表示第i个丑数的值
		dp         = make([]int, n+1)
		p2, p3, p5 = 1, 1, 1
	)
	// 初始化
	dp[1] = 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = common.MinInTwo(common.MinInTwo(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}

	return dp[n]
}
