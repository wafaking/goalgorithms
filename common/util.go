package common

func DiffSlice(sli1, sli2 []int) bool {
	if len(sli1) != len(sli2) {
		return false
	}
	for i := range sli1 {
		if sli1[i] != sli2[i] {
			return false
		}
	}
	return true
}

func DiffDoubleSlice(sli1, sli2 [][]int) bool {
	if len(sli1) != len(sli2) {
		return false
	}
	for i := range sli1 {
		if !DiffSlice(sli1[i], sli2[i]) {
			return false
		}
	}
	return true
}
