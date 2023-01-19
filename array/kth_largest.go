package array

import (
	"math/rand"
	"time"
)

// 数组中的第K个最大元素(leetcode-215/sword-2-76)
// 在未排序的数组中找到第k个最大的元素,要求实现时间复杂度为O(n)。
// 示例1: 输入:[3,2,1,5,6,4]和k=2,输出:5;
// 示例2: 输入:[3,2,3,1,2,4,5,5,6]和k=4,输出:4;

// findKthLargest1 借助快排分隔的方法
func findKthLargest1(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

// quickSelect 生成随机数r，用r将nums分成小大两部分,返回随机数所在的第kth位置
// kthNum为数组中的第k个数(如len(n)=10, 则第3大的数为,有序数组的索引为7)
func quickSelect(nums []int, left, right, kthIdx int) int {
	// 1. 根据left、right生成两者之间的随机数位置下标
	var randIdx = rand.Int()%(right-left+1) + left
	// 2. 将随机数放在尾部
	nums[randIdx], nums[right] = nums[right], nums[randIdx]

	// 3. 获取排序后随机数所在位置
	var pos = partition(nums, left, right)
	if pos == kthIdx {
		return nums[pos]
	} else if pos < kthIdx {
		// 说明在右侧,继续将右侧数据排序，再找出kth
		return quickSelect(nums, pos, right, kthIdx)
	} else { // 否则，一定在左侧
		return quickSelect(nums, left, pos-1, kthIdx)
	}
}

// partition 根据随机数将数组分成小,randNum,大三部分
func partition(nums []int, left, right int) int {
	var (
		ref = nums[right]
		pos = left
	)

	// 都与随机数比较，将数组"大致"有序排列
	for i := left; i < right; i++ {
		if nums[i] <= ref {
			nums[pos], nums[i] = nums[i], nums[pos]
			pos++
		}
	}

	// 将随机数移到大小两部分数据的中间位置
	nums[pos], nums[right] = nums[right], nums[pos]
	// 返回随机数的索引
	return pos
}

// findKthLargest2 使用大顶堆(构建大顶堆后，删除n-k+1个元素即可)
func findKthLargest2(nums []int, k int) int {
	var size = len(nums)
	buildMaxHeap(nums)
	for i := 0; i < k-1; i++ {
		nums[0], nums[size-1] = nums[size-1], nums[0]
		size--
		maxHeapify(nums, 0, size-1)
	}

	return nums[0]
}

// buildMaxHeap 构建大顶堆
func buildMaxHeap(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		maxHeapify(nums, i, len(nums))
	}
}

func maxHeapify(nums []int, parent, size int) {
	// largest用来记录值最大的节点
	subLeft, subRight, largest := 2*parent+1, 2*parent+2, parent

	// 获取左、右子节点中值较大的节点
	if subLeft < size && nums[subLeft] > nums[largest] {
		largest = subLeft
	}

	if subRight < size && nums[subRight] > nums[largest] {
		largest = subRight
	}

	// 如果 largest != parent表示最大的节点不是父节点
	if largest != parent {
		// 将父节点与子节点中值较大的元素交换
		nums[parent], nums[largest] = nums[largest], nums[parent]
		// 向下继续交换
		maxHeapify(nums, largest, size)
	}
}
