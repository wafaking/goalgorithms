package array

import (
	"container/heap"
	"sort"
)

// 数据流中的第K大元素(leetcode-703/sword2-59)
// 设计一个找到数据流中第k大元素的类（class）。注意是排序后的第k大元素，不是第k个不同的元素。
// 请实现KthLargest类：
//   KthLargest(int k, int[] nums) 使用整数k和整数流nums初始化对象。
//	 int add(int val)将val插入数据流nums后，返回当前数据流中第k大的元素。

// 示例：输入：["KthLargest", "add", "add", "add", "add", "add"],
// [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]], 输出：[null, 4, 5, 5, 8, 8]
// [4, 5, 8, 2]的第3大元素：4
// [4, 5, 8, 2, 3]---> 4
// [4, 5, 8, 2, 3, 5]---> 5
// [4, 5, 8, 2, 3, 5, 10]---> 5
// [4, 5, 8, 2, 3, 5, 10, 9]---> 8
// [4, 5, 8, 2, 3, 5, 10, 9, 4]---> 8

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// 使用含有k个元素的优先级队列，小顶堆的末尾元素即为第k大元素

// KthLargest 注：KthLargest必须实现sort.Interface接口
type KthLargest struct {
	sort.IntSlice     // 最小堆，堆顶元素
	size          int // 堆的大小
}

func Constructor(k int, nums []int) KthLargest {
	var kthLargest = KthLargest{size: k}
	for _, v := range nums { // 初始化，将元素添加到堆
		kthLargest.Add(v)
	}
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	// 未超过size或当要添加的元素大于堆顶元素时,直接添加
	if len(this.IntSlice) < this.size || this.IntSlice[0] < val {
		heap.Push(this, val)
	}

	if len(this.IntSlice) > this.size {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}

// Push push为堆接口必须实现的方法
func (this *KthLargest) Push(x interface{}) {
	this.IntSlice = append(this.IntSlice, x.(int))
}

// Pop pop为堆接口必须实现的方法(pop the minimum value)
func (this *KthLargest) Pop() interface{} {
	// 数组的尾部元素为第k大元素
	var v = this.IntSlice[len(this.IntSlice)-1]
	this.IntSlice = this.IntSlice[:len(this.IntSlice)-1]
	return v
}
