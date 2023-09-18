package dynamics

type item1 struct {
	nums1, nums2 []int
	expected     int
}

type item2 struct {
	num      int
	expected int
}

type item3 struct {
	nums     []int
	expected int
}

type item4 struct {
	num1, num2 int
	expected   int
}

type item5 struct {
	grid     [][]int
	expected int
}

type item6 struct {
	weight, value []int
	bagWeight     int
	expected      int
}
type item7 struct {
	weight, value, times []int
	bagWeight            int
	expected             int
}

type item8 struct {
	s, p     string
	expected bool
}

type item9 struct {
	text1, text2 string
	expected     int
}

type item10 struct {
	nums     []int
	target   int
	expected int
}

type item11 struct {
	nums     []int
	expected bool
}

//type item12 struct {
//	nums     []int
//	target   int
//	expected int
//}

type item13 struct {
	str []string
	item4
}

func maxInTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxInNums(l []int) int {
	for i := 1; i < len(l); i++ {
		if l[0] < l[i] {
			l[0] = l[i]
		}
	}
	return l[0]
}

func minInTwo(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minInThree(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		return c
	}
	return a
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}
