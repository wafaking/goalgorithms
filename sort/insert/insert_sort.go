package insert

func InsertSort(sli []int) {
	length := len(sli)
	if length < 2 {
		return
	}

	for i := 1; i < length; i++ {
		for j := i; j > 0 && sli[j] < sli[j-1]; j-- {
			sli[j], sli[j-1] = sli[j-1], sli[j]
		}
	}
}
