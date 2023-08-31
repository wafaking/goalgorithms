package search

// 二分查找(leetcode-704)
// 给定一个有n个元素有序的（升序）整型数组nums和目标值target，请搜索nums中的target，如果目标值存在返回下标，否则返回-1。
// 示例1:输入:nums=[-1,0,3,5,9,12],target=9,输出:4
// 示例2:输入:nums=[-1,0,3,5,9,12],target=2输出:-1

// 法一：循环-左闭右闭区间
func binarySearch1(nums []int, target int) int {
	var low, high = 0, len(nums) - 1 // 左闭右闭区间，因此low<=high
	for low <= high {                //
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1 // 右闭，因此high=mid-1,nums[mid]不满足条件，直接去除
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 法二：循环-左闭右开区间
func binarySearch2(nums []int, target int) int {
	var low, high = 0, len(nums) // 左闭右开，因此选择low<high
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid // 右开区间，因此右边界包含mid
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 法三：递归：左闭右开区间
func binarySearch4(nums []int, target int) int {
	var dfs func(low, high int) int
	dfs = func(low, high int) int {
		if low >= high {
			return -1
		}
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			//high = mid
			return dfs(low, mid)
		} else {
			//low = mid + 1
			return dfs(mid+1, high)
		}
	}

	return dfs(0, len(nums))
}
