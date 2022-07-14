package sort

func InsertSort1(sli []int) {
	for i := 1; i < len(sli); i++ {
		for j := i; j > 0 && sli[j] < sli[j-1]; j-- {
			sli[j], sli[j-1] = sli[j-1], sli[j]
		}
	}
}

func InsertSort2(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if sli[j] < sli[j-1] {
				sli[j-1], sli[j] = sli[j], sli[j-1]
			} else {
				break
			}
		}
	}
}
