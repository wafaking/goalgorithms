package list

//数组中的第K个最大元素(leet 215)

//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//示例 1:
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5

//示例 2:
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
func findKthLargest(nums []int, k int) int {
	return 0
}

// func findKthLargest(nums []int, k int) int {
// 	return doFindKthLargest(nums, k)
// }

// func doFindKthLargest(nums []int, k int) int {
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}
// 	var left, right []int
// 	var start = nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] <= start { // TODO 边界判断
// 			left = append(left, nums[i])
// 		} else {
// 			right = append(right, nums[i])
// 		}
// 	}
// 	if len(left) == k {
// 		return doFindKthLargest(left, k)
// 	}

// 	switch len(nums) {
// 	case 0:
// 		return 0
// 	case 1:
// 		return nums[0]
// 	case 2:
// 		fmt.Println("---: ", nums)
// 		if nums[0] > nums[1] {
// 			return nums[1]
// 		} else {
// 			return nums[0]
// 		}
// 	default:
// 		ref := nums[0]
// 		var newNums []int
// 		for i := 1; i < len(nums); i++ {
// 			if nums[i] > ref {
// 				newNums = append(newNums, nums[i])
// 			}
// 		}
// 		return doFindKthLargest(newNums)
// 	}
// }

// func findKthLargest1(nums []int, k int) int {
// 	return doFindKthLargest1(nums, k, 0, len(nums)-1)
// }

// func doFindKthLargest1(nums []int, k int, start, end int) int {
// 	nlen := len(nums)
// 	targetPos := nlen - k

// 	for {
// 		pivotIndex := partition(nums, start, end)
// 		if pivotIndex == targetPos {
// 			return nums[pivotIndex]
// 		} else if pivotIndex < targetPos {
// 			start = pivotIndex + 1
// 		} else {
// 			end = pivotIndex - 1
// 		}
// 	}
// }

// func partition(nums []int, start, end int) int {
// 	pivot := nums[start]
// 	left, right := start+1, end

// 	for left <= right {
// 		for left <= right && nums[left] <= pivot {
// 			left++
// 		}
// 		for left <= right && nums[right] > pivot {
// 			right--
// 		}
// 		if left <= right {
// 			nums[left], nums[right] = nums[right], nums[left]
// 		}
// 	}
// 	nums[right], nums[start] = nums[start], nums[right]
// 	return right
// }

func findKthLargest1(nums []int, k int) int {
	return _findKthSmallest(nums, 0, len(nums)-1, len(nums)-k)
}

func _findKthSmallest(arr []int, l, r, k int) int {
	if l >= r {
		return arr[l]
	}
	p := _partition(arr, l, r)
	if k > p {
		return _findKthSmallest(arr, p+1, r, k)
	} else if k == p {
		return arr[p]
	} else {
		return _findKthSmallest(arr, l, p-1, k)
	}
}

func _partition(arr []int, l, r int) int {
	v := arr[l]
	j := l
	i := j + 1
	for i = j + 1; i <= r; i++ {
		if arr[i] < v {
			arr[j+1], arr[i] = arr[i], arr[j+1]
			j++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}
