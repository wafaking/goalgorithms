package common

func DiffIntSlice(sli1, sli2 []int) bool {
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

func DiffDoubleIntSlice(sli1, sli2 [][]int) bool {
	if len(sli1) != len(sli2) {
		return false
	}
	for i := range sli1 {
		if !DiffIntSlice(sli1[i], sli2[i]) {
			return false
		}
	}
	return true
}

func DiffStrSlice(sli1, sli2 []string) bool {
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

func MaxInTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxInNums(l ...int) int {
	for i := 1; i < len(l); i++ {
		if l[0] < l[i] {
			l[0] = l[i]
		}
	}
	return l[0]
}

func MinInTwo(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MinInThree(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		return c
	}
	return a
}

func Abs(num int) int {
	if num > 0 {
		return num
	}
	return -1 * num
}
