package sort

func QuickSort1(sli []int) []int {
	if len(sli) == 0 {
		return []int{}
	}

	//rand.Seed(time.Now().UnixNano())
	//idx:= rand.Intn()
	//rand.Int()

	var (
		// TODO
		pivot       = sli[0]
		left, right = make([]int, 0), make([]int, 0)
	)
	for i := 1; i < len(sli); i++ {
		if sli[i] > pivot {
			right = append(right, sli[i])
		} else {
			left = append(left, sli[i])
		}
	}
	return append(append(QuickSort1(left), pivot), QuickSort1(right)...)

}
