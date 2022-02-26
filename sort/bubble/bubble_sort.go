package bubble

// 普通的冒泡排序（严格意义上非冒泡排序）
func BubbleSort(sli []int) {
	// fmt.Println("before sorted: ", sli)
	length := len(sli)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if sli[i] > sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
			}
		}
	}
	// fmt.Println("after sorted: ", sli)
}

// 标准冒泡排序
func BubbleSortStandard(sli []int) {
	length := len(sli)

	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if sli[j-1] > sli[j] {
				sli[j-1], sli[j] = sli[j], sli[j-1]
			}
		}
	}
}

// 标准冒泡排序优化
func BubbleSortStandardOptimize(sli []int) {
	length := len(sli)

	var flag = true // 设置标识位,默认为真
	for i := 1; i < length && flag; i++ {
		flag = false
		for j := length - 1; j >= i; j-- {
			if sli[j] < sli[j-1] {
				sli[j-1], sli[j] = sli[j], sli[j-1]
				flag = true // 有交换则为真
			}
		}
	}
}
