package array

import (
	"testing"
)

func TestMain(t *testing.M) {
	//sli := []int{5, 3, 7, 2, 0, 4, 1, 2}
	//sli := []int{4, 7, 5, 9, 2, 3}
	//sli := []int{2, 3, 1, 0, 2, 5, 3}
	//sli := []int{2, 7, 11, 15}
	//sli := []int{3, 2, 4}
	sli := []int{3, 0, -2, -1, -1, -1, 1, 2}
	//sli := []int{3, 2, 5, 4, 7, 2, 6, 5}
	//sli := []int{1, 2, 2, 3, 4, 5, 6, 7}
	//sli := []int{1, 2, 3, 4, 4}

	InitNums(sli)
	PrintNums()

	//Init2DNums()
	//Println2DNums()
	t.Run()
}

func TestTwoSum(t *testing.T) {
	var samples = []item{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, v := range samples {
		res := twoSum11(v.nums, v.target)
		if len(res) != len(v.expected) {
			t.Failed()
		}
		for i := 0; i < len(res); i++ {
			if res[i] != v.expected[i] {
				t.Fatal()
			}
		}
	}
	t.Log("------------SPLIT-------------")
	for _, v := range samples {
		res := twoSum12(v.nums, v.target)
		if len(res) != len(v.expected) {
			t.Failed()
		}
		for i := 0; i < len(res); i++ {
			if res[i] != v.expected[i] {
				t.Fatal()
			}
		}
	}
}

func TestTwoSum21(t *testing.T) {
	var list = []item{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}
	for _, item := range list {
		res := twoSum21(item.nums, item.target)
		t.Logf("res: %v, expected: %v", res, item.expected)
	}
	t.Log("-----------SPLIT---------")
}

func TestFindRepeatNumber(t *testing.T) {
	var list = []item{
		{[]int{2, 3, 1, 0, 2, 5, 3}, 0, nil},
		{[]int{3, 4, 2, 1, 1, 0}, 0, nil},
	}
	for _, v := range list {
		res := findRepeatNumber1(v.nums)
		t.Log("res: ", res)
	}
	t.Log("-----------SPLIT---------")
	for _, v := range list {
		res := findRepeatNumber2(v.nums)
		t.Log("res: ", res)
	}
	t.Log("-----------SPLIT---------")
	for _, v := range list {
		res := findRepeatNumber3(v.nums)
		t.Log("res: ", res)
	}
	t.Log("-----------SPLIT---------")
	for _, v := range list {
		res := findRepeatNumber4(v.nums)
		t.Log("res: ", res)
	}
}

func TestGetDuplications(t *testing.T) {
	res := getDuplications(nums, 8)
	t.Log("res: ", res)
}

func TestThreeSum(t *testing.T) {
	var samples = [][]int{
		//{0, 1, 1},
		//{0, 0, 0},
		//{-1, 0, 1, 2, -1, -4},
		{-4, -1, -1, -1, -1, 0, 1, 2, 2},
	}
	for _, nums := range samples {
		res := threeSum1(nums)
		t.Log("res: ", res)
	}
	t.Log("----------SPLIT------------")
	for _, nums := range samples {
		res := threeSum2(nums)
		t.Log("res: ", res)
	}
}

func TestThreeSumClosest(t *testing.T) {
	var samples = []struct {
		Nums     []int
		Target   int
		Expected int
	}{
		{Nums: []int{-1, 2, 1, -4}, Target: 1, Expected: 2},
		{Nums: []int{0, 0, 0}, Target: 1, Expected: 0},
		{Nums: []int{4, 0, 5, -5, 3, 3, 0, -4, -5}, Target: -1, Expected: -1},
	}
	for _, item := range samples {
		res := threeSumClosest1(item.Nums, item.Target)
		t.Logf("1--- %t, res: %d, expected:%d", res == item.Expected, res, item.Expected)
		res = threeSumClosest2(item.Nums, item.Target)
		t.Logf("2--- %t, res: %d, expected:%d", res == item.Expected, res, item.Expected)
		t.Log("----------SPLIT------------")
	}
}

func TestFourSumCount(t *testing.T) {
	var samples = [][4][]int{
		{
			{1, 2},
			{-2, -1},
			{-1, 2},
			{0, 2},
		},
		{
			{0},
			{0},
			{0},
			{0},
		},
	}
	for _, nums := range samples {
		res := fourSumCount11(nums[0], nums[1], nums[2], nums[3])
		t.Log("res: ", res)
		t.Log("----------SPLIT------------")
		res = fourSumCount12(nums[0], nums[1], nums[2], nums[3])
		t.Log("res: ", res)
		t.Log("----------SPLIT------------")
	}
}

func TestFourSum(t *testing.T) {
	var samples = []item{
		//{[]int{2, 2, 2, 2, 2}, 8, nil},
		//{[]int{1, 0, -1, 0, -2, 2}, 0, nil},
		{[]int{-4, -1, -1, -1, -1, 0, 1, 2, 2, 3}, 0, nil},
	}
	// 示例1： 输入：nums=[1,0,-1,0,-2,2], target=0, 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	// 示例2： 输入：nums = [2,2,2,2,2], target=8, 输出：[[2,2,2,2]]

	for _, item := range samples {
		res := fourSum(item.nums, item.target)
		t.Log("res: ", res)
	}
	t.Log("----------SPLIT------------")
}

func TestFindNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	matrix = [][]int{{-5}}
	res := findNumberIn2DArray(matrix, -5)
	t.Log("res: ", res)
}

