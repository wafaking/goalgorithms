package common

type Item1 struct {
	Nums     []int
	Target   int
	Expected []int
}

type Item2 struct {
	Nums     []int
	Expected int
}

type Item3 struct {
	Nums     []int
	Target   int
	Expected int
}

type Item4 struct {
	Target   int
	Expected int
}

type Item5 struct {
	Num      int
	Expected [][]string
}

type Item6 struct {
	Nums, Expected []int
}

type Item7 struct {
	Nums     []int
	Expected [][]int
}
