package sort

// BubbleSort 普通的冒泡排序（严格意义上非冒泡排序）
func BubbleSort(sli []int) {
	length := len(sli)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if sli[i] > sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
			}
		}
	}
}

// BubbleSortStandard 标准冒泡排序
func BubbleSortStandard(sli []int) {
	for i := 0; i < len(sli); i++ {
		for j := len(sli) - 1; j > i; j-- {
			if sli[j-1] > sli[j] {
				sli[j-1], sli[j] = sli[j], sli[j-1]
			}
		}
	}
}

// BubbleSortStandardOptimize 标准冒泡排序优化(针对大部分有序的情况)
func BubbleSortStandardOptimize(sli []int) {
	// 1, 3, 4, 5, 6, 7
	sli = []int{4, 1, 3, 5, 6, 7}
	var flag = true // 设置标识位,默认为真
	for i := 1; i < len(sli) && flag; i++ {
		flag = false
		for j := len(sli) - 1; j >= i; j-- {
			if sli[j] < sli[j-1] {
				sli[j-1], sli[j] = sli[j], sli[j-1]
				flag = true // 有交换则为真
			}
		}
	}
}
