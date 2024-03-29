package array

//旋转数组的最小数字(sword11)
//把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
//给你一个可能存在重复元素值的数组numbers，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。
//请返回旋转数组的最小元素。例如，数组[3,4,5,1,2]为[1,2,3,4,5]的一次旋转，该数组的最小值为1。
//示例1：输入：[3,4,5,1,2]输出：1
//示例2：输入：[2,2,2,0,1]输出：0

// 二分查找
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	// [3,4,5,1,2]
	// [2,1,5,6,7]
	var low, high = 0, len(numbers) - 1
	for low < high {
		mid := low + (high-low)>>1
		// 注：此处要比较mid与high
		if numbers[mid] < numbers[high] {
			high = mid
		} else if numbers[mid] > numbers[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return numbers[low]
}
