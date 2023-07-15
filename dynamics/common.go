package dynamics

type item1 struct {
	Sli1, Sli2 []int
	expected   int
}

type item2 struct {
	Num      int
	expected int
}

type item3 struct {
	Sli      []int
	expected int
}

type item4 struct {
	Num1, Num2 int
	expected   int
}

type item5 struct {
	Grid     [][]int
	expected int
}

type item6 struct {
	Weight, Value []int
	BagWeight     int
	expected      int
}
type item7 struct {
	Weight, Value, Times []int
	BagWeight            int
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
	sli      []int
	target   int
	expected int
}

func maxInTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
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
