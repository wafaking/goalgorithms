package array

// 两数之和II-输入有序数组(leetcode-167)
// 给定下标从1开始的整数数组numbers和目标数target，数组按非递减顺序排列，找出满足相加之和等于目标数target的两个数，并返回下标。
// 以长度为2的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
// 示例1：输入：numbers=[2,7,11,15], target = 9,输出：[1,2]
// 示例2：输入：numbers=[2,3,4], target = 6,输出：[1,3]
// 示例3：输入：numbers=[-1,0], target = -1,输出：[1,2]

// 双指针
func twoSum21(numbers []int, target int) []int {
	for start, end := 0, len(numbers)-1; start < end; {
		temp := numbers[start] + numbers[end]
		if temp > target {
			end--
		} else if temp < target {
			start++
		} else {
			return []int{start + 1, end + 1}
		}
	}
	return []int{}
}

// 二分查找
func twoSum22(numbers []int, target int) []int {
	var n = len(numbers)
	// 有序数组，固定一位，二分查找另一位
	for i := 0; i < n; i++ {
		expect := target - numbers[i]
		for l, r := i+1, n-1; l <= r; {
			mid := ((r - l) >> 1) + l
			if numbers[mid] == expect {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > expect {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return []int{}
}
