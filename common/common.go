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
	Num      int
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

type Item8 struct {
	Nums     []int
	Expected bool
}

type Item9 struct {
	Num1, Num2 int
	Expected   int
}

type Item10 struct {
	Str []string
	Item9
}

type Item11 struct {
	Nums1, Nums2 []int
	Expected     int
}

type Item12 struct {
	Grid     [][]int
	Expected int
}

type Item13 struct {
	Weight, Value []int
	BagWeight     int
	Expected      int
}

type Item14 struct {
	Weight, Value, Times []int
	BagWeight            int
	Expected             int
}

type Item15 struct {
	S, P     string
	Expected bool
}

type Item16 struct {
	Text1, Text2 string
	Expected     int
}

type Item17 struct {
	Nums     []int
	Expected []string
}

type Item18 struct {
	Nums     []int
	Target   int
	Expected bool
}

type Item19 struct {
	Nums1, Nums2 []int
	Expected     []int
}

type Item20 struct {
	Nums     []int
	Expected []int
}

type Item21 struct {
	Nums     []int
	Target   int
	Expected []int
}

type Item22 struct {
	Item2
	Nums1, Nums2 []int
}

type Item23 struct {
	Nums       []int
	Num1, Num2 int
	Expected   []int
}

type Item24 struct {
	NNode    *Node
	Expected []int
}
