package array

// 数组中的第K个最大元素(leetcode-215/sword-2-76)
// 在未排序的数组中找到第k个最大的元素,要求实现时间复杂度为O(n)。
// 示例1: 输入:[3,2,1,5,6,4]和k=2,输出:5;
// 示例2: 输入:[3,2,3,1,2,4,5,5,6]和k=4,输出:4;

func findKthLargest1(nums []int, k int) int {
	return 0
}

func findKthLargest2(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}
