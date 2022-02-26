package list

// 不修改数组找出重复的元素
//在一个长度为n+1的数组里所有数字都在1~n的范围里，所有数组中至少有一个数字是重复的，
//请找出任意一个重复的数字，但不修改原有的数组；如：输入长度为8的数组[3,2,5,4,7,2,6,5],
//则输出重复的元素为5或2

func getDuplications(array []int, length int) int {
	if len(array) == 0 {
		return -1
	}
	//[3,2,5,4,7,2,6,5] 长度为8，数值为1~7
	var start, end = 1, length - 1
	for start <= end {
		middle := ((end - start) >> 1) + start    // 获取中间值
		count := rangeCount(array, start, middle) // 获取此范围内数值的个数

		if start == end { // 结束条件：左右数值相等，且不只一个
			if count > 1 {
				return start
			} else {
				break
			}
		}

		if count > middle-start+1 { // 即此范围的个数比此范围应有的数量多，则重复元素一定在左侧，否则在右侧
			end = middle
		} else {
			start = middle + 1
		}

	}

	return -1
}

func rangeCount(array []int, start, end int) (count int) {
	for _, v := range array {
		if v >= start && v <= end {
			count++
		}
	}
	return count
}