func TestMerge(t *testing.T) {
	type Te struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	sli := []Te{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			nums2: []int{2, 5, 6},
			m:     3,
			n:     3,
		},
		{
			nums1: []int{0},
			nums2: []int{1},
			m:     0,
			n:     1,
		},
		{
			nums1: []int{2, 0},
			nums2: []int{1},
			m:     1,
			n:     1,
		},
		{
			nums1: []int{4, 5, 6, 0, 0, 0},
			nums2: []int{1, 2, 3},
			m:     3,
			n:     3,
		},
	}

	//nums1 := []int{0}
	//nums2 := []int{1}
	//m, n := 0, 1

	//nums1 := []int{2, 0}
	//nums2 := []int{1}
	//m, n := 1, 1

	for _, v := range sli {
		t.Log("nums1: ", v.nums1)
		t.Log("nums2: ", v.nums2)
		merge2(v.nums1, v.m, v.nums2, v.n)
		t.Log("res: ", v.nums1)
	}

}

func TestTotalDestinationPath(t *testing.T) {
	res := TotalDestinationPath(1, 1)
	t.Log("res: ", res)
}

func TestMinArray(t *testing.T) {
	var sliList = [][]int{
		{},
		{2},
		{3, 2},
		{3, 4, 5, 1, 2},
		{2, 2, 2, 0, 1},
		{3, 4, 5, 6, 1, 2},
		{1, 1, 1, 1, 1, 1, 0, 1, 1},
	}
	for _, sli := range sliList {
		// res := minArray(sli)
		res := minArray(sli)
		t.Logf("res: %d, sli: %+v", res, sli)
	}
}

func TestExistInMatrix(t *testing.T) {
	// var board = [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }
	// var word = "ABCCED"

	var board = [][]byte{
		{'a', 'b'},
		{'d', 'c'},
	}
	var word = "abcd"

	res := exist(board, word)
	t.Logf("res: %t", res)
}

func TestMovingCount(t *testing.T) {
	// res := movingCount(2, 3, 1) expected: 3
	// res := movingCount(3, 1, 0) // expected: 1
	// res := movingCount(3, 2, 17) // expected: 6
	res := movingCount(16, 8, 4) // expected: 15
	t.Logf("res: %d", res)
}

func TestGetWordSum(t *testing.T) {
	res := getWordSum(10)
	t.Logf("res: %d", res)
}

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea2(height)
	t.Logf("res: %d", res)
}

