package sort

func SelectionSort(sli []int) {
	length := len(sli)

	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if sli[min] > sli[j] {
				min = j
			}
		}
		if min != i {
			sli[i], sli[min] = sli[min], sli[i]
		}
	}
}

func selectSort(sli []int) {
	var min int
	for i := 0; i < len(sli); i++ {
		for j := i + 1; j < len(sli); j++ {
			if sli[j] < sli[min] {
				sli[min], sli[j] = sli[j], sli[min]
				break
			}
		}
	}
}
