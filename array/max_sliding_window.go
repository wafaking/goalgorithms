package array

import (
	"container/heap"
	"sort"
)

//滑动窗口最大值(leetcode-239)
//给你一个整数数组nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的k个数字。滑动窗口每次只向右移动一位。
//返回滑动窗口中的最大值。
//示例1：输入：nums=[1,3,-1,-3,5,3,6,7],k=3输出：[3,3,5,5,6,7]
//解释：
//	滑动窗口的位置                最大值
//	---------------               -----
//	[1  3  -1] -3  5  3  6  7      3
//	1 [3  -1  -3] 5  3  6  7       3
//	1  3 [-1  -3  5] 3  6  7       5
//	1  3  -1 [-3  5  3] 6  7       5
//	1  3  -1  -3 [5  3  6] 7       6
//	1  3  -1  -3  5 [3  6  7]      7
//示例2：输入：nums=[1],k=1输出：[1]
//注：1<=k<=nums.length,1<=nums.length<=105

// 队列
func maxSlidingWindow1(nums []int, k int) []int {
	// 使用队列；往Q中添加一个num时，将队前小于此num的元素都替换成当前元素
	var (
		ans   = make([]int, 0)
		queue = make([]int, 0, k)
		push  = func(num int) {
			for j := len(queue) - 1; j >= 0; j-- {
				if num <= queue[j] {
					break
				}
				queue[j] = num
			}
			// 将当前元素添加到queue尾部
			queue = append(queue, num)
		}
	)
	// 1.先构造queue
	for i := 0; i < k; i++ {
		push(nums[i])
	}

	if len(nums) <= k {
		return []int{queue[0]}
	}

	// 2.再添加其它元素
	for i := k; i < len(nums); i++ {
		ans = append(ans, queue[0])
		queue = queue[1:]
		push(nums[i])
	}
	ans = append(ans, queue[0])
	return ans
}

// 单调队列（TODO）
func maxSlidingWindow2(nums []int, k int) []int {
	var (
		q    = make([]int, 0)
		push = func(i int) {
			for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
				q = q[:len(q)-1]
			}
			q = append(q, i)
		}
	)

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// 优先级队列(TODO)
func maxSlidingWindow3(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

// 分块+预处理(TODO)
func maxSlidingWindow4(nums []int, k int) []int {
	n := len(nums)
	prefixMax := make([]int, n)
	suffixMax := make([]int, n)
	for i, v := range nums {
		if i%k == 0 {
			prefixMax[i] = v
		} else {
			prefixMax[i] = max(prefixMax[i-1], v)
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || (i+1)%k == 0 {
			suffixMax[i] = nums[i]
		} else {
			suffixMax[i] = max(suffixMax[i+1], nums[i])
		}
	}

	ans := make([]int, n-k+1)
	for i := range ans {
		ans[i] = max(suffixMax[i], prefixMax[i+k-1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