func TestPermute(t *testing.T) {
	var numsList = [][]int{
		//{},
		//{1},
		//{0, 1},
		{1, 2, 3},
	}
	for _, nums := range numsList {
		res := permute11(nums)
		t.Logf("res: %v\n", res)
		t.Log("-------------SPLIT----------")
		res = permute12(nums)
		t.Logf("res: %v\n", res)
		t.Log("-------------SPLIT----------")
		res = permute21(nums)
		t.Logf("res: %v\n", res)
		t.Log("-------------SPLIT----------")
		res = permute22(nums)
		t.Logf("res: %v\n", res)
		t.Log("-------------SPLIT----------")
		res = permute3(nums)
		t.Logf("res: %v\n", res)
		t.Log("-------------SPLIT----------")
		//res = permute32(nums)
		//t.Logf("res: %v\n", res)
		//t.Log("-------------SPLIT----------")
	}
}

func TestPermuteUnique(t *testing.T) {
	var numsList = [][]int{
		//{},
		//{1},
		//{0, 1},
		//{1, 2, 3},
		{1, 1, 2},
		//{1, 2, 1},
		//{1, 1, 1},
		{1, 2, 1, 1},
		{-1, 2, -1, 2, 1, -1, 2, 1},
		{0, 1, 0, 0, 9},
	}
	for _, nums := range numsList {
		res := permuteUnique11(nums)
		t.Logf("res: %v\n", res)
	}
	t.Log("----------SPLIT------------")
	for _, nums := range numsList {
		//res := permuteUnique2(nums)
		res := permuteUnique2(nums)
		t.Logf("res: %v\n", res)
	}
}

func TestCombine(t *testing.T) {
	var list = []item{
		{nums: []int{4}, target: 2, expected: nil},
		{nums: []int{1}, target: 1, expected: nil},
		//{nums: []int{1}, target: 2, expected: nil},
		{nums: []int{4}, target: 3, expected: nil},
	}
	for _, item := range list {
		res := combine11(item.nums[0], item.target)
		t.Logf("res: %v\n", res)
		t.Log("----------SPLIT------------")
		res = combine12(item.nums[0], item.target)
		t.Logf("res: %v\n", res)
		t.Log("----------SPLIT------------")
		res = combine13(item.nums[0], item.target)
		t.Logf("res: %v\n", res)
		t.Log("----------SPLIT------------")
		res = combine2(item.nums[0], item.target)
		t.Logf("res: %v\n", res)
		t.Log("----------SPLIT------------")
		res = combine3(item.nums[0], item.target)
		t.Logf("res: %v\n", res)
		t.Log("----------SPLIT------------")
	}
	//for _, nums := range numsList {
	//	//res := permuteUnique2(nums)
	//	res := permuteUnique2(nums)
	//	t.Logf("res: %v\n", res)
	//}
}

func TestCombinationSum1(t *testing.T) {
	var list = []item{
		{[]int{2}, 1, nil},
		{[]int{2, 3, 5}, 9, nil},
		{[]int{2, 3, 5}, 8, nil},
		{[]int{2, 3, 6, 7}, 7, nil},
		//{[]int{2, 2, 3, 5}, 9, nil},
		//{[]int{2, 2, 2}, 8, nil},
	}

	for _, item := range list {
		res := combinationSum11(item.nums, item.target)
		t.Logf("res: %v\n", res)
		t.Log("------------SPLIT-------------")
		res = combinationSum12(item.nums, item.target)
		t.Logf("res: %v\n", res)
		t.Log("------------SPLIT-------------")
		res = combinationSum13(item.nums, item.target)
		t.Logf("res: %v\n", res)
	}
}

