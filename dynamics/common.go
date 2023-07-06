package dynamics

type item1 struct {
	Sli1, Sli2 []int
	Expected   int
}

type item2 struct {
	Num      int
	Expected int
}

type item3 struct {
	Sli      []int
	Expected int
}

type item4 struct {
	Num1, Num2 int
	Expected   int
}

type item5 struct {
	Grid     [][]int
	Expected int
}

type item6 struct {
	Weight, Value []int
	BagWeight     int
	Expected      int
}
type item7 struct {
	Weight, Value, Times []int
	BagWeight            int
	Expected             int
}

func MaxInTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}