func TestCombinationSum2(t *testing.T) {
	var list = []item{
		{[]int{2, 2, 2, 2, 4, 6}, 8, nil},
		{[]int{2, 2, 2, 2, 2, 2}, 8, nil},
		{[]int{1, 1, 1, 2}, 4, nil},
		{[]int{0, 0, 0, 0}, 0, nil},
	}
	for _, item := range list {
		res := combinationSum21(item.nums, item.target)
		t.Logf("res: %v\n", res)
		t.Log("------------SPLIT-------------")
		res = combinationSum22(item.nums, item.target)
		t.Logf("res: %v\n", res)
		t.Log("------------SPLIT-------------")
		res = combinationSum23(item.nums, item.target)
		t.Logf("res: %v\n", res)
		t.Log("------------SPLIT-------------")
		res = combinationSum24(item.nums, item.target)
		t.Logf("res: %v\n", res)
	}
}

func TestCombinationSum3(t *testing.T) {
	var list = []item{
		{[]int{3}, 7, nil},
		{[]int{4}, 7, nil},
		{[]int{3}, 9, nil},
		{[]int{9}, 45, nil},
		{[]int{4}, 1, nil},
	}

	for _, item := range list {
		res := combinationSum31(item.nums[0], item.target)
		t.Logf("res: %v\n", res)
		t.Log("------------SPLIT-------------")
		res = combinationSum32(item.nums[0], item.target)
		t.Logf("res: %v\n", res)
		t.Log("------------SPLIT-------------")
	}
}

func TestCombinationSum4(t *testing.T) {
	var list = []item{
		{[]int{1, 2, 3}, 7, []int{44}},
		{[]int{1, 2, 3}, 4, []int{7}},
		//{[]int{9}, 3, []int{0}},
		//{[]int{4, 2, 1}, 32, []int{0}},
	}
	// 示例1：输入:nums=[1,2,3], target=4，输出：7
	// 	所有可能的组合为：(1, 1, 1, 1),(1, 1, 2),(1, 2, 1),(1, 3),(2, 1, 1),(2, 2),(3, 1)
	// 示例2：输入：nums=[9], target=3,输出：0

	for _, item := range list {
		//res := combinationSum41(item.nums, item.target)
		//t.Logf("res: %v\n", res)
		//t.Log("------------SPLIT-------------")
		res := combinationSum42(item.nums, item.target)
		t.Logf("res: %v\n", res)
		t.Log("------------SPLIT-------------")
	}
}

func TestFindDuplicate(t *testing.T) {
	var list = []item2{
		{[]int{1, 2, 2, 3}, 2},
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{1, 2, 3, 3, 3, 5, 4}, 3},
	}
	for _, item := range list {
		res := findDuplicate1(item.nums)
		t.Logf("res: %t\n", res == item.expected)
	}
	t.Log("----------SPLIT------------")
	for _, item := range list {
		res := findDuplicate21(item.nums)
		t.Logf("res: %t\n", res == item.expected)
	}
	t.Log("----------SPLIT------------")
	for _, item := range list {
		res := findDuplicate22(item.nums)
		t.Logf("res: %t\n", res == item.expected)
	}
	t.Log("----------SPLIT------------")
	for _, item := range list {
		res := findDuplicate3(item.nums)
		t.Logf("res: %t\n", res == item.expected)
	}
	t.Log("----------SPLIT------------")
	for _, item := range list {
		res := findDuplicate4(item.nums)
		t.Logf("res: %t\n", res == item.expected)
	}
}

func TestSingleNumber(t *testing.T) {
	var list = []item2{
		{[]int{1, 2, 2}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}
	for _, item := range list {
		res := singleNumber(item.nums)
		t.Logf("res: %t\n", res == item.expected)
	}
	t.Log("----------SPLIT------------")
}

func TestSingleNumber2(t *testing.T) {
	var list = []item2{
		//{[]int{6, 3, 6, 6}, 3},
		//{[]int{2, 2, 3, 2}, 3},
		//{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
		//{[]int{-2, -2, 1, -2, 1, 4, 1}, 4},
		//{[]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}, -4},
		{[]int{-2, -2, -2, -4}, -4},
	}
	for _, item := range list {
		res := singleNumber21(item.nums)
		t.Logf("res: %t\n", res == item.expected)
	}
	t.Log("----------SPLIT------------")
	for _, item := range list {
		res := singleNumber22(item.nums)
		t.Logf("res: %t\n", res == item.expected)
	}
	t.Log("----------SPLIT------------")
	for _, item := range list {
		res := singleNumber23(item.nums)
		t.Logf("res: %t\n", res == item.expected)
	}
	t.Log("----------SPLIT------------")
}

func TestSingleNumber3(t *testing.T) {
	var list = []item{
		{[]int{1, 2, 1, 2, 3}, 0, nil},
		{[]int{2, 2, 3, 3, 4, 5}, 0, []int{4, 5}},
		{[]int{-2, -2, -3, -4, 1, 1}, 0, []int{-3, -4}},
		{[]int{-1, 0}, 0, []int{-1, 0}},
	}
	for _, item := range list {
		res := singleNumber31(item.nums)
		t.Logf("res: %v, expected: %v\n", res, item.expected)
	}
	t.Log("----------SPLIT------------")
	for _, item := range list {
		res := singleNumber32(item.nums)
		t.Logf("res: %v, expected: %v\n", res, item.expected)
	}
	t.Log("----------SPLIT------------")
}

func TestSubsets1(t *testing.T) {
	var list = []item{
		{[]int{1, 2, 3}, 0, nil},
		{[]int{0}, 0, nil},
		// 示例1：输入：nums=[1,2,3], 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
		// 示例2：输入：nums = [0], 输出：[[],[0]]
	}
	for _, item := range list {
		res := subsets1(item.nums)
		t.Logf("res: %v\n", res)
		t.Log("----------SPLIT------------")
		res = subsets2(item.nums)
		t.Logf("res: %v\n", res)
	}
}

func TestKthLargest(t *testing.T) {
	var nums = []int{4, 5, 8, 2}
	var list = []item4{
		{3, 4},
		{5, 5},
		{10, 5},
		{9, 8},
		{4, 8},
	}

	// [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]], 输出：[null, 4, 5, 5, 8, 8]
	// [4, 5, 8, 2]的第3大元素：4
	// [4, 5, 8, 2, 3]---> 4
	// [4, 5, 8, 2, 3, 5]---> 5
	// [4, 5, 8, 2, 3, 5, 10]---> 5
	// [4, 5, 8, 2, 3, 5, 10, 9]---> 8
	// [4, 5, 8, 2, 3, 5, 10, 9, 4]---> 8
	var kth = Constructor(3, nums)
	for _, item := range list {
		res := kth.Add(item.target)
		t.Logf("%t, expected: %d, res:%d", res == item.expected, item.expected, res)
	}
}

func TestFindKthLargest(t *testing.T) {
	var list = []item3{
		{nums: []int{6, 8, 10, 7, 4, 5, 9, 2, 3, 1}, target: 3, expected: 8},
		{nums: []int{3, 2, 1, 5, 6, 4}, target: 2, expected: 5},
		{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, target: 4, expected: 4},
	}

	for _, item := range list {
		var nums = item.nums
		res := findKthLargest1(item.nums, item.target)
		if res != item.expected {
			t.Fatalf("res: %d, origin:%+v", res, item)
		}

		t.Log("-------SPLIT-----------SPLIT------SPLIT----------")

		nums = item.nums
		res = findKthLargest2(nums, item.target)
		if res != item.expected {
			t.Fatalf("res: %d, origin:%+v", res, item)
		}
	}
}

func TestSolveNQueens(t *testing.T) {
	var list = []item5{
		{
			num: 1, expected: [][]string{{"Q"}},
		},
		{
			num: 4, expected: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
	}
	for _, item := range list {
		res := solveNQueens1(item.num)
		t.Logf("res: %v\n, expect:%v", res, item.expected)
		t.Log("----------SPLIT------------")
		//res = subsets2(item.nums)
		//t.Logf("res: %v\n", res)
	}
}
